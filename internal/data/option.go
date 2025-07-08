package data

type Option func(*option)

type option struct {
	filterEmpty     bool
	filterDuplicate bool
	subOrder        bool
	snapshot        bool
	subRefund       bool
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

func withSubOrder() Option {
	return func(o *option) {
		o.subOrder = true
	}
}

func withSnapshot() Option {
	return func(o *option) {
		o.snapshot = true
	}
}

func withSubRefund() Option {
	return func(o *option) {
		o.subRefund = true
	}
}
