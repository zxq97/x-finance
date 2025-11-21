package refund

import "github.com/pkg/errors"

var (
	ErrRefundOrderNotFound = errors.New("refund: order not found")
)
