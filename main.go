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

	for _, utmp := range utmps {
		user = string(utmp.User[:])
		tty = string(utmp.Device[:])
		host = string(utmp.Host[:])
		break
	}

	arps, _ = arp.GetEntries()
	str := strings.Trim(host, "\x00")
	mac, _ = arps.GetMACFromAddr(str)
	log.Debugf("user:%s tty:%s host:%s mac:%s", user, tty, host, mac)
	cmd.Run()
}
