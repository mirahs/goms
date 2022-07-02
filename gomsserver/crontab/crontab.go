package crontab

import (
	"github.com/robfig/cron/v3"
	"log"
	"runtime"
)


var cs *cron.Cron


func Init() {
	cs = cron.New(cron.WithSeconds())

	hostInitCs()

	cs.Start()
}


func recoverCheck() {
	if err := recover(); err != nil {
		buf := make([]byte, 4096)
		l := runtime.Stack(buf, false)
		log.Printf("%v: %s", err, buf[:l])
	}
}
