package cpe


type Nexus []Box

func (n Nexus) IsSn(sn string) (bool, Box) {

	for _, b := range n {
		if sn == b.Sn {
			return true, b
		} 
		continue
	}
	return false, box
}

func (n Nexus) GetCpeStructBySn(sn string) Box {

	for _, b := range n {
		if sn == b.Sn {
			box = b
			break
		}
		continue
	}
	return box
}

