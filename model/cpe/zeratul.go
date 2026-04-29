package cpe

type Zeratul []Spe

func (z Zeratul) IsSn(sn string) (bool, Spe) {
	return IsSn(z, sn)
}

func (z Zeratul) SNs() []string {
	return SNs(z)
}

func (z Zeratul) IsExist(sn string) bool {
	return IsExist(z, sn)
}

func (z Zeratul) GetCpeStructBySn(sn string) (Spe, bool) {
	return GetBySn(z, sn)
}

func (z Zeratul) MaxVersion() string {
	return MaxVersion(z)
}

func (z Zeratul) GetCpesByModel(model string) []string {
	return GetByModel(z, model)
}

func (z Zeratul) GetCpesByVersion(version string) []string {
	return GetByVersion(z, version)
}

func (z Zeratul) GetCpesByPopId(id int) []string {
	return GetByPopId(z, id)
}
