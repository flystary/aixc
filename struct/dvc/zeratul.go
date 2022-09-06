package dvc


type Zeratul []Vdvc


func (z Zeratul) IsSn(sn string) (bool, Vdvc) {
	var vdvc Vdvc
	for _, d := range z {
		if sn == d.Sn {
			return true, d
		} 
		continue
	}
	return false, vdvc
}

func (z Zeratul) GetStructBySn(sn string) Vdvc {
	var dvc Vdvc
	for _, d := range z {
		if sn == d.Sn {
			dvc = d
			break
		}
		continue
	}
	return dvc
}

