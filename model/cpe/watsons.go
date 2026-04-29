package cpe

type Watsons struct {
	Total int   `json:"total"`
	Data  []Vox `json:"data"`
}

func (w Watsons) IsSn(sn string) (bool, Vox) {
	return IsSn(w.Data, sn)
}

func (w Watsons) SNs() []string {
	return SNs(w.Data)
}

func (w Watsons) IsExist(sn string) bool {
	return IsExist(w.Data, sn)
}

func (w Watsons) GetCpeStructBySn(sn string) (Vox, bool) {
	return GetBySn(w.Data, sn)
}

func (w Watsons) MaxVersion() string {
	return MaxVersion(w.Data)
}

func (w Watsons) GetCpesByModel(model string) []string {
	return GetByModel(w.Data, model)
}

func (w Watsons) GetCpesByVersion(version string) []string {
	return GetByVersion(w.Data, version)
}

func (w Watsons) GetCpesByPopId(id int) []string {
	return GetByPopId(w.Data, id)
}
