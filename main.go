package main

import (
	"log"
	"os"
	"time"
	"aixc/cmd"
	"aixc/utils"
)

var loger *log.Logger

func init() {
        file := "/var/log/xc" + time.Now().Format("20180102") + ".log"
        logFile, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
        if err != nil {
                panic(err)
        }
        loger = log.New(logFile, "[xc]",log.Ldate|log.Ltime|log.LstdFlags|log.Lshortfile|log.LUTC) // 将文件设置为loger作为输出
}

func main() {
		utmps := utils.LoadUtmp()
		for _, utmp := range utmps {
			loger.Println(string(utmp.User[:]), string(utmp.Device[:]),string(utmp.Host[:]))
		}
		cmd.Run()
}