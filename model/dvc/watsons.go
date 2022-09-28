package dvc

type Watsons []Wdvc


func (w Watsons) IsSn(sn string) (bool, Wdvc) {
	var wdvc Wdvc
	for _, d := range w {
		if sn == d.Sn {
			return true, d
		} 
		continue
	}
	return false, wdvc
}

func (w Watsons) GetDvcStructBySn(sn string) Wdvc {
	var dvc Wdvc
	for _, d := range w {
		if sn == d.Sn {
			dvc = d
			break
		}
		continue
	}
	return dvc
}

