package cpe

import (
	"strconv"
	"strings"
)

type Watsons struct {
	Total int   `json:"total"`
	Data  []Vox `json:"data"`
}

func (w Watsons) IsSn(sn string) (bool, Vox) {

	for i := 0; i < len(w.Data); i++ {
		if sn == w.Data[i].Sn {
			return true, w.Data[i]
		}
		continue
	}
	return false, vox
}

func (w Watsons) GetCpeStructBySn(sn string) Vox {

	for _, vb := range w.Data {
		if sn == vb.Sn {
			vox = vb
			break
		}
		continue
	}
	return vox
}

func (w Watsons) MaxVersion() string {
	var max int
	var maxs string

	for _, c := range w.Data {
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
