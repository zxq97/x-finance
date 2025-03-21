package biz

type RefundItemDO struct {
	MainID int64
	ID     int64
}

type RefundResultDO struct {
	RefundRealPay     int64
	RefundCoupon      int64
	RefundSettle      int64
	RefundProfit      int64
	SelfSettleRefund  int64
	OtherSettleRefund int64
	SelfProfitRefund  int64
	OtherProfitRefund int64
}
