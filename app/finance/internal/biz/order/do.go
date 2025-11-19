package order

const (
	StatusUnpaid   = 0
	StatusPaid     = 1
	StatusFailed   = 2
	StatusCanceled = 3

	OrderTypeAll    = 0
	OrderTypeNormal = 1
	OrderTypeModify = 2

	RefundTypeOrderCancel = 1

	refundNoTmp = "refund_no_%d_%d_cancel" // id, typ
)

type NotifyPayParam struct {
	ID        int64
	OrderType int8
	Status    int8
}

type CreateParam struct {
	MainID    int64
	ID        int64
	OrderType int8
	Target    int64
	RealPay   int64
	Coupon    int64
	BizType   int8
}

type CancelParam struct {
	MainID    int64
	ID        int64
	OrderType int8
	Deduct    int64
}

type RefundParam struct {
	MainID     int64
	RefundType int8
	RefundAmt  int64
	RefundNo   string
	BizType    int8
}
