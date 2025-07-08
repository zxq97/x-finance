package calculate

import (
	"github.com/zxq97/x-finance/internal/biz"
	"github.com/zxq97/x-finance/internal/util"
)

type common struct{}

func (c *common) checkAndCalc(param *biz.RefundAmtParam) (*refundAmtResult, error) {
	if len(param.RefundItems) == 0 {
		return nil, biz.ErrRefundAmtInvalid
	}

	var balance, discount, settle, target, realPay int64
	for _, item := range param.RefundItems {
		balance += item.Balance
		discount += item.Discount
		settle += item.Balance + item.Discount
		realPay += item.RealPay
		target += item.RealPay + item.Coupon
	}

	if settle < param.RefundAmt {
		return nil, biz.ErrRefundAmtInvalid
	}

	res := new(refundAmtResult)
	switch param.RefundType {
	case 0, 4, 6: // 钱优先
		res.refundReal = util.Min(param.RefundAmt, balance)
		res.refundCoupon = util.Min(param.RefundAmt-res.refundReal, discount)
	case 1: // 券优先
		res.refundCoupon = util.Min(param.RefundAmt, discount)
		res.refundReal = util.Min(param.RefundAmt-res.refundCoupon, balance)
	case 2: // 比例
		res.refundReal = util.Min(realPay/target*param.RefundAmt, balance)
		res.refundCoupon = util.Min(param.RefundAmt-res.refundReal, discount)
	default:
		return nil, biz.ErrRefundTypeNotFound
	}

	if res.refundReal > balance || res.refundCoupon > discount {
		return nil, biz.ErrRefundAmtInvalid
	}

	return res, nil
}
