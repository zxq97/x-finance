package biz

import (
	"context"

	"github.com/zxq97/x-finance/internal/data"
)

type OrderRepo interface {
	GetOne(ctx context.Context, cond map[string]interface{}) (*data.Order, error)
	GetAll(ctx context.Context, cond map[string]interface{}) ([]*data.Order, error)
	Create(ctx context.Context, order *data.Order) error
	Paid(ctx context.Context, id int64) error
	GetRefundByNo(ctx context.Context, refundNo string) (*data.Refund, error)
}

type DepositRepo interface {
	Create(ctx context.Context) error
}

type Calculate interface {
	Refund(ctx context.Context, itmes []*RefundItemDO) (map[int64]*RefundResultDO, error)
}
