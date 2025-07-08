package calculate

import (
	"context"

	"github.com/zxq97/x-finance/internal/biz"
)

var _ biz.Calculate = (*without)(nil)

type without struct{}

func (*without) Refund(ctx context.Context, param *biz.RefundAmtParam) (map[int64]*biz.RefundAmtResult, error) {
	return nil, nil
}

func (*without) Profit(ctx context.Context, param *biz.ProfitAmtParam) (*biz.ProfitAmtResult, error) {
	return nil, nil
}