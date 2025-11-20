package biz

type RefundDetail struct {
	SubID           int64
	TargetType      int8
	RefundNo        string
	RefundAmt       int64
	MerchantID      string
	GrantMerchantID string
}

type Refund struct {
	BizType       int8
	ID            int64
	SubID         int64
	OrderType     int8
	RefundType    int8
	ThirdRefundNo string

	RefundAmt       int64
	RefundRealPay   int64
	RefundCoupon    int64
	RefundPromotion int64

	RefundDetails []*RefundDetail
}
