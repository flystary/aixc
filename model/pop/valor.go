package pop

type Valor struct {
	Data Collection[Pop]
}

func (v Valor) Collection() Collection[Pop] { return v.Data }

func (v Valor) GetCollection() Collection[PopInfo] {
	res := make(Collection[PopInfo], len(v.Data))
	for i, p := range v.Data {
		res[i] = p
	}
	return res
}
