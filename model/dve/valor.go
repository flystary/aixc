package dve

type Valor []Vde

func (v Valor) IsSn(sn string) (bool, Vde) {
	return IsSn(v, sn)
}

func (v Valor) GetDveStructBySn(sn string) Vde {
	return GetBySn(v, sn)
}

func (v Valor) GetCpesByEnterprise(enterprise string) []string {
	return GetByEnterprise(v, enterprise)
}
