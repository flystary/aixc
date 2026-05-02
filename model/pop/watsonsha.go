package pop

type WatsonsHa struct {
	Data Collection[Entry]
}

func (wh WatsonsHa) Collection() Collection[Entry] { return wh.Data }

func (wh WatsonsHa) GetCollection() Collection[PopInfo] {
	res := make(Collection[PopInfo], len(wh.Data))
	for i, p := range wh.Data {
		res[i] = p
	}
	return res
}
