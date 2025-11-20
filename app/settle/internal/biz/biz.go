package biz

import (
	"context"
)

type OrderRepo interface {
	GetOneByID(ctx context.Context, bizType int8, id, subID int64, opts ...Option) (*Order, error)
	GetAllByID(ctx context.Context, bizType int8, id int64, orderType int8, opts ...Option) ([]*Order, error)
	MultiGetBySubIDs(ctx context.Context, bizType int8, id int64, ids []int64, opts ...Option) ([]*Order, error)
	Create(ctx context.Context, param *OrderCreateParam) error
	UpdatePayID(ctx context.Context, bizType int8, id, subID int64, payID string) error

	CheckRefunded(ctx context.Context, bizType int8, id int64, refundNo string) error
	InsertRefunds(ctx context.Context, params []*Refund) error
}

type DepositRepo interface {
}

type MerchantClient interface {
	GetMerchantInfo(ctx context.Context, source int32, supplierCode string) (*MerchantInfo, error)
}

type CashierClient interface {
	CreateOrder(ctx context.Context, param *CreateOrderParam) (*PayInfo, error)
}

type SubletClient interface {
	GetSubletInfo(ctx context.Context, merchantCode string) (*SubletInfo, error)
}
