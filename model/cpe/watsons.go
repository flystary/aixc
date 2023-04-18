package cpe

import (
	"strings"
)

type Watsons struct {
	Total int   `json:"total"`
	Data  []Vox `json:"data"`
}

func (w Watsons) IsSn(sn string) (bool, Vox) {

	for i := 0; i < len(w.Data); i++ {
		if sn == w.Data[i].Sn {
			return true, w.Data[i]
		}
		continue
	}
	return false, vox
}

func (w Watsons) GetCpeStructBySn(sn string) Vox {

	for _, vb := range w.Data {
		if sn == vb.Sn {
			vox = vb
			break
		}
		continue
	}
	return vox
}

func (w Watsons) MaxVersion() string {
	var max string
	for _, c := range w.Data {
		if c.SoftwareVersion == "" {
			continue
		}
		ves := strings.Split(c.SoftwareVersion, ".")
		mves := strings.Split(max, ".")
		lens := len(ves)

		if lens < 2 {
			continue
		} else if lens == 2 {
			ves = append(ves, "0")
		}

		for i := 0; i < lens; i++ {
			if ves[i] > mves[i] {
				max = c.SoftwareVersion
				break
			}
		}
	}
	return max
}

func (w Watsons) GetCpesByModel(model string) []string {
	for _, data := range w.Data {
		if model == data.Model {
			sns = append(sns, data.Sn)
		}
	}
	return sns
}

func (w Watsons) GetCpesByVersion(version string) []string {
	for _, data := range w.Data {
		if version == data.SoftwareVersion {
			sns = append(sns, data.Sn)
		}
	}
	return sns
}

func (w Watsons) GetCpesByPopId(id int) []string {
	for _, data := range w.Data {
		if id == data.MasterEntryID || id == data.BackupEntryID {
			sns = append(sns, data.Sn)
		}
	}
	return sns
}
