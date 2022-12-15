package cpe

import (
	"strconv"
	"strings"
)

type WatsonsHa struct {
	Total 			int   `json:"total"`
	Data 			[]Vox `json:"data"`
}
func (wh WatsonsHa) IsSn(sn string) (bool, Vox) {

	for i := 0; i < len(wh.Data); i++ {
		if sn == wh.Data[i].Sn {
			return true, wh.Data[i]
		}
		continue
	}
	return false, vox
}

func (wh WatsonsHa) GetCpeStructBySn(sn string) Vox {

	for _, vb := range wh.Data {
		if sn == vb.Sn {
			vox = vb
			break
		}
		continue
	}
	return vox
}

func (wh WatsonsHa) MaxVersion() string {
	var max  int
	var maxs string
	// var min int
	for _, c := range wh.Data {
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