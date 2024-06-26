package model

//M 自定义int
type M int

const (
	watsons M = iota + 1
	valor
    yifeng
	watsonsHa
	tassadar
)

// Enum 枚举
func (m M) Enum() string {
	var mode string
	switch m {
	case valor:
		{
			mode = "valor"
		}
	case yifeng:
		{
			mode = "yifeng"
		}
	case watsons:
		{
			mode = "watsons"
		}
	case watsonsHa:
		{
			mode = "watsonsha"
		}
	case tassadar:
		{
			mode = "tassadar"
		}
	}
	return mode
}
