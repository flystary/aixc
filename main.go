package main

import (
	"runtime"

	"aixc/cmd"
	"aixc/tools"
	"aixc/tools/arp"
	xlog "aixc/tools/log"
	"aixc/tools/utmp"
)

func main() {

	goos := runtime.GOOS
	arch := runtime.GOARCH
	err := xlog.Init("/var/run/aixc.log")
	if err != nil {
		panic(err)
	}
	defer xlog.Close()

	if goos == "linux" && arch == "amd64" {
		utmps, err := utmp.LoadUtmp()
		if err != nil {
			xlog.Errorf("load utmp failed : %v", err)
			return
		}

		entrys, err := arp.GetEntries()
		if err != nil {
			xlog.Errorf("load arp failed : %v", err)
			return
		}

		for _, u := range utmps {
			user := tools.TrimNull(u.User[:])
			tty := tools.TrimNull(u.Device[:])
			host := tools.TrimNull(u.Host[:])

			mac, err := entrys.GetMACFromAddr(host)
			if err != nil {
				mac = ""
			}
			xlog.Debugf("user:%s tty:%s host:%s mac:%s", user, tty, host, mac)
		}
	}
	cmd.Run()
}
