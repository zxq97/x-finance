package calculate

import (
	"context"
	"sort"

	"github.com/zxq97/x-finance/internal/biz"
	"github.com/zxq97/x-finance/internal/util"
)

var _ biz.Calculate = (*normal)(nil)

type normal struct {
	*common
}

func (n *normal) Refund(ctx context.Context, parma *biz.RefundAmtParam) (map[int64]*biz.RefundAmtResult, error) {
	res, err := n.checkAndCalc(parma)
	if err != nil {
		return nil, err
	}

	m := make(map[int64]*biz.RefundAmtResult, len(parma.RefundItems))

	sort.Slice(parma.RefundItems, func(i, j int) bool {
		return parma.RefundItems[i].OrderType < parma.RefundItems[j].OrderType
	})

	var tmpRefundReal, tmpRefundCoupon int64
	for _, v := range parma.RefundItems {
		if res.refundReal == 0 && res.refundCoupon == 0 {
			break
		}

		tmpRefundReal = util.Min(res.refundReal, v.Balance)
		tmpRefundCoupon = util.Min(res.refundCoupon, v.Discount)
		if tmpRefundReal == 0 && tmpRefundCoupon == 0 {
			continue
		}

		res.refundReal -= tmpRefundReal
		res.refundCoupon -= tmpRefundCoupon

		m[v.ID] = &biz.RefundAmtResult{
			RefundReal:   tmpRefundReal,
			RefundCoupon: tmpRefundCoupon,
		}
	}

	if res.refundReal > 0 || res.refundCoupon > 0 {
		return nil, biz.ErrRefundAmtInvalid
	}

	return m, nil
}

func (n *normal) Profit(ctx context.Context, param *biz.ProfitAmtParam) (*biz.ProfitAmtResult, error) {
	res := new(biz.ProfitAmtResult)

	res.Profit = param.Profit - param.RefundAmt
	res.Settle = param.Settlt - param.RefundAmt
	res.SelfProfit = res.Profit * param.Rate / 100
	res.OtherProfit = res.Profit - res.SelfProfit
	res.SelfRefund = res.SelfProfit - param.SelfProfit
	res.OtherRefund = param.RefundAmt - res.SelfRefund
	res.SelfSettle = res.SelfSettle - res.SelfRefund
	res.OtherSettle = res.OtherSettle - res.OtherRefund

	if res.SelfSettle < 0 || res.OtherSettle < 0 {
		return nil, biz.ErrCancelDuplicate
	}

	return res, nil
}
