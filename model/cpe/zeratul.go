package cpe

import (
	"strconv"
	"strings"
)

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

func (z Zeratul) MaxVersion() string {
	var max  int
	var maxs string
	// var min int
	for _, c := range z {
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