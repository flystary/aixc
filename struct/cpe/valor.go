package cpe

type Valor []Cpe

func (v Valor) IsSn(sn string) (bool, Cpe) {

	for _, c := range v {
		if sn == c.Sn {
			return true, c
		} 
		continue
	}
	return false, cpe
}

func (v Valor) GetCpeStruct(sn string) Cpe {

	for _, c := range v {
		if sn == c.Sn {
			cpe = c
			break
		}
		continue
	}
	return cpe
}
