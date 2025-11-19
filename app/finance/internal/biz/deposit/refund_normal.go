package deposit

import (
	"context"
	"github.com/zxq97/x-finance/app/finance/internal/biz"
	"github.com/zxq97/x-finance/pkg/refund"
)

var _ refund.RefundBiz = (*refundNormal)(nil)

type refundNormal struct {
	repo biz.DepositRepo
}

func (r *refundNormal) CheckIdempotent(ctx context.Context, param *refund.CheckIdempotentParam) error {
	if err := r.repo.HasDeductingAndDeductFailed(ctx, param.MainID); err != nil {
		return err
	}

	return r.repo.CheckRefund(ctx, param.MainID, param.RefundNo)
}

func (r *refundNormal) GetNeedRefundOrders(ctx context.Context, param *refund.GetRefundOrderParam) ([]*refund.Order, error) {
	deducts, err := r.repo.GetDeductsByMainID(ctx, param.MainID, 1)
	if err != nil {
		return nil, err
	}

	res := make([]*refund.Order, 0, len(deducts))
	ids := make([]int64, 0, len(deducts))
	m := make(map[int64]*refund.Order, len(deducts))
	for _, v := range deducts {
		m[v.DeductID] = &refund.Order{
			BizType:   v.BizType,
			MainID:    v.MainID,
			ID:        v.DeductID,
			OrderType: v.DeductType,
			TargetAmt: v.Amount,
		}
		res = append(res, m[v.DeductID])
		ids = append(ids, v.DeductID)
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
	if param.RefundAmt != param.RefundReal || param.RefundCoupon != 0 || param.RefundPromo != 0 {
		return nil, biz.ErrRefundAmtInvalid
	}

	var amt, balance int64
	for _, v := range param.Orders {
		amt += v.SettleAmt
		balance += v.Balance

	}

	if amt < param.RefundAmt || balance < param.RefundReal {
		return nil, biz.ErrRefundAmtInvalid
	}

	return &refund.RefundDetail{RealAmt: param.RefundAmt}, nil
}

func (r *refundNormal) InsertRefundRecords(ctx context.Context, records []*refund.RefundRecord) error {
	panic("unimplemented")
}

func (r *refundNormal) RealRefund(ctx context.Context, record *refund.RefundRecord) error {
	panic("unimplemented")
}
