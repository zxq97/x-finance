package order

import (
	"context"

	"github.com/google/wire"
	"github.com/zxq97/x-finance/internal/biz"
	"github.com/zxq97/x-finance/internal/data"
)

var ProviderSet = wire.NewSet(NewOrderUseCase)

type OrderUseCase struct {
	repo biz.OrderRepo
}

func NewOrderUseCase(repo biz.OrderRepo) *OrderUseCase {
	return &OrderUseCase{repo: repo}
}

func (uc *OrderUseCase) getWithOption(ctx context.Context, id int64, opts ...Option) (*data.Order, error) {
	opt := &option{}

	for _, o := range opts {
		o(opt)
	}

	cond := map[string]interface{}{
		"id": id,
	}

	if len(opt.status) != 0 {
		cond["status in (?)"] = opt.status
	}

	if len(opt.typ) != 0 {
		cond["type in (?)"] = opt.typ
	}

	if opt.filterNoBalance {
		cond["balance > ?"] = 0
	}

	order, err := uc.repo.GetOne(ctx, cond)
	if err != nil {
		return nil, err
	} else if order == nil && opt.filterEmpty {
		return nil, biz.ErrOrderNotFound
	} else if order != nil && opt.filterDuplicate {
		return nil, biz.ErrOrderDuplicate
	}

	return order, nil
}

func (uc *OrderUseCase) getAllWtihOption(ctx context.Context, mainID int64, opts ...Option) ([]*data.Order, error) {
	opt := &option{}

	for _, o := range opts {
		o(opt)
	}

	cond := map[string]interface{}{
		"main_id": mainID,
	}

	if len(opt.status) != 0 {
		cond["status in (?)"] = opt.status
	}

	if len(opt.typ) != 0 {
		cond["type in (?)"] = opt.typ
	}

	if opt.filterNoBalance {
		cond["balance > ?"] = 0
	}

	orders, err := uc.repo.GetAll(ctx, cond)
	if err != nil {
		return nil, err
	}

	if opt.filterEmpty && len(orders) == 0 {
		return nil, biz.ErrOrderNotFound
	}

	return orders, nil
}
