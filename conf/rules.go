package conf

import "fmt"

type Rules struct {
	URL	        string		`yaml:"url"`
	Operation 	string		`yaml:"operation"`
	Toekn 		string		`yaml:"token"`
	Valor		Valor		`yaml:"valor"`
	Nexus		Nexus		`yaml:"nexus"`
	Watsons		Watsons		`yaml:"watsons"`
	Tassadar	Tassadar	`yaml:"tassadar"`
	WatsonsHa	WatsonsHa	`yaml:"watsons_ha"`
	Modes   	[]string	`yaml:"modes"`
}

type Mode struct {
	Cpe  string  	`yaml:"cpe"`
	Pop  string  	`yaml:"pop"`
	Dvc  string  	`yaml:"dvc"`
}

type Valor     Mode
type Nexus     Mode
type Watsons   Mode
type Tassadar  Mode
type WatsonsHa Mode

func (r *Rules) PopRouteByMode(mode string) string {
	switch mode {
		case "nexus":
			return fmt.Sprintf("%s/nexus/%s?", r.URL, r.Nexus.Pop)
		case "valor":
			return fmt.Sprintf("%s/valor/%s?", r.URL, r.Valor.Pop)
		case "watsons":
			return fmt.Sprintf("%s/watsons/%s?", r.URL, r.Watsons.Pop)
		case "watsons_ha":
			return fmt.Sprintf("%s/watsons_ha/%s?", r.URL, r.WatsonsHa.Pop)
		case "tassadar":
			return fmt.Sprintf("%s/tassadar/%s?", r.URL, r.Tassadar.Pop)
		default:
			return ""
	}
}

func (r *Rules) CpeRouteByMode(mode string) string {
	switch mode {
		case "nexus":
			return fmt.Sprintf("%s/nexus/%s?", r.URL, r.Nexus.Cpe)
		case "valor":
			return fmt.Sprintf("%s/valor/%s?", r.URL, r.Valor.Cpe)
		case "watsons":
			return fmt.Sprintf("%s/watsons/%s?", r.URL, r.Watsons.Cpe)
		case "watsons_ha":
			return fmt.Sprintf("%s/watsons_ha/%s?", r.URL, r.WatsonsHa.Cpe)
		case "tassadar":
			return fmt.Sprintf("%s/tassadar/%s?", r.URL, r.Tassadar.Cpe)
		default:
			return ""
	}
}

func (r *Rules) DeviceRouteByMode(mode string) string {
	switch mode {
		case "nexus":
			return fmt.Sprintf("%s/nexus/%s?", r.URL, r.Nexus.Dvc)
		case "valor":
			return fmt.Sprintf("%s/valor/%s?", r.URL, r.Valor.Dvc)
		case "watsons":
			return fmt.Sprintf("%s/watsons/%s?", r.URL, r.Watsons.Dvc)
		case "watsons_ha":
			return fmt.Sprintf("%s/watsons_ha/%s?", r.URL, r.WatsonsHa.Dvc)
		case "tassadar":
			return fmt.Sprintf("%s/tassadar/%s?", r.URL, r.Tassadar.Dvc)
		default:
			return ""
	}
}

func (r *Rules) OperationRouteByMode() string {
	return fmt.Sprintf("%s/%s?", r.URL, r.Operation)
}

func (r *Rules) TokenRouteByMode() string {
	return fmt.Sprintf("%s/%s?", r.URL, r.Toekn)
}
