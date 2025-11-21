package biz

type Option func(o *Opt)

type Opt struct {
	WithDynamic bool
	OrderType   []int8
}

func WithDynamic() Option {
	return func(o *Opt) {
		o.WithDynamic = true
	}
}

func WithOrderType(types []int8) Option {
	return func(o *Opt) {
		o.OrderType = append(o.OrderType, types...)
	}
}
