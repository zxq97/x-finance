package deposit

import (
	"github.com/google/wire"
	"github.com/zxq97/x-finance/internal/biz"
)

var ProviderSet = wire.NewSet(NewDepositUseCase)

var normal *refundNormal

type DepositUseCase struct {
	repo      biz.DepositRepo
	orderRepo biz.OrderRepo
}

func NewDepositUseCase(repo biz.DepositRepo, orderRepo biz.OrderRepo) *DepositUseCase {
	normal = &refundNormal{repo: repo}
	return &DepositUseCase{repo: repo, orderRepo: orderRepo}
}
