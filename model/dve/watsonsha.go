package dve

type WatsonsHa []Wde


func (wh WatsonsHa) IsSn(sn string) (bool, Wde) {
	var dve Wde
	for _, d := range wh {
		if sn == d.Sn {
			return true, d
		}
		continue
	}
	return false, dve
}

func (wh WatsonsHa) GetDveStructBySn(sn string) Wde {
	var dve Wde
	for _, d := range wh {
		if sn == d.Sn {
			dve = d
			break
		}
		continue
	}
	return dve
}
