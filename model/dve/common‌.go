package dve

type Collection[T DveInfo] []T

func (c Collection[T]) Len() int {
	return len(c)
}

func (c Collection[T]) IsEmpty() bool {
	return len(c) == 0
}

func (c Collection[T]) IsSn(sn string) (bool, T) {
	for _, v := range c {
		if v.GetSn() == sn {
			return true, v
		}
	}
	var zero T
	return false, zero
}

func (c Collection[T]) GetBySn(sn string) T {
	for _, v := range c {
		if v.GetSn() == sn {
			return v
		}
	}
	var zero T
	return zero
}

// 通用过滤器
func (c Collection[T]) Filter(predicate func(T) bool) []string {
	sns := make([]string, 0)
	for _, v := range c {
		if predicate(v) && v.GetSn() != "" {
			sns = append(sns, v.GetSn())
		}
	}
	return sns
}

func (c Collection[T]) GetByEnterprise(enterprise string) []string {
	return c.Filter(func(v T) bool { return v.GetEnterprise() == enterprise })
}

/*
func (c Collection[T]) GetByEnterprise(enterprise string) []string {
	sns := make([]string, 0)

	for _, v := range c {
		if v.GetEnterprise() == enterprise && v.GetSn() != "" {
			sns = append(sns, v.GetSn())
		}
	}
	return sns
}
*/

type DveProvider interface {
	GetCollection() Collection[DveInfo]
}
