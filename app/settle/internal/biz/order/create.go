package order

import (
	"context"

	"github.com/zxq97/x-finance/api/settle/service/v1"
	"github.com/zxq97/x-finance/app/settle/internal/biz"
)

func (uc *OrderUseCase) preCheckCreate(ctx context.Context, bizType v1.BizType, id, subID int64, orderType v1.OrderType) error {
	if bizType == v1.BizType_BizTypeUnknown {
		return biz.ErrBizTypeIsInvalid
	} else if orderType == v1.OrderType_OrderTypeUnknown {
		return biz.ErrOrderTypeIsInvalid
	} else if id == 0 || subID == 0 {
		return biz.ErrOrderIDIsInvalid
	}

	if orderType == v1.OrderType_OrderTypeNormal && id != subID {
		return biz.ErrOrderIDIsInvalid
	}

	order, err := uc.repo.GetOneByID(ctx, int8(bizType), id, subID)
	if err != nil {
		return err
	} else if order != nil {
		return biz.ErrOrderIDIsDuplicate
	}

	return nil
}

func (uc *OrderUseCase) createNormal(ctx context.Context, order *biz.OrderCreateParam) (*biz.PayInfo, error) {
	mchInfo, err := uc.merchant.GetMerchantInfo(ctx, order.Source, order.SupplierCode)
	if err != nil {
		return nil, err
	}

	platformRate := 100 - mchInfo.Rate
	platformAmt := order.TargetAmt * int64(platformRate) / 100
	mchAmt := order.TargetAmt - platformAmt

	order.Order.SettleDetails = make([]*biz.SettleDetail, 2)

	order.Order.SettleDetails[0] = &biz.SettleDetail{
		MerchantID: platformMerchantID,
		Rate:       platformRate,
		SettleAmt:  platformAmt,
		SubID:      order.SubID,
		TargetType: targetTypePlatform,
	}
	order.Order.SettleDetails[1] = &biz.SettleDetail{
		MerchantID: mchInfo.MerchantID,
		Rate:       mchInfo.Rate,
		SettleAmt:  mchAmt,
		SubID:      order.SubID,
		TargetType: targetTypeMerchant,
	}

	if err = uc.repo.Create(ctx, order); err != nil {
		return nil, err
	}

	return uc.cashier.CreateOrder(ctx, &biz.CreateOrderParam{})
}

func (uc *OrderUseCase) createMerge(ctx context.Context, order *biz.OrderCreateParam) (*biz.PayInfo, error) {
	panic("implement me")
}

func (uc *OrderUseCase) createGoods(ctx context.Context, order *biz.OrderCreateParam) (*biz.PayInfo, error) {
	order.SettleDetails = append(order.SettleDetails, &biz.SettleDetail{
		MerchantID: platformMerchantID,
		Rate:       100,
		SettleAmt:  order.TargetAmt,
		SubID:      order.SubID,
		TargetType: targetTypePlatform,
	})

	if err := uc.repo.Create(ctx, order); err != nil {
		return nil, err
	}

	return uc.cashier.CreateOrder(ctx, &biz.CreateOrderParam{})
}

func (uc *OrderUseCase) createSublet(ctx context.Context, order *biz.OrderCreateParam) (*biz.PayInfo, error) {
	if order.MerchantCode == "" {
		return nil, biz.ErrMerchantCodeIsInvalid
	}

	mchInfo, err := uc.merchant.GetMerchantInfo(ctx, order.Source, order.SupplierCode)
	if err != nil {
		return nil, err
	}

	subletInfo, err := uc.sublet.GetSubletInfo(ctx, order.MerchantCode)
	if err != nil {
		return nil, err
	}

	platformRate := 100 - mchInfo.Rate - subletInfo.Rate
	platformAmt := order.TargetAmt * int64(platformRate) / 100
	subletAmt := order.TargetAmt * int64(subletInfo.Rate) / 100
	mchAmt := order.TargetAmt*platformAmt - subletAmt

	order.SettleDetails = make([]*biz.SettleDetail, 3)
	order.SettleDetails[0] = &biz.SettleDetail{
		MerchantID: platformMerchantID,
		Rate:       platformRate,
		SettleAmt:  platformAmt,
		SubID:      order.SubID,
		TargetType: targetTypePlatform,
	}
	order.SettleDetails[1] = &biz.SettleDetail{
		MerchantID: subletInfo.ThirdMerchantID,
		Rate:       subletInfo.Rate,
		SettleAmt:  subletAmt,
		SubID:      order.SubID,
		TargetType: targetTypeThird,
	}
	order.SettleDetails[2] = &biz.SettleDetail{
		MerchantID: subletInfo.SubletMerchantID,
		Rate:       mchInfo.Rate,
		SettleAmt:  mchAmt,
		SubID:      order.SubID,
		TargetType: targetTypeSublet,
	}

	if err = uc.repo.Create(ctx, order); err != nil {
		return nil, err
	}

	return uc.cashier.CreateOrder(ctx, &biz.CreateOrderParam{})
}

func (uc *OrderUseCase) Create(ctx context.Context, req *v1.CreateOrderRequest) (res *v1.CreateOrderResponse, err error) {
	if err = uc.preCheckCreate(ctx, req.BizType, req.OrderId, req.SubOrderId, req.OrderType); err != nil {
		return nil, err
	}

	order := &biz.OrderCreateParam{
		Order: &biz.Order{
			BizType:        int8(req.BizType),
			UID:            req.Uid,
			ID:             req.OrderId,
			SubID:          req.SubOrderId,
			StoreAmt:       req.StoreAmount,
			TargetAmt:      req.TargetAmount,
			RealPayAmt:     req.RealPayAmount,
			CouponAmt:      req.CouponAmount,
			PromotionAmt:   req.PromotionAmount,
			MchDiscountAmt: req.MerchantDiscount,
			OrderType:      int8(req.OrderType),
		},
		Source:       req.Source,
		SupplierCode: req.SupplierCode,
		MerchantCode: req.MerchantCode,
	}

	var payInfo *biz.PayInfo

	switch req.SettleType {
	case v1.SettleType_SettleTypeUnknown:
		return nil, biz.ErrSettleTypeIsInvalid
	case v1.SettleType_SettleTypeNormal:
		payInfo, err = uc.createNormal(ctx, order)
	case v1.SettleType_SettleTypeMerge:
		payInfo, err = uc.createMerge(ctx, order)
	case v1.SettleType_SettleTypeGoods:
		payInfo, err = uc.createGoods(ctx, order)
	case v1.SettleType_SettleTypeSublet:
		payInfo, err = uc.createSublet(ctx, order)
	default:
		return nil, biz.ErrSettleTypeIsInvalid
	}

	if err != nil {
		return nil, err
	}

	if err = uc.repo.UpdatePayID(ctx, order.BizType, order.ID, order.SubID, payInfo.OutTradeID); err != nil {
		return nil, err
	}

	return &v1.CreateOrderResponse{
		PayInfo: &v1.PayInfo{
			OutTradeId: payInfo.OutTradeID,
		},
	}, nil
}
