package pop

type Zeratul struct {
	Data Collection[SPop]
}

func (z Zeratul) Collection() Collection[SPop] { return z.Data }

func (z Zeratul) GetCollection() Collection[PopInfo] {
	res := make(Collection[PopInfo], len(z.Data))
	for i, p := range z.Data {
		res[i] = p
	}
	return res
}
