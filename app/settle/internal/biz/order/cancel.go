package order

import (
	"context"

	"github.com/zxq97/x-finance/api/settle/service/v1"
)

func (uc *OrderUseCase) Cancel(ctx context.Context, req *v1.CancelRequest) (*v1.CancelResponse, error) {
	return &v1.CancelResponse{}, nil
}
