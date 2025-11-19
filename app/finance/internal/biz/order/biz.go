package order

import (
	"github.com/google/wire"
	"github.com/zxq97/x-finance/app/finance/internal/biz"
)

var ProviderSet = wire.NewSet(NewOrderUseCase)

var normal *refundNormal

type OrderUseCase struct {
	repo biz.OrderRepo
}

func NewOrderUseCase(repo biz.OrderRepo) *OrderUseCase {
	normal = &refundNormal{repo: repo}
	return &OrderUseCase{repo: repo}
}
