package cpe

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
