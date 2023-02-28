package main

import (
	"aixc/cmd"
	utmp "aixc/utils/utmp"
	log  "aixc/utils/log"
)

func main() {
	utmps := utmp.LoadUtmp()
	for _, utmp := range utmps {
           log.Debugf("user:%s tty:%s host:%s", string(utmp.User[:]), string(utmp.Device[:]), string(utmp.Host[:]))
	}
	cmd.Run()
}
