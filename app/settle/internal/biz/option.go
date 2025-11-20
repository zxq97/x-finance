package biz

type Option func(o *Opt)

type Opt struct {
	WithDynamic bool
}

func WithDynamic() Option {
	return func(o *Opt) {
		o.WithDynamic = true
	}
}
