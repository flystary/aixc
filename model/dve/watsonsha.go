package dve

type WatsonsHa []Wde

func (wh WatsonsHa) IsSn(sn string) (bool, Wde) {
	return IsSn(wh, sn)
}

func (wh WatsonsHa) GetDveStructBySn(sn string) Wde {
	return GetBySn(wh, sn)
}

func (wh WatsonsHa) GetCpesByEnterprise(enterprise string) []string {
	return GetByEnterprise(wh, enterprise)
}
