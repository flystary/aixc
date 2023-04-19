package cpe

import (
	"strings"
)

type Nexus []Box

func (n Nexus) IsSn(sn string) (bool, Box) {

	for _, b := range n {
		if sn == b.Sn {
			return true, b
		}
		continue
	}
	return false, box
}

func (n Nexus) GetCpeStructBySn(sn string) Box {

	for _, b := range n {
		if sn == b.Sn {
			box = b
			break
		}
		continue
	}
	return box
}

func (n Nexus) MaxVersion() string {
	var max string

	for _, c := range n {
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

func (n Nexus) GetCpesByModel(model string) []string {
	for _, data := range n {
		if model == data.Model {
			sns = append(sns, data.Sn)
		}
	}
	return sns
}
func (n Nexus) GetCpesByVersion(version string) []string {
	for _, data := range n {
		if version == data.SoftwareVersion {
			sns = append(sns, data.Sn)
		}
	}
	return sns
}

func (n Nexus) GetCpesByPopId(id int) []string {
	for _, data := range n {
		if id == data.MasterEntryID || id == data.BackupEntryID {
			sns = append(sns, data.Sn)
		}
	}
	return sns
}
