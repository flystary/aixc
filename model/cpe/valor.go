package cpe

import (
	"strconv"
	"strings"
)

type Valor struct {
	Total 			int   `json:"total"`
	Data 			[]Cpe `json:"data"`
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
	var max  = 0
	var maxs string
	// var min int
	for _, c := range v.Data {
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
