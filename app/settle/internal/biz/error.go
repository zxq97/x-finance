package biz

import "github.com/pkg/errors"

var (
	ErrBizTypeIsInvalid      = errors.New("biz: type is invalid")
	ErrOrderTypeIsInvalid    = errors.New("order: type is invalid")
	ErrOrderIDIsInvalid      = errors.New("order: id is invalid")
	ErrOrderIDIsDuplicate    = errors.New("order: id is duplicate")
	ErrSettleTypeIsInvalid   = errors.New("settle: type is invalid")
	ErrSettleDetailNotFound  = errors.New("settle: detail not found")
	ErrMerchantCodeIsInvalid = errors.New("merchant: code is invalid")
	ErrRefundAmtInvalid      = errors.New("refund: amount is invalid")
	ErrRefundTypeNotFound    = errors.New("refund: type not found")
)
