package order

import (
	"context"
	"fmt"

	"github.com/zxq97/x-finance/internal/biz"
	"github.com/zxq97/x-finance/internal/biz/calculate"
)

func (uc *OrderUseCase) Cancel(ctx context.Context, param *CancelParam) error {
	refundNo := fmt.Sprintf(refundNoTmp, param.ID, param.OrderType)
	if err := uc.repo.CheckRefund(ctx, param.MainID, refundNo); err != nil {
		return err
	}

	orders, err := uc.repo.GetSubOrdersByMainID(ctx, param.MainID, param.OrderType)
	if err != nil {
		return err
	}

	var sumSettle int64
	list := make([]*biz.RefundItem, len(orders))
	for k, v := range orders {
		sumSettle += v.Settle
		list[k] = &biz.RefundItem{
			ID:        v.ID,
			OrderType: v.Type,
			Balance:   v.Balance,
			Discount:  v.Discount,
			RealPay:   v.RealPay,
			Coupon:    v.Coupon,
		}
	}

	if sumSettle < param.Deduct {
		return biz.ErrCancelDeductInvalid
	}

	calc, err := calculate.NewCalculate(calculate.StrategyNormal)
	if err != nil {
		return err
	}

	uapAmt, err := calc.Refund(ctx, &biz.RefundAmtParam{
		RefundAmt:   sumSettle - param.Deduct,
		RefundType:  RefundTypeOrderCancel,
		RefundItems: list,
	})
	if err != nil {
		return err
	}

	res := &biz.RefundResult{
		MainID:     param.MainID,
		RefundType: RefundTypeOrderCancel,
		RefundAmt:  sumSettle - param.Deduct,
		RefundNo:   refundNo,
		Items:      make([]*biz.RefundResultItem, 0, len(orders)),
	}
	for _, v := range orders {
		if val, ok := uapAmt[v.ID]; ok {
			pasAmt, err := calc.Profit(ctx, &biz.ProfitAmtParam{
				Profit:     v.Profit,
				SelfProfit: v.SelfPrifit,
				RefundAmt:  val.RefundCoupon + val.RefundReal,
				Settlt:     v.Settle,
				Rate:       10,
			})
			if err != nil {
				return err
			}

			res.Items = append(res.Items, &biz.RefundResultItem{
				ID:           v.ID,
				RefundReal:   val.RefundReal,
				RefundCoupon: val.RefundCoupon,
				SelfRefund:   pasAmt.SelfRefund,
				OtherRefund:  pasAmt.OtherRefund,
			})
		}
	}

	return uc.repo.Refund(ctx, res)
}
