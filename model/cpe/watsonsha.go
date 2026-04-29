package cpe

type WatsonsHa struct {
	Total int   `json:"total"`
	Data  []Vox `json:"data"`
}

func (wh WatsonsHa) IsSn(sn string) (bool, Vox) {
	return IsSn(wh.Data, sn)
}

func (wh WatsonsHa) SNs() []string {
	return SNs(wh.Data)
}

func (wh WatsonsHa) IsExist(sn string) bool {
	return IsExist(wh.Data, sn)
}

func (wh WatsonsHa) GetCpeStructBySn(sn string) (Vox, bool) {
	return GetBySn(wh.Data, sn)
}

func (wh WatsonsHa) MaxVersion() string {
	return MaxVersion(wh.Data)
}

func (wh WatsonsHa) GetCpesByModel(model string) []string {
	return GetByModel(wh.Data, model)
}

func (wh WatsonsHa) GetCpesByVersion(version string) []string {
	return GetByVersion(wh.Data, version)
}

func (wh WatsonsHa) GetCpesByPopId(id int) []string {
	return GetByPopId(wh.Data, id)
}
