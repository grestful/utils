package utils

import (
	//idworker "github.com/gitstliu/go-id-worker"
	"github.com/sony/sonyflake"

	//"time"
)

//var IdWorker *idworker.IdWorker
var sf *sonyflake.Sonyflake
func init() {
	var st sonyflake.Settings
	sf = sonyflake.NewSonyflake(st)
	if sf == nil {
		panic("sonyflake not created")
	}
	//IdWorker = &idworker.IdWorker{}
	//_ = IdWorker.InitIdWorker(int64(IPStringToInt(GetLocalIP())), time.Now().Unix()/1e6)
}

func NextId() int64 {
	id, _ := sf.NextID()
	return int64(id)
	//id, _ := IdWorker.NextId()
	//return id
}
