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

func (wh WatsonsHa) GetPopStructById(id int) Entry {

	for _, e := range wh {
		if id == e.ID {
			entry = e
		}
	}
	return entry
}

func (wh WatsonsHa) GetIdByAddr(addr string) (id int) {
	for _, pop := range wh {
		if addr == pop.EntryIP {
			id = pop.ID
		}
	}
	return id
}