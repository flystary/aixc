package net

import (
	"net/http"
	"strings"
)

func Put(URL string, Payload *strings.Reader) *http.Response {
	req, _ := http.NewRequest("PUT", URL, Payload)
	req.Header.Add("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	return resp
}
