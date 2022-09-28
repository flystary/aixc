package cpe

type Watsons []VBox

func (w Watsons) IsSn(sn string) (bool, VBox){

	for i := 0; i < len(w); i++ {
		if sn == w[i].Sn {
			return true, w[i]
		} 
		continue
	}
	return false, vox
}

func (w Watsons) GetCpeStructBySn(sn string) VBox {

	for _, vb := range w {
		if sn == vb.Sn {
			vox = vb
			break
		}
		continue
	}
	return vox
}
