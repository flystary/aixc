package net

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sync"
)

// token单例
var onces = &sync.Once{}

const (
	username string = "xxxxxxxxxxxxxxxx"
	password string = "xxxxxxxxxxxxxxxx"
)

func newMD5(code string) string {
	m := md5.New()
	// m.Write([]byte(code))
	io.WriteString(m, code)
	return hex.EncodeToString(m.Sum(nil))
}

type tokenResp struct {
	AccessToken string `json:"access_token"`
}

// GetToken 获取token
func GetToken(URL string) (string, error) {
	data := url.Values{}
	data.Set("username", username)
	data.Set("password", newMD5(newMD5(password)))
	data.Set("client_id", "browser")
	data.Set("client_secret", "b7n3i7kzg22y3p035rw3rd9sfzvs4cv0")
	data.Set("grant_type", "password")

	res, err := http.PostForm(URL, data)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	var resp tokenResp
	if err := json.Unmarshal(body, &resp); err != nil {
		return "", err
	}

	if resp.AccessToken == "" {
		return "", fmt.Errorf("empty token")
	}

	return resp.AccessToken, nil
}
