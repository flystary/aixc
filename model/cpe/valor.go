package cpe

import "strings"

type Valor struct {
	Total int   `json:"total"`
	Data  []Cpe `json:"data"`
}

func (v Valor) IsSn(sn string) (bool, Cpe) {

	for _, c := range v.Data {
		if sn == c.Sn {
			return true, c
		}
		continue
	}
	return false, cpe
}

func (v Valor) IsExist(sn string) bool {

	for _, c := range v.Data {
		if sn == c.Sn {
			return true
		}
		continue
	}
	return false
}

func (v Valor) SNs() []string {
	var sns = make([]string, 0)
	for _, c := range v.Data {
		if c.Sn != "" {
			sns = append(sns, c.Sn)
		}
		continue
	}
	return sns
}

func (v Valor) GetCpeStructBySn(sn string) Cpe {

	for _, c := range v.Data {
		if sn == c.Sn {
			cpe = c
			break
		}
		continue
	}
	return cpe
}

func (v Valor) MaxVersion() string {
	var max string

	for _, c := range v.Data {
		if c.SoftwareVersion == "" {
			continue
		}
		vers := strings.Split(c.SoftwareVersion, ".")
		maxs := strings.Split(max, ".")
		lens := len(vers)

		if lens < 2 {
			continue
		} else if lens == 2 {
			vers = append(vers, "0")
		}

		for i := 0; i < lens; i++ {
			if vers[i] > maxs[i] {
				max = c.SoftwareVersion
				break
			}
		}
	}
	return max
}

func (v Valor) GetCpesByModel(model string) []string {
	for _, data := range v.Data {
		if model == data.Model {
			sns = append(sns, data.Sn)
		}
	}
	return sns
}

func (v Valor) GetCpesByVersion(version string) []string {
	for _, data := range v.Data {
		if version == data.SoftwareVersion {
			sns = append(sns, data.Sn)
		}
	}
	return sns
}

func (v Valor) GetCpesByPopId(id int) []string {
	for _, data := range v.Data {
		if id == data.MasterPopID || id == data.BackupPopID {
			sns = append(sns, data.Sn)
		}
	}
	return sns
}
