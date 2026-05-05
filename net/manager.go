package net

import (
	"aixc/model/cpe"
	"aixc/model/dve"
	"aixc/model/pop"
	r "aixc/route"
	. "aixc/tools"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

type NetManager struct {
	Token       string
	Router      Router
	CPERegistry map[string]cpe.CpeProvider
	POPRegistry map[string]pop.PopProvider
	DVERegistry map[string]dve.DveProvider
	mu          sync.RWMutex
}

func NewNetManager(routePath string) (*NetManager, error) {
	route, err := r.LoadRoute(routePath)
	if err != nil {
		return nil, err
	}
	token, _ := GetToken(route.GetTokenFromRoute())
	return &NetManager{
		Token:       token,
		Router:      route,
		CPERegistry: make(map[string]cpe.CpeProvider),
		POPRegistry: make(map[string]pop.PopProvider),
		DVERegistry: make(map[string]dve.DveProvider),
	}, nil
}

func (m *NetManager) FetchAny(url string, target any) error {
	fullURL := fmt.Sprintf("%s?access_token=%s&_=%d", url, m.Token, TimeUnix(time.Now()))
	res, err := http.Get(fullURL)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	bytes, _ := io.ReadAll(res.Body)
	return json.Unmarshal(bytes, target)
}

func (m *NetManager) SyncData(mode string) error {
	m.mu.Lock()
	if _, ok := m.CPERegistry[mode]; !ok {
		m.registerType(mode)
	}
	m.mu.Unlock()

	urls := []string{m.Router.GetCpeFromRoute(mode), m.Router.GetPopFromRoute(mode), m.Router.GetDveFromRoute(mode)}
	targets := []any{m.CPERegistry[mode], m.POPRegistry[mode], m.DVERegistry[mode]}

	var wg sync.WaitGroup
	errChan := make(chan error, 3)

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(idx int) {
			defer wg.Done()
			if err := m.FetchAny(urls[idx], targets[idx]); err != nil {
				errChan <- err
			}
		}(i)
	}
	wg.Wait()
	close(errChan)
	if len(errChan) > 0 {
		return <-errChan
	}
	return nil
}

func (m *NetManager) registerType(mode string) {
	switch mode {
	case "valor", "yifeng":
		m.CPERegistry[mode] = &cpe.Valor{}
		m.POPRegistry[mode] = &pop.Valor{}
		m.DVERegistry[mode] = &dve.Valor{}
	case "watsons":
		m.CPERegistry[mode] = &cpe.Watsons{}
		m.POPRegistry[mode] = &pop.Watsons{}
		m.DVERegistry[mode] = &dve.Watsons{}
	case "watsonsha":
		m.CPERegistry[mode] = &cpe.WatsonsHa{}
		m.POPRegistry[mode] = &pop.WatsonsHa{}
		m.DVERegistry[mode] = &dve.WatsonsHa{}
	case "tassadar":
		m.CPERegistry[mode] = &cpe.Zeratul{}
		m.POPRegistry[mode] = &pop.Zeratul{}
		m.DVERegistry[mode] = &dve.Zeratul{}
	}
}

func (m *NetManager) SyncAll(modes []string) {
	var wg sync.WaitGroup
	for _, mode := range modes {
		wg.Add(1)
		go func(md string) {
			defer wg.Done()
			_ = m.SyncData(md)
		}(mode)
	}
	wg.Wait()
}
