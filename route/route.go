package route

import (
	"fmt"
	"os"
	"sync"
	"gopkg.in/yaml.v3"
)



var (
	// 文件单例
	once = &sync.Once{}
)

// LoadRoute 加载Route
func LoadRoute(path string) Route {
	var route Route
	once.Do(func() {
		io, err := os.ReadFile(path)
		if err != nil {
			fmt.Printf("Open File Error: %v", err)
		}
		err = yaml.Unmarshal(io, &route)
		if err != nil {
			fmt.Printf("Unmarshal File Error: %v", err)
		}
	})
	return route
}

func (r Route) GetCpeFromRoute(mode string) string {
	switch mode {
	case "nexus":
		return fmt.Sprintf("%s/nexus/%s?", r.InitURL, r.Nexus.Cpe)
	case "valor":
		return fmt.Sprintf("%s/valor/%s?&pageSize=%v&", r.InitURL, r.Valor.Cpe, r.Valor.Pse)
	case "watsons":
		return fmt.Sprintf("%s/watsons/%s?", r.InitURL, r.Watsons.Cpe)
	case "watsonsha":
		return fmt.Sprintf("%s/watsons_ha/%s?", r.InitURL, r.WatsonsHa.Cpe)
	case "tassadar":
		return fmt.Sprintf("%s/tassadar/%s?", r.InitURL, r.Tassadar.Cpe)
	default:
		return ""
	}
}

func (r Route) GetPopFromRoute(mode string) string {
	switch mode {
	case "nexus":
		return fmt.Sprintf("%s/nexus/%s?", r.InitURL, r.Nexus.Pop)
	case "valor":
		return fmt.Sprintf("%s/valor/%s?", r.InitURL, r.Valor.Pop)
	case "watsons":
		return fmt.Sprintf("%s/watsons/%s?", r.InitURL, r.Watsons.Pop)
	case "watsonsha":
		return fmt.Sprintf("%s/watsons_ha/%s?", r.InitURL, r.WatsonsHa.Pop)
	case "tassadar":
		return fmt.Sprintf("%s/tassadar/%s?", r.InitURL, r.Tassadar.Pop)
	default:
		return ""
	}
}

func (r Route) GetDveFromRoute(mode string) string {
	switch mode {
	case "nexus":
		return fmt.Sprintf("%s/nexus/%s", r.InitURL, r.Nexus.Dve)
	case "valor":
		return fmt.Sprintf("%s/valor/%s", r.InitURL, r.Valor.Dve)
	case "watsons":
		return fmt.Sprintf("%s/watsons/%s", r.InitURL, r.Watsons.Dve)
	case "watsonsha":
		return fmt.Sprintf("%s/watsons_ha/%s", r.InitURL, r.WatsonsHa.Dve)
	case "tassadar":
		return fmt.Sprintf("%s/tassadar/%s", r.InitURL, r.Tassadar.Dve)
	default:
		return ""
	}
}

func (r Route) GetTokenFromRoute() string {
	return fmt.Sprintf("%s/%s?", r.InitURL, r.Tokenurl)
}

func (r Route) GetOperationFromRoute() string {
	return fmt.Sprintf("%s/%s?", r.InitURL, r.Operation)
}