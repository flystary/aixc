package pop

type Watsons struct {
	Data Collection[Entry]
}

func (w Watsons) Collection() Collection[Entry] { return w.Data }

func (w Watsons) GetCollection() Collection[PopInfo] {
	res := make(Collection[PopInfo], len(w.Data))
	for i, p := range w.Data {
		res[i] = p
	}
	return res
}
