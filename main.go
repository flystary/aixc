package main

import (
	"aixc/cmd"
	arp "aixc/utils/arp"
	log "aixc/utils/log"
	utmp "aixc/utils/utmp"
	"strings"
)

func main() {
	var user, tty, host, mac string
	var arps = make(arp.Entrys, 0)
	utmps := utmp.LoadUtmp()

	arps, _ = arp.GetEntries()
	for _, utmp := range utmps {
		user = strings.Trim(string(utmp.User[:]), "\x00")
		tty  = strings.Trim(string(utmp.Device[:]), "\x00")
		host = strings.Trim(string(utmp.Host[:]), "\x00")

		mac, _ = arps.GetMACFromAddr(host)
		log.Debugf("user:%s tty:%s host:%s mac:%s", user, tty, host, mac)
	}
	cmd.Run()
}
