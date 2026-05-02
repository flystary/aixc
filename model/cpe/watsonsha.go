package cpe

type WatsonsHa struct {
	Total int             `json:"total"`
	Data  Collection[Vox] `json:"data"`
}

func (wh WatsonsHa) Collection() Collection[Vox] { return wh.Data }

func (wh WatsonsHa) GetCollection() Collection[CpeInfo] {
	res := make(Collection[CpeInfo], wh.Data.Len())
	for i, p := range wh.Data {
		res[i] = p
	}
	return res

}
