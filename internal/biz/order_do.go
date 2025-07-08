package biz

type Order struct {
	BizType int8

	MainID      int64
	ID          int64
	Type        int8
	Status      int8

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
	Rate       int8
	Commission bool
}
