package cpe

type Valor struct {
	Total int   `json:"total"`
	Data  []Cpe `json:"data"`
}

func (v Valor) IsSn(sn string) (bool, Cpe) {
	return IsSn(v.Data, sn)
}

func (v Valor) SNs() []string {
	return SNs(v.Data)
}

func (v Valor) IsExist(sn string) bool {
	return IsExist(v.Data, sn)
}

func (v Valor) GetCpeStructBySn(sn string) (Cpe, bool) {
	return GetBySn(v.Data, sn)
}

func (v Valor) MaxVersion() string {
	return MaxVersion(v.Data)
}

func (v Valor) GetCpesByModel(model string) []string {
	return GetByModel(v.Data, model)
}

func (v Valor) GetCpesByVersion(version string) []string {
	return GetByVersion(v.Data, version)
}

func (v Valor) GetCpesByPopId(id int) []string {
	return GetByPopId(v.Data, id)
}
