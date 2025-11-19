package service

import (
	"github.com/google/wire"
	"github.com/zxq97/x-finance/app/settle/internal/biz/deposit"
	"github.com/zxq97/x-finance/app/settle/internal/biz/order"
)

var ProviderSet = wire.NewSet(NewSettleService)

type SettleService struct {
	order   *order.OrderUseCase
	deposit *deposit.DepositUseCase
}

func NewSettleService(orderUseCase *order.OrderUseCase, depositUseCase *deposit.DepositUseCase) *SettleService {
	return &SettleService{
		order:   orderUseCase,
		deposit: depositUseCase,
	}
}
