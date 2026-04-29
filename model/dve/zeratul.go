package dve

type Zeratul []Vde

func (z Zeratul) IsSn(sn string) (bool, Vde) {
	return IsSn(z, sn)
}

func (z Zeratul) GetDveStructBySn(sn string) Vde {
	return GetBySn(z, sn)
}

func (z Zeratul) GetCpesByEnterprise(enterprise string) []string {
	return GetByEnterprise(z, enterprise)
}
