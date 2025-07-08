package order

import (
	"context"

	"github.com/zxq97/x-finance/internal/biz"
	"github.com/zxq97/x-finance/internal/biz/calculate"
)

func (uc *OrderUseCase) Refund(ctx context.Context, param *RefundParam) error {
	if err := uc.repo.CheckRefund(ctx, param.MainID, param.RefundNo); err != nil {
		return err
	}

	orders, err := uc.repo.GetSubOrdersByMainID(ctx, param.MainID, OrderTypeAll)
	if err != nil {
		return err
	}

	calc, err := calculate.NewCalculate(calculate.StrategyNormal)
	if err != nil {
		return err
	}

	list := make([]*biz.RefundItem, len(orders))
	for k, v := range orders {
		list[k] = &biz.RefundItem{
			ID:        v.ID,
			OrderType: v.Type,
			Balance:   v.Balance,
			Discount:  v.Discount,
			RealPay:   v.RealPay,
			Coupon:    v.Coupon,
		}
	}
	uapAmt, err := calc.Refund(ctx, &biz.RefundAmtParam{
		RefundAmt:   param.RefundAmt,
		RefundType:  param.RefundType,
		RefundItems: list,
	})
	if err != nil {
		return err
	}

	res := &biz.RefundResult{
		MainID:     param.MainID,
		RefundType: param.RefundType,
		RefundAmt:  param.RefundAmt,
		RefundNo:   param.RefundNo,
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
