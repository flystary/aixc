package conf

import "fmt"

// Rules 文件
type Rules struct {
	URL	        string		`yaml:"url"`
	Operation 	string		`yaml:"operation"`
	Toekn 		string		`yaml:"token"`
	Valor		Valor		`yaml:"valor"`
	Nexus		Nexus		`yaml:"nexus"`
	Watsons		Watsons		`yaml:"watsons"`
	Tassadar	Tassadar	`yaml:"tassadar"`
	WatsonsHa	WatsonsHa	`yaml:"watsonsha"`
	Modes   	[]string	`yaml:"modes"`
}

// Mode 信息类别
type Mode struct {
	Cpe  string  	`yaml:"cpe"`
	Pop  string  	`yaml:"pop"`
	Dvc  string  	`yaml:"dvc"`
}
// Valor sdwan7
type Valor     Mode
// Nexus sdwan6
type Nexus     Mode
// Watsons 屈臣氏
type Watsons   Mode
// Tassadar sase2
type Tassadar  Mode
// WatsonsHa 屈臣氏仓库
type WatsonsHa Mode

// PopRouteByMode pop路由
func (r *Rules) PopRouteByMode(mode string) string {
	switch mode {
		case "nexus":
			return fmt.Sprintf("%s/nexus/%s?", r.URL, r.Nexus.Pop)
		case "valor":
			return fmt.Sprintf("%s/valor/%s?", r.URL, r.Valor.Pop)
		case "watsons":
			return fmt.Sprintf("%s/watsons/%s?", r.URL, r.Watsons.Pop)
		case "watsonsha":
			return fmt.Sprintf("%s/watsons_ha/%s?", r.URL, r.WatsonsHa.Pop)
		case "tassadar":
			return fmt.Sprintf("%s/tassadar/%s?", r.URL, r.Tassadar.Pop)
		default:
			return ""
	}
}

// CpeRouteByMode cpe路由
func (r *Rules) CpeRouteByMode(mode string) string {
	switch mode {
		case "nexus":
			return fmt.Sprintf("%s/nexus/%s?", r.URL, r.Nexus.Cpe)
		case "valor":
			return fmt.Sprintf("%s/valor/%s?", r.URL, r.Valor.Cpe)
		case "watsons":
			return fmt.Sprintf("%s/watsons/%s?", r.URL, r.Watsons.Cpe)
		case "watsonsha":
			return fmt.Sprintf("%s/watsons_ha/%s?", r.URL, r.WatsonsHa.Cpe)
		case "tassadar":
			return fmt.Sprintf("%s/tassadar/%s?", r.URL, r.Tassadar.Cpe)
		default:
			return ""
	}
}

// DeviceRouteByMode devic路由
func (r *Rules) DeviceRouteByMode(mode string) string {
	switch mode {
		case "nexus":
			return fmt.Sprintf("%s/nexus/%s?", r.URL, r.Nexus.Dvc)
		case "valor":
			return fmt.Sprintf("%s/valor/%s?", r.URL, r.Valor.Dvc)
		case "watsons":
			return fmt.Sprintf("%s/watsons/%s?", r.URL, r.Watsons.Dvc)
		case "watsonsha":
			return fmt.Sprintf("%s/watsons_ha/%s?", r.URL, r.WatsonsHa.Dvc)
		case "tassadar":
			return fmt.Sprintf("%s/tassadar/%s?", r.URL, r.Tassadar.Dvc)
		default:
			return ""
	}
}

// OperationRoute Operation路由
func (r *Rules) OperationRoute() string {
	return fmt.Sprintf("%s/%s?", r.URL, r.Operation)
}

// TokenRoute token路由
func (r *Rules) TokenRoute() string {
	return fmt.Sprintf("%s/%s?", r.URL, r.Toekn)
}