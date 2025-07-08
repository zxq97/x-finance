package data

import (
	"context"

	"github.com/pkg/errors"
	"github.com/zxq97/x-finance/internal/biz"
	"gorm.io/gorm"
)

func (r *repo) getOneOrderWithOption(ctx context.Context, mainID, orderID int64, opts ...Option) error {
	opt := &option{}
	for _, o := range opts {
		o(opt)
	}

	m := &MainOrder{}
	if err := r.db.Select("*").Table(mainOrder).Where("main_id = ?", mainID).First(m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			if opt.filterEmpty {
				return biz.ErrOrderNotFound
			}
		}
		return err
	}

	return nil
}

func (r *repo) CheckCreate(ctx context.Context, mainID, id int64, orderType int8) error {
	return nil
}

func (r * repo) CheckRefund(ctx context.Context, mainID int64, refundNo string) error {
	panic(1)
}

func (r *repo) Create(ctx context.Context, param *biz.Order) error {
	panic(1)
}

func (r *repo) GetSubOrderByID(ctx context.Context, id int64) (*biz.Order, error) {
	panic(1)
}
func (r *repo) GetSubOrdersByMainID(ctx context.Context, mainID int64, orderType int8) ([]*biz.Order, error) {
	panic(1)
}

func (r *repo) Refund(ctx context.Context, param *biz.RefundResult) error {
	panic(1)
}

func (r *repo) UpdateOrderStatus(ctx context.Context, id int64, status int8) error {
	panic(1)
}