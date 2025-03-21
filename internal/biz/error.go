package biz

import "github.com/pkg/errors"

var (
	ErrOrderPayFailed = errors.New("order: pay failed")

	ErrOrderNotFound  = errors.New("order: not found")
	ErrOrderDuplicate = errors.New("order: duplicate")

	ErrCancelDuplicate = errors.New("cancel: duplicate")

	ErrStrategyNotFound = errors.New("strategy: not found")
)
