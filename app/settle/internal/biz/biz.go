package biz

import (
	"context"
)

type OrderRepo interface {
	GetOneByID(ctx context.Context, bizType int8, id, subID int64, opts ...Option) (*Order, error)
	GetAllByID(ctx context.Context, bizType int8, id int64, opts ...Option) ([]*Order, error)
	MultiGetBySubIDs(ctx context.Context, bizType int8, id int64, ids []int64, opts ...Option) ([]*Order, error)
	GetDetailsBySubID(ctx context.Context, subID int64) ([]*SettleDetail, error)
	GetDetailsBySubIDAndType(ctx context.Context, subID int64, targetType int8) ([]*SettleDetail, error)
	MultiGetDetailsBySubIDs(ctx context.Context, subIDs []int64) (map[int64][]*SettleDetail, error)
	MultiGetDetailsBySubIDsAndType(ctx context.Context, subIDs []int64, targetType int8) (map[int64]*SettleDetail, error)
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
