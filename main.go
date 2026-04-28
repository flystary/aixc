package main

import (
	"runtime"
	"strings"

	"aixc/cmd"
	. "aixc/tools/arp"
	log "aixc/tools/log"
	. "aixc/tools/utmp"
)

func main() {

	os := runtime.GOOS
	arch := runtime.GOARCH
	xlog, _ := log.New("/var/run/aixc.log")

	if os == "linux" && arch == "amd64" {
		var user, tty, host, mac string
		// var utmps = make([]*Utmp, 0)
		var entrys = make(Entries, 0)

		utmps, err := LoadUtmp()
		if err != nil {

		}
		entrys, _ = GetEntries()

		for _, utmp := range utmps {
			user = strings.Trim(string(utmp.User[:]), "\x00")
			tty = strings.Trim(string(utmp.Device[:]), "\x00")
			host = strings.Trim(string(utmp.Host[:]), "\x00")

			mac, _ = entrys.GetMACFromAddr(host)
			xlog.Debugf("user:%s tty:%s host:%s mac:%s", user, tty, host, mac)
		}
	}
	cmd.Run()
}
