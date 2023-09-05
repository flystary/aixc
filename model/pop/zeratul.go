package pop

type Zeratul []SPop
var spop  SPop
func (z Zeratul) IsId(id int) (bool, SPop) {
	for _ , sp := range z {
		if id == sp.ID {
			return true, sp
		}
		continue
	}
	return false, spop
}

func (z Zeratul) GetPopStructById(id int) SPop {
	var pop  SPop
	for _, p := range z {
		if id == p.ID {
			pop = p
		}
	}
	return pop
}

func (z Zeratul) GetIdByAddr(addr string) (id int) {
	for _, pop := range z {
		if addr == pop.PopIP {
			id = pop.ID
		}
	}
	return id
}
