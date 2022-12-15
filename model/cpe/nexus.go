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
	var max  int
	var maxs string
	// var min int
	for _, c := range n {
		versions := strings.Split(c.SoftwareVersion, ".")
		hundred, _:= strconv.Atoi(versions[0])
		unit, _:= strconv.Atoi(versions[1])
		num := ( hundred * 100 ) + unit

		if max <= num {
			max  = num
			maxs = c.SoftwareVersion
		}
	}
	return maxs
}