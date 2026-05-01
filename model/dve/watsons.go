package dve

type Watsons struct {
	Total int             `json:"total"`
	Data  Collection[Wde] `json:"data"` // 定义的泛型切片类型
}

func (w Watsons) Collection() Collection[Wde] { return w.Data }
