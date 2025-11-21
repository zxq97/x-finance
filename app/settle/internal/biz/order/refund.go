package order

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"strconv"
	"strings"

	v1 "github.com/zxq97/x-finance/api/settle/service/v1"
	"github.com/zxq97/x-finance/pkg/refund"
)

func buildRefundNo(id int64, thirdRefundNo string) string {
	h := md5.New()
	h.Write([]byte(strconv.FormatInt(id, 10) + thirdRefundNo))
	return strings.ToUpper(hex.EncodeToString(h.Sum(nil)))
}

func (uc *OrderUseCase) Refund(ctx context.Context, req *v1.RefundRequest) (*v1.RefundResponse, error) {
	if err := refund.NewStrategy(normal).Refund(ctx, &refund.RefundParam{}); err != nil {
		return nil, err
	}

	return &v1.RefundResponse{}, nil
}
