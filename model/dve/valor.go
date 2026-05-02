package dve

type Valor struct {
	Data Collection[Vde]
}

func (v Valor) Collection() Collection[Vde] { return v.Data }

func (v Valor) GetCollection() Collection[DveInfo] {
	res := make(Collection[DveInfo], len(v.Data))
	for i, p := range v.Data {
		res[i] = p
	}
	return res

}
