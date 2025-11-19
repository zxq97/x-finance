package biz

type Deduct struct {
	BizType int8

	MainID     int64
	DeductID   int64
	Amount     int64
	Status     int8
	DeductNo   string
	DeductType int8
}
