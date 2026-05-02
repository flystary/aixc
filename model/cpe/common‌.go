package cpe

type Collection[T CpeInfo] []T

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

func (c Collection[T]) SNs() []string {
	sns := make([]string, 0, c.Len())

	for _, v := range c {
		if sn := v.GetSn(); sn != "" {
			sns = append(sns, sn)
		}
	}
	return sns
}

func (c Collection[T]) IsExist(sn string) bool {
	for _, v := range c {
		if v.GetSn() == sn {
			return true
		}
	}
	return false
}

func (c Collection[T]) MaxVersion() string {
	var max string

	for _, v := range c {
		ver := v.GetVersion()
		if ver == "" {
			continue
		}

		if max == "" || compareVersion(ver, max) > 0 {
			max = ver
		}
	}
	return max
}

func (c Collection[T]) GetBySn(sn string) (T, bool) {
	for _, v := range c {
		if v.GetSn() == sn {
			return v, true
		}
	}
	var zero T
	return zero, false
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

func (c Collection[T]) GetByModel(model string) []string {
	return c.Filter(func(v T) bool { return v.GetModel() == model })
}

func (c Collection[T]) GetByVersion(version string) []string {
	return c.Filter(func(v T) bool { return v.GetVersion() == version })
}

func (c Collection[T]) GetByPopId(id int) []string {
	return c.Filter(func(v T) bool {
		return v.GetMasterPopID() == id || v.GetBackupPopID() == id
	})
}

/*
func (c Collection[T]) GetByModel(model string) []string {
	sns := make([]string, 0)

	for _, v := range c {
		if v.GetModel() == model && v.GetSn() != "" {
			sns = append(sns, v.GetSn())
		}
	}
	return sns
}

func (c Collection[T]) GetByVersion(version string) []string {
	sns := make([]string, 0)

	for _, v := range c {
		if v.GetVersion() == version && v.GetSn() != "" {
			sns = append(sns, v.GetSn())
		}
	}
	return sns
}

func (c Collection[T]) GetByPopId(id int) []string {
	sns := make([]string, 0)

	for _, v := range c {
		if (v.GetMasterPopID() == id || v.GetBackupPopID() == id) && v.GetSn() != "" {
			sns = append(sns, v.GetSn())
		}
	}
	return sns
}
*/

type CpeProvider interface {
	GetCollection() Collection[CpeInfo]
}
