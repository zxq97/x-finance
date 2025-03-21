package order

import (
	"context"

	"github.com/zxq97/x-finance/internal/biz"
	"github.com/zxq97/x-finance/internal/data"
)

func (uc *OrderUseCase) build(ctx context.Context, do *OrderDO) (*data.Order, error) {
	return &data.Order{
		ID: do.ID,
	}, nil
}

func (uc *OrderUseCase) checkPayStatus(ctx context.Context, do *NotifyDO) error {
	if do.Status != StatusPaid {
		return biz.ErrOrderPayFailed
	}
	return nil
}

func (uc *OrderUseCase) Create(ctx context.Context, do *OrderDO) error {
	// fixme 改成option形式 携带option 直接判断err
	o, err := uc.repo.GetOne(ctx, nil)
	if err != nil {
		return err
	} else if o != nil {
		return biz.ErrOrderDuplicate
	}

	order, err := uc.build(ctx, do)
	if err != nil {
		return err
	}

	return uc.repo.Create(ctx, order)
}

func (uc *OrderUseCase) NotifyPay(ctx context.Context, do *NotifyDO) error {
	// fixme 改成option形式 携带option 直接判断err
	_, err := uc.repo.GetOne(ctx, nil)
	if err != nil {
		return err
	}

	if err = uc.checkPayStatus(ctx, do); err!= nil {
		return err
	}

	return uc.repo.Paid(ctx, do.ID)
}
