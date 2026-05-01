package cpe

type Watsons struct {
	Total int             `json:"total"`
	Data  Collection[Vox] `json:"data"`
}

func (w Watsons) Collection() Collection[Vox] { return w.Data }
