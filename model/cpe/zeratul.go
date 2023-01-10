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

func (z Zeratul) SNs() []string {
	var sns  = make([]string, 0)
	for _, c := range z {
		if c.Sn != "" {
			sns = append(sns, c.Sn)
		}
		continue
	}
	return sns
}


func (z Zeratul) MaxVersion() string {
	var max int
	var maxs string

	for _, c := range z {
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
