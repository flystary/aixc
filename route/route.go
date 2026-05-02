package route

import (
	"fmt"
	"os"
	"sync"

	"gopkg.in/yaml.v3"
)

// Route
type Route struct {
	BaseURL       string              `yaml:"baseurl"`
	TokenPath     string              `yaml:"tokenpath"`
	OperationPath string              `yaml:"operationpath"`
	Modes         []string            `yaml:"modes"`
	Platforms     map[string]Classify `yaml:",inline"`
}

// Classify
type Classify struct {
	CpePath string `yaml:"cpe"`
	PopPath string `yaml:"pop"`
	DvePath Device `yaml:"dve"`
	Pse     *int   `yaml:"pse,omitempty"`
}

type Device struct {
	PoolPath   string `yaml:"pool"`
	RemotePath string `yaml:"remote"`
}

var (
	global  Route
	once    sync.Once
	errInit error
)

func LoadRoute(path string) (*Route, error) {
	once.Do(func() {
		b, err := os.ReadFile(path)
		if err != nil {
			errInit = err
			return
		}

		errInit = yaml.Unmarshal(b, &global)
	})

	return &global, errInit
}
func (r Route) getCfg(mode string) (Classify, bool) {
	c, ok := r.Platforms[mode]
	return c, ok
}

func (r Route) GetCpeFromRoute(mode string) string {
	c, ok := r.getCfg(mode)
	if !ok {
		return ""
	}

	baseURL := fmt.Sprintf("%s/%s/%s", r.BaseURL, mode, c.CpePath)
	if c.Pse != nil {
		return fmt.Sprintf("%s?page=1&pageSize=%d&", baseURL, *c.Pse)
	}
	return baseURL + "?"
}

func (r Route) GetPopFromRoute(mode string) string {
	c, ok := r.getCfg(mode)
	if !ok {
		return ""
	}
	return fmt.Sprintf("%s/%s/%s", r.BaseURL, mode, c.PopPath)
}

func (r Route) GetDveFromRoute(mode string) string {
	c, ok := r.getCfg(mode)
	if !ok {
		return ""
	}

	baseURL := fmt.Sprintf("%s/%s/%s", r.BaseURL, mode, c.DvePath.PoolPath)
	// 兼容 watsons 的特殊分页逻辑
	if mode == "watsons" && c.Pse != nil {
		return fmt.Sprintf("%s?page=1&pageSize=%d&", baseURL, *c.Pse)
	}
	return baseURL + "?"
}

func (r Route) GetDveRemoteFromRoute(mode string) string {
	c, ok := r.getCfg(mode)
	if !ok {
		return ""
	}

	return fmt.Sprintf("%s/%s/%s/%%v/%s", r.BaseURL, mode, c.DvePath.PoolPath, c.DvePath.RemotePath)
}

func (r Route) GetTokenFromRoute() string {
	return fmt.Sprintf("%s/%s?", r.BaseURL, r.TokenPath)
}

func (r Route) GetOperationFromRoute() string {
	return fmt.Sprintf("%s/%s?", r.BaseURL, r.OperationPath)
}
