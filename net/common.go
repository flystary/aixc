package net

import (
	. "aixc/tools"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// T ~ cpe.Valor, pop.Watsons dev.Watsons
func FetchData[T any](TOKEN, URL string, target *T) error {
	Unix := TimeUnix(time.Now())
	fullURL := fmt.Sprintf("%s?access_token=%s&_=%d", URL, TOKEN, Unix)

	res, err := http.Get(fullURL)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	bytes, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(bytes, target)
}
