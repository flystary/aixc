package cpe

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
