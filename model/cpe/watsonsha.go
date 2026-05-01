package cpe

type WatsonsHa struct {
	Total int             `json:"total"`
	Data  Collection[Vox] `json:"data"`
}

func (wh WatsonsHa) Collection() Collection[Vox] { return wh.Data }
