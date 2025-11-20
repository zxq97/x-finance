package order

import (
	"github.com/google/wire"
	"github.com/zxq97/x-finance/app/settle/internal/biz"
)

var ProviderSet = wire.NewSet(NewOrderUseCase)

const (
	targetTypePlatform = 0
	targetTypeMerchant = 1
	targetTypeThird    = 2
	targetTypeSublet   = 3

	platformMerchantID = ""
)

type OrderUseCase struct {
	repo     biz.OrderRepo
	merchant biz.MerchantClient
	cashier  biz.CashierClient
	sublet   biz.SubletClient
}

var normal *refundNormal

func NewOrderUseCase(repo biz.OrderRepo, merchant biz.MerchantClient, cashier biz.CashierClient) *OrderUseCase {
	normal = &refundNormal{repo: repo, cashier: cashier}
	return &OrderUseCase{repo: repo, merchant: merchant, cashier: cashier}
}
