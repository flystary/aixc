package dve

type WatsonsHa struct {
	Data Collection[Wde]
}

func (wh WatsonsHa) Collection() Collection[Wde] { return wh.Data }

func (wh WatsonsHa) GetCollection() Collection[DveInfo] {
	res := make(Collection[DveInfo], len(wh.Data))
	for i, p := range wh.Data {
		res[i] = p
	}
	return res
}
