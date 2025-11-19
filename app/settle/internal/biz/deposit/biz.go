package deposit

import (
	"github.com/google/wire"
	"github.com/zxq97/x-finance/app/settle/internal/biz"
)

var ProviderSet = wire.NewSet(NewDepositUseCase)

type DepositUseCase struct {
	repo biz.DepositRepo
}

func NewDepositUseCase(repo biz.DepositRepo) *DepositUseCase {
	return &DepositUseCase{repo: repo}
}
