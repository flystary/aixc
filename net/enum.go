package net

//Mtype 自定义int
type Mtype int

const (
	valor Mtype = iota + 1
	nexus
	watsons
	watsonsHa
	tassadar
)

func (m Mtype) enum() string {
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
			mode = "watsons_ha"
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