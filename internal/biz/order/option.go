package order

type Option func(*option)

type option struct {
	status          []int8
	typ             []int8
	filterNoBalance bool
	filterEmpty     bool
	filterDuplicate bool
}

func withStatus(status ...int8) Option {
	return func(o *option) {
		o.status = status
	}
}

func withType(typ ...int8) Option {
	return func(o *option) {
		o.typ = typ
	}
}

func withFilterNoBalance() Option {
	return func(o *option) {
		o.filterNoBalance = true
	}
}

func withFilterEmpty() Option {
	return func(o *option) {
		o.filterEmpty = true
	}
}

func withFilterDuplicate() Option {
	return func(o *option) {
		o.filterDuplicate = true
	}
}
