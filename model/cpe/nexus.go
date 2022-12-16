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
			one, _ := strconv.Atoi(versions[0])
			two, _ := strconv.Atoi(versions[1])
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
