package deposit

const (
	deducting = 1
	deducted = 2
	deductFailed = 3
	deductCanceled = 4
)

type DeductParam struct {
	BizType int8

	MainID     int64
	DeductID   int64
	Amount     int64
	Status     int8
	DeductNo   string
	DeductType int8
}