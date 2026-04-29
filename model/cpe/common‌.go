package cpe

func IsSn[T CpeGet](data []T, sn string) (bool, T) {
	for _, v := range data {
		if v.GetSn() == sn {
			return true, v
		}
	}
	var zero T
	return false, zero
}

func SNs[T CpeGet](data []T) []string {
	sns := make([]string, 0, len(data))

	for _, v := range data {
		if sn := v.GetSn(); sn != "" {
			sns = append(sns, sn)
		}
	}
	return sns
}

func IsExist[T CpeGet](data []T, sn string) bool {
	for _, v := range data {
		if v.GetSn() == sn {
			return true
		}
	}
	return false
}

func MaxVersion[T CpeGet](data []T) string {
	var max string

	for _, v := range data {
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

func GetBySn[T CpeGet](data []T, sn string) (T, bool) {
	for _, v := range data {
		if v.GetSn() == sn {
			return v, true
		}
	}
	var zero T
	return zero, false
}

func GetByModel[T CpeGet](data []T, model string) []string {
	sns := make([]string, 0)

	for _, v := range data {
		if v.GetModel() == model && v.GetSn() != "" {
			sns = append(sns, v.GetSn())
		}
	}
	return sns
}

func GetByVersion[T CpeGet](data []T, version string) []string {
	sns := make([]string, 0)

	for _, v := range data {
		if v.GetVersion() == version && v.GetSn() != "" {
			sns = append(sns, v.GetSn())
		}
	}
	return sns
}

func GetByPopId[T CpeGet](data []T, id int) []string {
	sns := make([]string, 0)

	for _, v := range data {
		if v.GetMasterPopID() == id || v.GetBackupPopID() == id && v.GetSn() != "" {
			sns = append(sns, v.GetSn())
		}
	}
	return sns
}
