package order

import (
	"context"

	"github.com/zxq97/x-finance/api/settle/service/v1"
	"github.com/zxq97/x-finance/pkg/refund"
)

func (uc *OrderUseCase) Refund(ctx context.Context, req *v1.RefundRequest) (*v1.RefundResponse, error) {
	if err := refund.NewStrategy(normal).Refund(ctx, &refund.RefundParam{}); err != nil {
		return nil, err
	}

	return &v1.RefundResponse{}, nil
}
