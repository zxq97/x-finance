package order

import (
	"github.com/google/wire"
	"github.com/zxq97/x-finance/app/settle/internal/biz"
)

var ProviderSet = wire.NewSet(NewOrderUseCase)

type OrderUseCase struct {
	repo biz.OrderRepo
}

var normal *refundNormal

func NewOrderUseCase(repo biz.OrderRepo) *OrderUseCase {
	normal = &refundNormal{repo: repo}
	return &OrderUseCase{repo: repo}
}
