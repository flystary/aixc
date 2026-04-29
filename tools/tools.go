package tools

import (
	"bytes"
	"time"

	"github.com/mitchellh/go-homedir"
)

// HOME 用户根路径
func HOME() string {
	dir, _ := homedir.Dir()
	return dir
}

func TrimNull(b []byte) string {
	if i := bytes.IndexByte(b, 0); i >= 0 {
		return string(b[:i])
	}
	return string(b)
}

func TimeUnix(e time.Time) int64 {
	return e.UnixNano() / 1e6
}
