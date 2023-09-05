package dve


type Zeratul []Vde


func (z Zeratul) IsSn(sn string) (bool, Vde) {
	var vdve Vde
	for _, d := range z {
		if sn == d.Sn {
			return true, d
		}
		continue
	}
	return false, vdve
}

func (z Zeratul) GetDveStructBySn(sn string) Vde {
	var dve Vde
	for _, d := range z {
		if sn == d.Sn {
			dve = d
			break
		}
		continue
	}
	return dve
}

func (z Zeratul) GetCpesByEnterprise(enterprise string) []string {
	for _, data := range z {
		if enterprise == data.CustomerName {
			sns = append(sns, data.Sn)
		}
	}
	return sns
}