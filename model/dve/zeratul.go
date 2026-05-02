package dve

type Zeratul struct {
	Data Collection[Vde]
}

func (z Zeratul) Collection() Collection[Vde] { return z.Data }

func (z Zeratul) GetCollection() Collection[DveInfo] {
	res := make(Collection[DveInfo], len(z.Data))
	for i, p := range z.Data {
		res[i] = p
	}
	return res

}
