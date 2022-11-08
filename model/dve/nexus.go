package dve
// Nexus 数组nDve
type Nexus []Nde

//IsSn Nexus是否存在Sn
func (n Nexus) IsSn(sn string) (bool, Nde) {
	var dve Nde
	for _, d := range n {
		if sn == d.Sn {
			return true, d
		}
		continue
	}
	return false, dve
}
//GetStructBySn 根据SN获取struct
func (n Nexus) GetDveStructBySn(sn string) Nde {
	var dve Nde
	for _, d := range n {
		if sn == d.Sn {
			dve = d
			break
		}
		continue
	}
	return dve
}

func (n Nexus) GetCpesByEnterprise(enterprise string) []string {
	for _, data := range n {
		if enterprise == data.CustomerName {
			sns = append(sns, data.Sn)
		}
	}
	return sns
}