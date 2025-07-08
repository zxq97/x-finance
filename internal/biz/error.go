package biz

import "github.com/pkg/errors"

var (
	ErrOrderPayFailed = errors.New("order: pay failed")
	ErrOrderNotFound  = errors.New("order: not found")
	ErrOrderDuplicate = errors.New("order: duplicate")

	ErrCancelDuplicate = errors.New("cancel: duplicate")

	ErrStrategyNotFound = errors.New("strategy: not found")

	ErrRefundTypeNotFound = errors.New("refund: type not found")
	ErrRefundAmtInvalid   = errors.New("refund: amount invalid")
	ErrRefundNoIdempotent = errors.New("refund: no idempotent")

	ErrCancelDeductInvalid = errors.New("cancel: deduct invalid")

	ErrProfitSettleInvalid = errors.New("profit: settle invalid")
)
