package dao

import (
	"gomsserver/util/time"
	"sync"
)


var hostOnlineMutex sync.Mutex
var hostOnline = make(map[string]int64)


func HostOnlineHeart(host string) {
	hostOnlineMutex.Lock()
	defer hostOnlineMutex.Unlock()

	hostOnline[host] = time.UnixTime()
}

func HostOnlineGet(host string) bool {
	hostOnlineMutex.Lock()
	defer hostOnlineMutex.Unlock()

	heart := hostOnline[host]

	return (time.UnixTime() - heart) <= 8
}
