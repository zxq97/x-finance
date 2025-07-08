package biz

import (
	"context"
)

type OrderRepo interface {
	CheckCreate(ctx context.Context, mainID, id int64, orderType int8) error
	CheckRefund(ctx context.Context, mainID int64, refundNo string) error
	Create(ctx context.Context, param *Order) error
	GetSubOrderByID(ctx context.Context, id int64) (*Order, error)
	GetSubOrdersByMainID(ctx context.Context, mainID int64, orderType int8) ([]*Order, error)
	Refund(ctx context.Context, param *RefundResult) error

	UpdateOrderStatus(ctx context.Context, id int64, status int8) error
}

type DepositRepo interface {
	Create(ctx context.Context) error
}

type Calculate interface {
	Refund(ctx context.Context, param *RefundAmtParam) (map[int64]*RefundAmtResult, error)
	Profit(ctx context.Context, param *ProfitAmtParam) (*ProfitAmtResult, error)
}
