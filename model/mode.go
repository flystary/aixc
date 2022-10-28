package model

//M 自定义int
type M int

const (
	valor M = iota + 1
	nexus
	watsons
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
	case nexus:
		{
			mode = "nexus"
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
	default:
		{
			mode = "unknown"
		}
	}
	return mode
}
