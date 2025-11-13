package order

import (
	"context"

	"github.com/zxq97/x-finance/internal/biz"
	"github.com/zxq97/x-finance/internal/pkg/refund"
	"github.com/zxq97/x-finance/internal/util"
)

var _ refund.RefundBiz = (*refundNormal)(nil)

type refundNormal struct {
	repo biz.OrderRepo
}

func (r *refundNormal) CheckIdempotent(ctx context.Context, param *refund.CheckIdempotentParam) error {
	return r.repo.CheckRefund(ctx, param.MainID, param.RefundNo)
}

func (r *refundNormal) GetNeedRefundOrders(ctx context.Context, param *refund.GetRefundOrderParam) ([]*refund.Order, error) {
	var orderType int8
	if param.Type != 0 {
		orderType = param.Type
	}
	orders, err := r.repo.GetSubOrdersByMainID(ctx, param.MainID, orderType)
	if err != nil {
		return nil, err
	}

	res := make([]*refund.Order, 0, len(orders))
	ids := make([]int64, 0, len(orders))
	m := make(map[int64]*refund.Order, len(orders))
	for _, v := range orders {
		m[v.ID] = &refund.Order{
			BizType:   v.BizType,
			MainID:    v.MainID,
			ID:        v.ID,
			OrderType: v.Type,
			TargetAmt: v.Amount,
			RealAmt:   v.RealPay,
			CouponAmt: v.Coupon,
			PromoAmt:  v.Promo,
			SettleAmt: v.Amount,
			Balance:   v.Balance,
			Discount:  v.Discount,
			Promotion: v.Promotion,
		}
		ids = append(ids, v.ID)
		res = append(res, m[v.ID])
	}

	refunds, err := r.repo.GetRefundRecordsByID(ctx, param.MainID, ids)
	if err != nil {
		return nil, err
	}

	for _, v := range refunds {
		if o, ok := m[v.ID]; ok {
			o.SettleAmt -= v.RefundAmt
			o.RealAmt -= v.RefundReal
			o.CouponAmt -= v.RefundCoupon
			o.PromoAmt -= v.RefundPromo
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
	switch param.RefundType {
	case 0, 4, 6: // 钱优先
		res.RealAmt = util.Min(param.RefundAmt, balance)
		res.CouponAmt = util.Min(param.RefundAmt-res.RealAmt, discount)
	case 1: // 券优先
		res.CouponAmt = util.Min(param.RefundAmt, discount)
		res.RealAmt = util.Min(param.RefundAmt-res.CouponAmt, balance)
	case 2: // 比例
		res.RealAmt = util.Min(realPay/target*param.RefundAmt, balance)
		res.CouponAmt = util.Min(param.RefundAmt-res.RealAmt, discount)
	default:
		return nil, biz.ErrRefundTypeNotFound
	}

	return res, nil
}

func (r *refundNormal) InsertRefundRecords(ctx context.Context, records []*refund.RefundRecord) error {
	// return r.repo.InsertRefundRecords(ctx, records)
	return nil
}

func (r *refundNormal) RealRefund(ctx context.Context, record *refund.RefundRecord) error {
	// return r.repo.RealRefund(ctx, record)
	return nil
}
