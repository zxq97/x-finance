package order

const (
	StatusUnpaid   = 0
	StatusPaid     = 1
	StatusFailed   = 2
	StatusCanceled = 3

	OrderTypeNormal = 1
	OrderTypeModify = 2

	refundNo = "refund_no_%d_%d_cancel" // id, typ
)

type OrderDO struct {
	MainID      int64
	ID          int64
	Type        int8
	Amount      int64
	RealPay     int64
	Coupon      int64
	Balance     int64
	Discount    int64
	Settle      int64
	Profit      int64
	SelfSettle  int64
	OtherSettle int64
	SelfPrifit  int64
	OtherProfit int64

	// shapshot
	Rate       float64
	Commission bool
}

type NotifyDO struct {
	ID     int64
	Status int8
}

type CancelDO struct {
	MainID    int64
	ID        int64
	Type      int8
	Deduct    int64
	Reason    string
	IsRecycle bool
}

type RefundDO struct {
	MainID        int64
	ID            int64
	Type          int8
	Deduct        int64
	RefundRealPay int64
	RefundCoupon  int64
	RefundSettle  int64
	RefundProfit  int64
}
