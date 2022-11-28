package dve

type Valor []Vde

func (v Valor) IsSn(sn string) (bool, Vde) {
	var vdve Vde
	for _, d := range v {
		if sn == d.Sn {
			return true, d
		}
		continue
	}
	return false, vdve
}

func (v Valor) GetDveStructBySn(sn string) Vde {
	var dve Vde
	for _, d := range v {
		if sn == d.Sn {
			dve = d
			break
		}
		continue
	}
	return dve
}

func (v Valor) GetCpesByEnterprise(enterprise string) []string {
	for _, data := range v {
		if enterprise == data.CustomerName {
			sns = append(sns, data.Sn)
		}
	}
	return sns
}
