package cpe

import "strings"

type Zeratul []Spe

func (z Zeratul) IsSn(sn string) (bool, Spe) {
	for i := 0; i < len(z); i++ {
		if sn == z[i].Sn {
			return true, z[i]
		}
		continue
	}
	return false, spe
}

func (z Zeratul) IsExist(sn string) bool {
	for i := 0; i < len(z); i++ {
		if sn == z[i].Sn {
			return true
		}
		continue
	}
	return false
}

func (z Zeratul) GetCpeStructBySn(sn string) Spe {
	for _, c := range z {
		if sn == c.Sn {
			spe = c
			break
		}
		continue
	}
	return spe
}

func (z Zeratul) SNs() []string {
	var sns = make([]string, 0)
	for _, c := range z {
		if c.Sn != "" {
			sns = append(sns, c.Sn)
		}
		continue
	}
	return sns
}

func (z Zeratul) MaxVersion() string {
	var max string
	for _, c := range z {
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

func (z Zeratul) GetCpesByModel(model string) []string {
	for _, data := range z {
		if model == data.Model {
			sns = append(sns, data.Sn)
		}
	}
	return sns
}

func (z Zeratul) GetCpesByVersion(version string) []string {
	for _, data := range z {
		if version == data.SoftwareVersion {
			sns = append(sns, data.Sn)
		}
	}
	return sns
}

func (z Zeratul) GetCpesByPopId(id int) []string {
	for _, data := range z {
		if id == data.MasterPopID || id == data.BackupPopID {
			sns = append(sns, data.Sn)
		}
	}
	return sns
}
