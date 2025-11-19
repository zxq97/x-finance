package refund

type RefundParam struct {
	BizType      int8
	MainID       int64
	ID           int64
	Type         int8
	RefundType   int8
	RefundAmt    int64
	RefundReal   int64
	RefundCoupon int64
	RefundPromo  int64
	RefundNo     string
	DeductNo     string
}

type CheckIdempotentParam struct {
	BizType  int8
	MainID   int64
	RefundNo string
}

type GetRefundOrderParam struct {
	BizType  int8
	MainID   int64
	ID       int64
	Type     int8
	DeductNo string
}

type Order struct {
	BizType   int8
	MainID    int64
	ID        int64
	OrderType int8
	TargetAmt int64
	RealAmt   int64
	CouponAmt int64
	PromoAmt  int64
	SettleAmt int64
	Balance   int64
	Discount  int64
	Promotion int64
}

type CheckRefundAmtParam struct {
	Orders       []*Order
	RefundType   int8
	RefundAmt    int64
	RefundReal   int64
	RefundCoupon int64
	RefundPromo  int64
}

type RefundRecord struct {
	BizType       int8
	MainID        int64
	ID            int64
	Type          int8
	RefundType    int8
	RefundNo      string
	ThirdRefundNo string
	DeductAmt     int64
	RefundAmt     int64
	RefundReal    int64
	RefundCoupon  int64
	RefundPromo   int64
}

type RefundDetail struct {
	RealAmt   int64
	CouponAmt int64
	PromoAmt  int64
}

type UpdateRefundStatusParam struct {
	BizType  int8
	MainID   int64
	RefundNo string
	Status   int8
}
