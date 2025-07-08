package order

import (
	"github.com/google/wire"
	"github.com/zxq97/x-finance/internal/biz"
)

var ProviderSet = wire.NewSet(NewOrderUseCase)

type OrderUseCase struct {
	repo biz.OrderRepo
}

func NewOrderUseCase(repo biz.OrderRepo) *OrderUseCase {
	return &OrderUseCase{repo: repo}
}
