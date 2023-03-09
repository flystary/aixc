package cpe

import (
	"strconv"
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
	var max int
	var maxs string

	for _, c := range n {
		softwareVersion := c.SoftwareVersion
		if softwareVersion != "" {
			versions := strings.Split(softwareVersion, ".")
			one, _   := strconv.Atoi(versions[0])
			two, _   := strconv.Atoi(versions[1])
			three, _ := strconv.Atoi(versions[2])
			num := (one * 1000) + (two * 10) + three

			if num > max {
				max = num
				maxs = c.SoftwareVersion
			}
		}
	}
	return maxs
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