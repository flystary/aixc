package cpe

type WatsonsHa []Vox

func (wh WatsonsHa) IsSn(sn string) (bool, Vox) {

	for i := 0; i < len(wh); i++ {
		if sn == wh[i].Sn {
			return true, wh[i]
		}
		continue
	}
	return false, vox
}

func (wh WatsonsHa) GetCpeStructBySn(sn string) Vox {

	for _, vb := range wh {
		if sn == vb.Sn {
			vox = vb
			break
		}
		continue
	}
	return vox
}
