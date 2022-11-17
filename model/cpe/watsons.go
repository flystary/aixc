package cpe

type Watsons struct {
	Total int   `json:"total"`
	Data  []Vox `json:"data"`
}

func (w Watsons) IsSn(sn string) (bool, Vox){

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
