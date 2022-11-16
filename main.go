package main

import (
	"aixc/cmd"
	"aixc/utils/logger"
	"aixc/utils/utmp"
)

func main() {
	utmps := utmp.LoadUtmp()
	for _, utmp := range utmps {
	    logger.Debugf("user:%s tty:%s host:%s", string(utmp.User[:]), string(utmp.Device[:]), string(utmp.Host[:]))
	}
	cmd.Run()
}
