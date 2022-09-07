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

func (n Valor) GetPopStructBySn(id int) Pop {

	for _, p := range n {
		if id == p.ID {
			pop = p
		}
	}
	return pop
}
