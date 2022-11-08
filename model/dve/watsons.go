package dve

type Watsons []Wde


func (w Watsons) IsSn(sn string) (bool, Wde) {
	var wdve Wde
	for _, d := range w {
		if sn == d.Sn {
			return true, d
		}
		continue
	}
	return false, wdve
}

func (w Watsons) GetDveStructBySn(sn string) Wde {
	var dve Wde
	for _, d := range w {
		if sn == d.Sn {
			dve = d
			break
		}
		continue
	}
	return dve
}

