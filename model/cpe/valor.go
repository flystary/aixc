package cpe

type Valor struct {
	Total int             `json:"total"`
	Data  Collection[Cpe] `json:"data"` // 定义的泛型切片类型
}

func (v Valor) Collection() Collection[Cpe] { return v.Data }
