package refund

import (
	"context"

	"github.com/zxq97/x-finance/internal/util"
)

type Strategy interface {
	Refund(ctx context.Context, param *RefundParam) error
}

type RefundBiz interface {
	CheckIdempotent(ctx context.Context, param *CheckIdempotentParam) error
	GetNeedRefundOrders(ctx context.Context, param *GetRefundOrderParam) ([]*Order, error)
	CheckRefundAmt(ctx context.Context, param *CheckRefundAmtParam) (*RefundDetail, error)
	InsertRefundRecords(ctx context.Context, records []*RefundRecord) error
	RealRefund(ctx context.Context, record *RefundRecord) error
}

var _ Strategy = (*example)(nil)

type example struct	{
	biz RefundBiz
}

func NewStrategy(biz RefundBiz) Strategy {
	return &example{biz: biz}
}

func (e *example) Refund(ctx context.Context, param *RefundParam) error {
	if err := e.biz.CheckIdempotent(ctx, &CheckIdempotentParam{
		BizType:  param.BizType,
		MainID:   param.MainID,
		RefundNo: param.RefundNo,
	}); err != nil {
		return err
	}

	orders, err := e.biz.GetNeedRefundOrders(ctx, &GetRefundOrderParam{
		BizType:  param.BizType,
		MainID:   param.MainID,
		ID:       param.ID,
		Type:     param.Type,
		DeductNo: param.DeductNo,
	})
	if err != nil {
		return err
	} else if len(orders) == 0 {
		return err // error dingyi
	}

	refundDetail, err := e.biz.CheckRefundAmt(ctx, &CheckRefundAmtParam{
		Orders:       orders,
		RefundAmt:    param.RefundAmt,
		RefundReal:   param.RefundReal,
		RefundCoupon: param.RefundCoupon,
		RefundPromo:  param.RefundPromo,
	})
	if err != nil {
		return err
	}

	refundRecords := make([]*RefundRecord, 0, len(orders))
	for _, v := range orders {
		if refundDetail.RealAmt == 0 && refundDetail.CouponAmt == 0 && refundDetail.PromoAmt == 0 {
			break
		}

		tmpRefundReal := util.Min(v.Balance, refundDetail.RealAmt)
		tmpRefundCoupon := util.Min(v.CouponAmt, refundDetail.CouponAmt)
		tmpRefundPromo := util.Min(v.PromoAmt, refundDetail.PromoAmt)

		if tmpRefundReal == 0 && tmpRefundCoupon == 0 && tmpRefundPromo == 0 {
			continue
		}

		refundDetail.RealAmt -= tmpRefundReal
		refundDetail.CouponAmt -= tmpRefundCoupon
		refundDetail.PromoAmt -= tmpRefundPromo

		refundRecords = append(refundRecords, &RefundRecord{
			BizType:       param.BizType,
			MainID:        param.MainID,
			ID:            v.ID,
			Type:          v.OrderType,
			RefundType:    param.RefundType,
			RefundNo:      param.RefundNo,
			RefundAmt:     param.RefundAmt,
			RefundReal:    tmpRefundReal,
			RefundCoupon:  tmpRefundCoupon,
			RefundPromo:   tmpRefundPromo,
		})
	}

	if err = e.biz.InsertRefundRecords(ctx, refundRecords); err != nil {
		return err
	}

	for _, v := range refundRecords {
		_ = e.biz.RealRefund(ctx, v)
	}

	return nil
}
