package dve

func IsSn[T DveGet](data []T, sn string) (bool, T) {
	for _, v := range data {
		if v.GetSn() == sn {
			return true, v
		}
	}
	var zero T
	return false, zero
}

func GetBySn[T DveGet](data []T, sn string) T {
	for _, v := range data {
		if v.GetSn() == sn {
			return v
		}
	}
	var zero T
	return zero
}

func GetByEnterprise[T DveGet](data []T, enterprise string) []string {
	sns := make([]string, 0)

	for _, v := range data {
		if v.GetEnterprise() == enterprise && v.GetSn() != "" {
			sns = append(sns, v.GetSn())
		}
	}
	return sns
}
