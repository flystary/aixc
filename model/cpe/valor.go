package cpe

import (
	"strconv"
	"strings"
)

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

func (v Valor) SNs() []string {
	var sns  = make([]string, 0)
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
	var max = 0
	var maxs string

	for _, c := range v.Data {
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
