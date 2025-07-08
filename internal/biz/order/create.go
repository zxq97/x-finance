package order

import (
	"context"

	"github.com/zxq97/x-finance/internal/biz"
)

func (uc *OrderUseCase) buildOrder(param *CreateParam) *biz.Order {
	var rate int64 = 15

	order := &biz.Order{
		BizType:  param.BizType,
		MainID:   param.MainID,
		ID:       param.ID,
		Type:     param.OrderType,
		Amount:   param.Target,
		RealPay:  param.RealPay,
		Coupon:   param.Coupon,
		Balance:  param.RealPay,
		Discount: param.Coupon,
		Settle:   param.Target,
		Profit:   param.Target,
		Rate:     int8(rate),
	}

	order.SelfSettle = param.Target * rate / 100
	order.OtherSettle = param.Target - order.SelfSettle
	order.SelfPrifit = param.Target * rate / 100
	order.OtherProfit = param.Target - order.SelfPrifit
	return order
}

func (uc *OrderUseCase) Create(ctx context.Context, param *CreateParam) error {
	if err := uc.repo.CheckCreate(ctx, param.MainID, param.ID, param.OrderType); err != nil {
		return err
	}

	return uc.repo.Create(ctx, uc.buildOrder(param))
}

func (uc *OrderUseCase) NotifyPay(ctx context.Context, param *NotifyPayParam) error {
	order, err := uc.repo.GetSubOrderByID(ctx, param.ID)
	if err != nil {
		return err
	}

	if order.Status == param.Status {
		return nil
	}

	// todo -> ch notify

	return uc.repo.UpdateOrderStatus(ctx, param.ID, param.Status)
}
