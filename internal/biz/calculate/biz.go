package calculate

import (
	"github.com/zxq97/x-finance/internal/biz"
)

const (
	StrategyNormal  = 1
	StrategyWithout = 2
)

func NewCalculate(strategy int8) (biz.Calculate, error) {
	switch strategy {
	case StrategyNormal:
		return &normal{}, nil
	case StrategyWithout:
		return &without{}, nil
	default:
		return nil, biz.ErrStrategyNotFound
	}
}
