package cpe

type Watsons struct {
	Total int             `json:"total"`
	Data  Collection[Vox] `json:"data"`
}

func (w Watsons) Collection() Collection[Vox] { return w.Data }

func (w Watsons) GetCollection() Collection[CpeInfo] {
	res := make(Collection[CpeInfo], w.Data.Len())
	for i, p := range w.Data {
		res[i] = p
	}
	return res

}
