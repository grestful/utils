package utils

import (
	"github.com/gitstliu/go-id-worker"
	"time"
)

var IdWorker *idworker.IdWorker
func init()  {
	IdWorker = &idworker.IdWorker{}
	_ = IdWorker.InitIdWorker(int64(IPStringToInt(GetLocalIP())), time.Now().Unix() / 1e6)
}

func NextId() int64{
	id,_ := IdWorker.NextId()
	return id
}
