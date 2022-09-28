package cpe

type Zeratul []SCpe

func (z Zeratul) IsSn(sn string) (bool, SCpe) {
	for i := 0; i < len(z); i++ {
		if sn == z[i].Sn {
			return true, z[i]
		}
		continue
	}
	return false, spe
}

func (z Zeratul) GetCpeStructBySn(sn string) SCpe {
	for _, c := range z {
		if sn == c.Sn {
			spe = c
			break
		}
		continue
	}
	return spe
}
