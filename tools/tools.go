package tools

import (
	"time"

	"github.com/mitchellh/go-homedir"
)

// HOME 用户根路径
func HOME() string {
	dir, _ := homedir.Dir()
	return dir
}

func TimeUnix(e time.Time) int64 {
	return e.UnixNano() / 1e6
}
