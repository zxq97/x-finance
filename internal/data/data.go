package data

import (
	"net/http"

	"github.com/google/wire"
	"github.com/zxq97/x-finance/internal/biz"
	"gorm.io/gorm"
)

var ProviderSet = wire.NewSet()

var _ biz.OrderRepo = (*repo)(nil)

type repo struct {
	client *http.Client
	db     *gorm.DB
}

func NewRepo(client *http.Client, db *gorm.DB) *repo {
	return &repo{}
}

const (
	mainOrder     = "t_main_order"
	subOrder      = "t_sub_order"
	orderSnapshot = "t_order_snapshot"
	mainRefund    = "t_main_refund"
	subRefund     = "t_sub_refund"
)

type MainOrder struct {
}

type SubOrder struct {
}

type OrderSnapshot struct {
}

type MainRefund struct {
}

type SubRefund struct {
}
