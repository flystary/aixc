package dve

type Watsons struct {
	Total int   `json:"total"`
	Data  []Wde `json:"data"`
}

func (w Watsons) IsSn(sn string) (bool, Wde) {
	return IsSn(w.Data, sn)
}

func (w Watsons) GetDveStructBySn(sn string) Wde {
	return GetBySn(w.Data, sn)
}

func (w Watsons) GetCpesByEnterprise(enterprise string) []string {
	return GetByEnterprise(w.Data, enterprise)
}
