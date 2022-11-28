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
		return fmt.Sprintf("%s/valor/%s?page=1&pageSize=%v&", r.InitURL, r.Valor.Cpe, r.Valor.Pse)
	case "watsons":
		return fmt.Sprintf("%s/watsons/%s?page=1&pageSize=%v&", r.InitURL, r.Watsons.Cpe, r.Watsons.Pse)
	case "watsonsha":
		return fmt.Sprintf("%s/watsons_ha/%s?page=1&pageSize=%v&", r.InitURL, r.WatsonsHa.Cpe, r.WatsonsHa.Pse)
	case "tassadar":
		return fmt.Sprintf("%s/tassadar/%s?", r.InitURL, r.Tassadar.Cpe)
	default:
		return ""
	}
}

func (r Route) GetPopFromRoute(mode string) string {
	switch mode {
	case "nexus":
		return fmt.Sprintf("%s/nexus/%s", r.InitURL, r.Nexus.Pop)
	case "valor":
		return fmt.Sprintf("%s/valor/%s", r.InitURL, r.Valor.Pop)
	case "watsons":
		return fmt.Sprintf("%s/watsons/%s", r.InitURL, r.Watsons.Pop)
	case "watsonsha":
		return fmt.Sprintf("%s/watsons_ha/%s", r.InitURL, r.WatsonsHa.Pop)
	case "tassadar":
		return fmt.Sprintf("%s/tassadar/%s", r.InitURL, r.Tassadar.Pop)
	default:
		return ""
	}
}

func (r Route) GetDveFromRoute(mode string) string {
	switch mode {
	case "nexus":
		return fmt.Sprintf("%s/nexus/%s?", r.InitURL, r.Nexus.Dve.Pool)
	case "valor":
		return fmt.Sprintf("%s/valor/%s?", r.InitURL, r.Valor.Dve.Pool)
	case "watsons":
		return fmt.Sprintf("%s/watsons/%s?page=1&pageSize=%v&", r.InitURL, r.Watsons.Dve.Pool, r.Watsons.Pse)
	case "watsonsha":
		return fmt.Sprintf("%s/watsons_ha/%s?", r.InitURL, r.WatsonsHa.Dve.Pool)
	case "tassadar":
		return fmt.Sprintf("%s/tassadar/%s?", r.InitURL, r.Tassadar.Dve.Pool)
	default:
		return ""
	}
}

func (r Route) GetDveRemoteFromRoute(mode string) string {
	switch mode {
	case "nexus":
		return fmt.Sprintf("%s/nexus/%s/%%v/%s", r.InitURL, r.Nexus.Dve.Pool, r.Nexus.Dve.Remote)
	case "valor":
		return fmt.Sprintf("%s/valor/%s/%%v/%s", r.InitURL, r.Valor.Dve.Pool, r.Valor.Dve.Remote)
	case "watsons":
		return fmt.Sprintf("%s/watsons/%%v/%v/%s", r.InitURL, r.Watsons.Dve.Pool, r.Watsons.Dve.Remote)
	case "watsonsha":
		return fmt.Sprintf("%s/watsons_ha/%%v/%v/%s", r.InitURL, r.WatsonsHa.Dve.Pool, r.WatsonsHa.Dve.Remote)
	case "tassadar":
		return fmt.Sprintf("%s/tassadar/%s/%%v/%s", r.InitURL, r.Tassadar.Dve.Pool, r.Tassadar.Dve.Remote)
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