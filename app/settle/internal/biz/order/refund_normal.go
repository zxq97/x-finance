package order

import (
	"context"

	v1 "github.com/zxq97/x-finance/api/settle/service/v1"
	"github.com/zxq97/x-finance/app/settle/internal/biz"
	"github.com/zxq97/x-finance/pkg/refund"
	"github.com/zxq97/x-finance/pkg/util"
)

var _ refund.RefundBiz = (*refundNormal)(nil)

type refundNormal struct {
	repo    biz.OrderRepo
	cashier biz.CashierClient
}

func (r *refundNormal) CheckIdempotent(ctx context.Context, param *refund.CheckIdempotentParam) error {
	return r.repo.CheckRefunded(ctx, param.BizType, param.MainID, param.RefundNo)
}

func (r *refundNormal) GetNeedRefundOrders(ctx context.Context, param *refund.GetRefundOrderParam) ([]*refund.Order, error) {
	orders, err := r.repo.GetAllByID(ctx, param.BizType, param.MainID, param.Type, biz.WithDynamic())
	if err != nil {
		return nil, err
	}

	res := make([]*refund.Order, len(orders))
	for k, v := range orders {
		res[k] = &refund.Order{
			BizType:   v.BizType,
			MainID:    v.ID,
			ID:        v.SubID,
			OrderType: v.OrderType,
			TargetAmt: v.TargetAmt,
			RealAmt:   v.RealPayAmt,
			CouponAmt: v.CouponAmt,
			PromoAmt:  v.PromotionAmt,
			SettleAmt: v.SettleAmt,
			Balance:   v.Balance,
			Discount:  v.Discount,
			Promotion: v.Promotion,
		}
	}

	return res, nil
}

func (r *refundNormal) CheckRefundAmt(ctx context.Context, param *refund.CheckRefundAmtParam) (*refund.RefundDetail, error) {
	var amt, balance, discount, promotion, realPay, target int64
	for _, v := range param.Orders {
		amt += v.SettleAmt
		balance += v.Balance
		discount += v.Discount
		promotion += v.Promotion
		realPay += v.RealAmt
		target += v.TargetAmt
	}

	if amt < param.RefundAmt || balance < param.RefundReal || discount < param.RefundCoupon || promotion < param.RefundPromo {
		return nil, biz.ErrRefundAmtInvalid
	}

	res := new(refund.RefundDetail)
	switch v1.RefundType(param.RefundType) {
	case v1.RefundType_RefundTypeEnergy: // 钱优先
		res.RealAmt = util.Min(param.RefundAmt, balance)
		res.PromoAmt = util.Min(param.RefundAmt-res.RealAmt, promotion)
		res.CouponAmt = util.Min(param.RefundAmt-res.RealAmt-res.PromoAmt, discount)
	case v1.RefundType_RefundTypePenalty: // 券优先
		res.CouponAmt = util.Min(param.RefundAmt, discount)
		res.PromoAmt = util.Min(param.RefundAmt-res.CouponAmt, promotion)
		res.RealAmt = util.Min(param.RefundAmt-res.CouponAmt-res.PromoAmt, balance)
	case v1.RefundType_RefundTypeUserCancel, v1.RefundType_RefundTypeSource: // 比例
		res.RealAmt = util.Min(realPay/target*param.RefundAmt, balance)
		res.PromoAmt = util.Min(param.RefundAmt-res.RealAmt, promotion)
		res.CouponAmt = util.Min(param.RefundAmt-res.RealAmt-res.PromoAmt, discount)
	default:
		return nil, biz.ErrRefundTypeNotFound
	}

	return res, nil
}

func (r *refundNormal) InsertRefundRecords(ctx context.Context, records []*refund.RefundRecord) error {
	refunds := make([]*biz.Refund, len(records))
	ids := make([]int64, len(records))

	for k, v := range records {
		ids = append(ids, v.ID)
		refunds[k] = &biz.Refund{
			BizType:         v.BizType,
			ID:              v.MainID,
			SubID:           v.ID,
			OrderType:       v.Type,
			RefundType:      v.RefundType,
			ThirdRefundNo:   v.ThirdRefundNo,
			RefundAmt:       v.RefundAmt,
			RefundRealPay:   v.RefundReal,
			RefundCoupon:    v.RefundCoupon,
			RefundPromotion: v.RefundPromo,
		}
	}

	

	return r.repo.InsertRefunds(ctx, refunds)
}

func (r *refundNormal) RealRefund(ctx context.Context, record *refund.RefundRecord) error {
	panic("unimplemented")
}
