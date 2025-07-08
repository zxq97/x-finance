package biz

type RefundAmtParam struct {
	RefundAmt   int64
	RefundType  int8
	RefundItems []*RefundItem
}

type RefundItem struct {
	ID        int64
	OrderType int8
	Balance   int64
	Discount  int64
	RealPay   int64
	Coupon    int64
}

type RefundAmtResult struct {
	RefundReal   int64
	RefundCoupon int64
}

type ProfitAmtParam struct {
	Rate       int64
	RefundAmt  int64
	Profit     int64
	Settlt     int64
	SelfProfit int64
}

type ProfitAmtResult struct {
	Settle      int64
	Profit      int64
	SelfSettle  int64
	OtherSettle int64
	SelfProfit  int64
	OtherProfit int64
	SelfRefund  int64
	OtherRefund int64
}

type RefundResult struct {
	MainID     int64
	RefundType int8
	RefundAmt  int64
	RefundNo   string
	Items      []*RefundResultItem
}

type RefundResultItem struct {
	ID           int64
	RefundReal   int64
	RefundCoupon int64
	SelfRefund   int64
	OtherRefund  int64
}
