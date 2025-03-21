package order

import (
	"context"
	"fmt"

	"github.com/zxq97/x-finance/internal/biz"
	"github.com/zxq97/x-finance/internal/biz/calculate"
	"github.com/zxq97/x-finance/internal/data"
)

func (uc *OrderUseCase) buildRefundNo(ctx context.Context, do *CancelDO) string {
	return fmt.Sprintf(refundNo, do.ID, do.Type)
}

func (uc *OrderUseCase) getStrategy(ctx context.Context, order *data.Order) int8 {
	if order.Commission {
		return calculate.StrategyNormal
	}
	return calculate.StrategyWithout
}

func (*OrderUseCase) transformRefundItem(ctx context.Context, order ...*data.Order) []*biz.RefundItemDO {
	if len(order) == 0 {
		return nil
	}

	items := make([]*biz.RefundItemDO, len(order))
	for k, v := range order {
		items[k] = &biz.RefundItemDO{
			ID: v.ID,
		}
	}
	return items
}

func (uc *OrderUseCase) Cancel(ctx context.Context, do *CancelDO) error {
	order, err := uc.getWithOption(ctx, do.ID, withStatus(StatusUnpaid), withType(OrderTypeNormal), withFilterEmpty(), withFilterNoBalance())
	if err != nil {
		return err
	}

	refundNo := uc.buildRefundNo(ctx, do)
	refund, err := uc.repo.GetRefundByNo(ctx, refundNo)
	if err != nil {
		return err
	} else if refund != nil {
		return biz.ErrCancelDuplicate
	}

	orders, err := uc.getAllWtihOption(ctx, do.MainID, withStatus(StatusPaid), withType(OrderTypeNormal), withFilterNoBalance())
	if err != nil {
		return err
	}

	calc, err := calculate.NewCalculate(uc.getStrategy(ctx, order))
	if err != nil {
		return err
	}

	res, err := calc.Refund(ctx, uc.transformRefundItem(ctx, orders...))
	if err != nil {
		return err
	}

	
}
