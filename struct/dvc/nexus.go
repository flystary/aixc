package dvc
// Nexus 数组nDvc
type Nexus []Ndvc

//IsSn Nexus是否存在Sn
func (n Nexus) IsSn(sn string) (bool, Ndvc) {
	var ndvc Ndvc
	for _, d := range n {
		if sn == d.Sn {
			return true, d
		} 
		continue
	}
	return false, ndvc
}
//GetStructBySn 根据SN获取struct
func (n Nexus) GetStructBySn(sn string) Ndvc {
	var dvc Ndvc
	for _, d := range n {
		if sn == d.Sn {
			dvc = d
			break
		}
		continue
	}
	return dvc
}

