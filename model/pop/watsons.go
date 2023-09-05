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

func (w Watsons) GetPopStructById(id int) Entry {

	for _, e := range w {
		if id == e.ID {
			entry = e
		}
	}
	return entry
}

func (w Watsons) GetIdByAddr(addr string) (id int) {
	for _, pop := range w {
		if addr == pop.EntryIP {
			id = pop.ID
		}
	}
	return id
}