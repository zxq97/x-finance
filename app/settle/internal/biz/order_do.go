package biz

type SettleDetail struct {
	SubID int64
	TargetType int8
	MerchantID string
	SettleAmt int64
}

type Order struct {
	BizType int8
	UID int64
	ID int64
	SubID int64
	OrderType int8

	StoreAmt int64
	TargetAmt int64
	RealPayAmt int64
	CouponAmt int64
	PromotionAmt int64
	MchDiscountAmt int64

	SettleAmt int64
	Balance int64
	Discount int64
	Promotion int64

	SettleDetails []*SettleDetail
}
