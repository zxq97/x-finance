package order

import (
	"github.com/zxq97/x-finance/app/settle/internal/biz"
)

// var _ refund.RefundBiz = (*refundNormal)(nil)

type refundNormal struct {
	repo biz.OrderRepo
}
