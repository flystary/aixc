package cpe

type Zeratul struct {
	Data Collection[Spe]
}

func (z Zeratul) Collection() Collection[Spe] { return z.Data }

func (z Zeratul) GetCollection() Collection[CpeInfo] {
	res := make(Collection[CpeInfo], len(z.Data))
	for i, p := range z.Data {
		res[i] = p
	}
	return res

}
