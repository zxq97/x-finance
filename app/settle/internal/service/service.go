package service

import (
	"context"
	"github.com/google/wire"
	"github.com/zxq97/x-finance/api/settle/service/v1"
	"github.com/zxq97/x-finance/app/settle/internal/biz/deposit"
	"github.com/zxq97/x-finance/app/settle/internal/biz/order"
	"google.golang.org/protobuf/types/known/emptypb"
)

var ProviderSet = wire.NewSet(NewSettleService)

type SettleService struct {
	v1.UnimplementedSettleSvcServer
	order   *order.OrderUseCase
	deposit *deposit.DepositUseCase
}

// OrderCancel implements v1.SettleSvcServer.
func (s *SettleService) OrderCancel(ctx context.Context, req *v1.CancelRequest) (*v1.CancelResponse, error) {
	panic("unimplemented")
}

// OrderCreate implements v1.SettleSvcServer.
func (s *SettleService) OrderCreate(ctx context.Context, req *v1.CreateOrderRequest) (*v1.CreateOrderResponse, error) {
	panic("unimplemented")
}

// OrderRefund implements v1.SettleSvcServer.
func (s *SettleService) OrderRefund(ctx context.Context, req *v1.RefundRequest) (*v1.RefundResponse, error) {
	panic("unimplemented")
}

// OrderSettle implements v1.SettleSvcServer.
func (s *SettleService) OrderSettle(ctx context.Context, req *v1.SettleRequest) (*emptypb.Empty, error) {
	panic("unimplemented")
}

func NewSettleService(orderUseCase *order.OrderUseCase, depositUseCase *deposit.DepositUseCase) *SettleService {
	return &SettleService{
		order:   orderUseCase,
		deposit: depositUseCase,
	}
}
