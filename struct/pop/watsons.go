package pop

type Watsons []Entry

func (w Watsons) IsId(id int) (bool, Entry) {

	for _ , e := range w {
		if id == e.ID {
			return true, e
		} 
		continue
	}
	return false, entry
}

func (w Watsons) GetPopStructBySn(id int) Entry {

	for _, e := range w {
		if id == e.ID {
			entry = e
		}
	}
	return entry
}
