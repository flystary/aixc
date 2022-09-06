package dvc

type Valor []Vdvc


func (v Valor) IsSn(sn string) (bool, Vdvc) {
	var vdvc Vdvc
	for _, d := range v {
		if sn == d.Sn {
			return true, d
		} 
		continue
	}
	return false, vdvc
}

func (v Valor) GetStructBySn(sn string) Vdvc {
	var dvc Vdvc
	for _, d := range v {
		if sn == d.Sn {
			dvc = d
			break
		}
		continue
	}
	return dvc
}

