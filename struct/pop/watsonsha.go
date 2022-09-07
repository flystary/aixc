package pop

type WatsonsHa []Entry

func (wh WatsonsHa) IsId(id int) (bool, Entry) {

	for _ , e := range wh {
		if id == e.ID {
			return true, e
		} 
		continue
	}
	return false, entry
}

func (wh WatsonsHa) GetPopStructBySn(id int) Entry {

	for _, e := range wh {
		if id == e.ID {
			entry = e
		}
	}
	return entry
}
