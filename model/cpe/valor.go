package cpe

type Valor struct {
	Total int             `json:"total"`
	Data  Collection[Cpe] `json:"data"` // 定义的泛型切片类型
}

func (v Valor) Collection() Collection[Cpe] { return v.Data }

func (v Valor) GetCollection() Collection[CpeInfo] {
	res := make(Collection[CpeInfo], v.Data.Len())
	for i, p := range v.Data {
		res[i] = p
	}
	return res

}
