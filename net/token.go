package net


import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sync"
	"crypto/md5"
	"encoding/hex"
)

// token单例
var onces  = &sync.Once{}

const (
    username string = "matrix"
    password string = "pXQL4m"
)

func newMD5(code string) string {
	MD5 := md5.New()
	_, _ = io.WriteString(MD5, code)
	return hex.EncodeToString(MD5.Sum(nil))
}

// GetToken 获取token
func GetToken(URL string) string {
	var Token string
	var result = make(map[string]interface{})
	var reData = make(url.Values)

	reData["username"] = []string{username}
	reData["password"] = []string{newMD5(newMD5(password))}
	reData["client_id"] = []string{"browser"}
	reData["client_secret"] = []string{"b7n3i7kzg22y3p035rw3rd9sfzvs4cv0"}
	reData["grant_type"] = []string{"password"}

	onces.Do(func() {
		res, err := http.PostForm(URL, reData)
		if err != nil {
			fmt.Printf("Login Error: %v", err)
		}
		defer res.Body.Close()
		body, err := io.ReadAll(res.Body)
		if err != nil {
			fmt.Printf("ReadAll IO Error: %v", err)
		}
		err = json.Unmarshal(body, &result)
		if err != nil {
			fmt.Printf("Unmarshal body Error: %v", err)
		}
		Token = result["access_token"].(string)
	})
	return Token
}
