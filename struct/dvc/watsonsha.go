package dvc

type WatsonsHa []Wdvc


func (wh WatsonsHa) IsSn(sn string) (bool, Wdvc) {
	var wdvc Wdvc
	for _, d := range wh {
		if sn == d.Sn {
			return true, d
		} 
		continue
	}
	return false, wdvc
}

func (wh WatsonsHa) GetStructBySn(sn string) Wdvc {
	var dvc Wdvc
	for _, d := range wh {
		if sn == d.Sn {
			dvc = d
			break
		}
		continue
	}
	return dvc
}

