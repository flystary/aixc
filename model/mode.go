package model

// M 自定义int
type M int

const (
	watsons M = iota + 1
	valor
	yifeng
	watsonsHa
	tassadar
)

var (
	mToStr = map[M]string{
		valor:     "valor",
		yifeng:    "yifeng",
		watsons:   "watsons",
		watsonsHa: "watsonsha",
		tassadar:  "tassadar",
	}

	strToM = map[string]M{}
)

func (m M) String() string {
	if v, ok := mToStr[m]; ok {
		return v
	}
	return "unknown"
}

func (m M) Valid() bool {
	switch m {
	case watsons, valor, yifeng, watsonsHa, tassadar:
		return true
	default:
		return false
	}
}

func ParseM(s string) (M, bool) {
	m, ok := strToM[s]
	return m, ok
}
