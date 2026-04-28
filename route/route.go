package route

import (
	"fmt"
	"os"
	"sync"

	"gopkg.in/yaml.v3"
)

// Route
type Route struct {
	InitURL   string   `yaml:"initurl"`
	Tokenurl  string   `yaml:"tokenurl"`
	Operation string   `yaml:"operation"`
	Modes     []string `yaml:"modes"`

	Valor     Classify `yaml:"valor"`
	Nexus     Classify `yaml:"nexus"`
	Watsons   Classify `yaml:"watsons"`
	Tassadar  Classify `yaml:"tassadar"`
	WatsonsHa Classify `yaml:"watsonsha"`
}

// Classify
type Classify struct {
	Cpe string `yaml:"cpe"`
	Pop string `yaml:"pop"`
	Dve Device `yaml:"dve"`
	Pse *int   `yaml:"pse,omitempty"`
}

type Device struct {
	Pool   string `yaml:"pool"`
	Remote string `yaml:"remote"`
}

var (
	global  Route
	once    sync.Once
	errInit error
)

func LoadRoute(path string) (Route, error) {
	once.Do(func() {
		b, err := os.ReadFile(path)
		if err != nil {
			errInit = err
			return
		}

		errInit = yaml.Unmarshal(b, &global)
	})

	return global, errInit
}

func (r Route) get(mode string) (string, Classify, bool) {
	switch mode {
	case "nexus":
		return "nexus", r.Nexus, true
	case "valor":
		return "valor", r.Valor, true
	case "watsons":
		return "watsons", r.Watsons, true
	case "watsonsha":
		return "watsons_ha", r.WatsonsHa, true
	case "tassadar":
		return "tassadar", r.Tassadar, true
	default:
		return "", Classify{}, false
	}
}

func (r Route) build(mode, path string, extra string) string {
	name, _, ok := r.get(mode)
	if !ok {
		return ""
	}
	return fmt.Sprintf("%s/%s/%s%s", r.InitURL, name, path, extra)
}

func (r Route) GetCpeFromRoute(mode string) string {
	name, c, ok := r.get(mode)
	if !ok {
		return ""
	}

	switch mode {
	case "nexus", "tassadar":
		return fmt.Sprintf("%s/%s/%s?", r.InitURL, name, c.Cpe)
	default:
		return fmt.Sprintf("%s/%s/%s?page=1&pageSize=%v&", r.InitURL, name, c.Cpe, c.Pse)
	}
}

func (r Route) GetPopFromRoute(mode string) string {
	name, c, ok := r.get(mode)
	if !ok {
		return ""
	}
	return fmt.Sprintf("%s/%s/%s", r.InitURL, name, c.Pop)
}

func (r Route) GetDveFromRoute(mode string) string {
	name, c, ok := r.get(mode)
	if !ok {
		return ""
	}

	if mode == "watsons" {
		return fmt.Sprintf("%s/%s/%s?page=1&pageSize=%v&",
			r.InitURL, name, c.Dve.Pool, c.Pse)
	}

	return fmt.Sprintf("%s/%s/%s?", r.InitURL, name, c.Dve.Pool)
}

func (r Route) GetDveRemoteFromRoute(mode string) string {
	name, c, ok := r.get(mode)
	if !ok {
		return ""
	}

	return fmt.Sprintf(
		"%s/%s/%s/%%v/%s",
		r.InitURL,
		name,
		c.Dve.Pool,
		c.Dve.Remote,
	)
}

func (r Route) GetTokenFromRoute() string {
	return fmt.Sprintf("%s/%s?", r.InitURL, r.Tokenurl)
}

func (r Route) GetOperationFromRoute() string {
	return fmt.Sprintf("%s/%s?", r.InitURL, r.Operation)
}
