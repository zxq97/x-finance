package calculate

import (
	"context"

	"github.com/zxq97/x-finance/internal/biz"
)

var _ biz.Calculate = (*normal)(nil)

type normal struct{}

func (*normal) Refund(ctx context.Context, itmes []*biz.RefundItemDO) (map[int64]*biz.RefundResultDO, error) {
	return nil, nil
}
