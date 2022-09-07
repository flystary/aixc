package cpe

type Zeratul []Cpe

func (z Zeratul) IsSn(sn string) (bool, Cpe) {
	for i := 0; i < len(z); i++ {
		if sn == z[i].Sn {
			return true, z[i]
		}
		continue
	}
	return false, cpe
}

func (z Zeratul) GetCpeStructBySn(sn string) Cpe {
	for _, c := range z {
		if sn == c.Sn {
			cpe = c
			break
		}
		continue
	}
	return cpe
}
