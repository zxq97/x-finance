package refund

import "sort"

var _ sort.Interface = (*sortByBCP)(nil)
var _ sort.Interface = (*sortByBPC)(nil)
var _ sort.Interface = (*sortByPCB)(nil)

type baseSorter struct {
	orders []*Order
}

func (s baseSorter) Len() int {
	return len(s.orders)
}

func (s baseSorter) Swap(i, j int) {
	s.orders[i], s.orders[j] = s.orders[j], s.orders[i]
}

type sortByBCP struct {
	baseSorter
}
func (s sortByBCP) Less(i, j int) bool {
	if s.orders[i].Balance == s.orders[j].Balance {
		if s.orders[i].CouponAmt == s.orders[j].CouponAmt {
			return s.orders[i].PromoAmt > s.orders[j].PromoAmt
		}

		return s.orders[i].CouponAmt > s.orders[j].CouponAmt
	}

	return s.orders[i].Balance > s.orders[j].Balance
}

type sortByBPC struct {
	baseSorter
}

func (s sortByBPC) Less(i, j int) bool {
	if s.orders[i].Balance == s.orders[j].Balance {
		if s.orders[i].PromoAmt == s.orders[j].PromoAmt {
			return s.orders[i].CouponAmt > s.orders[j].CouponAmt
		}

		return s.orders[i].PromoAmt > s.orders[j].PromoAmt
	}

	return s.orders[i].Balance > s.orders[j].Balance
}

type sortByPCB struct {
	baseSorter
}

func (s sortByPCB) Less(i, j int) bool {
	if s.orders[i].PromoAmt == s.orders[j].PromoAmt {
		if s.orders[i].CouponAmt == s.orders[j].CouponAmt {
			return s.orders[i].Balance > s.orders[j].Balance
		}
		return s.orders[i].CouponAmt > s.orders[j].CouponAmt
	}
	return s.orders[i].PromoAmt > s.orders[j].PromoAmt
}

func SortByBCP(orders []*Order) {
    sorter := sortByBCP{baseSorter{orders}}
    sort.Sort(sorter)
}

func SortByBPC(orders []*Order) {
    sorter := sortByBPC{baseSorter{orders}}
    sort.Sort(sorter)
}

func SortByPCB(orders []*Order) {
    sorter := sortByPCB{baseSorter{orders}}
    sort.Sort(sorter)
}
