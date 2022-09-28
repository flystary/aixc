package cpe

type WatsonsHa []VBox

func (wh WatsonsHa) IsSn(sn string) (bool, VBox) {

	for i := 0; i < len(wh); i++ {
		if sn == wh[i].Sn {
			return true, wh[i]
		} 
		continue
	}
	return false, vox
}

func (wh WatsonsHa) GetCpeStructBySn(sn string) VBox {

	for _, vb := range wh {
		if sn == vb.Sn {
			vox = vb
			break
		}
		continue
	}
	return vox
}
