package pop


type Nexus []Entry


func (n Nexus) IsId(id int) (bool, Entry) {

	for _ , e := range n {
		if id == e.ID {
			return true, e
		} 
		continue
	}
	return false, entry
}

func (n Nexus) GetPopStructById(id int) Entry {

	for _, e := range n {
		if id == e.ID {
			entry = e
		}
	}
	return entry
}

