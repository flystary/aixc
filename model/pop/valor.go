package pop


type Valor []Pop

func (v Valor) IsId(id int) (bool, Pop) {

	for _ , p := range v {
		if id == p.ID {
			return true, p
		}
		continue
	}
	return false, pop
}

func (n Valor) GetPopStructById(id int) Pop {

	for _, p := range n {
		if id == p.ID {
			pop = p
		}
	}
	return pop
}

func (v Valor) GetIdByAddr(addr string) (id int) {
	for _, pop := range v {
		if addr == pop.PopIP {
			id = pop.ID
		}
	}
	return id
}