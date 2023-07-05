package initialize

import (
	"github.com/orca-zhang/ecache"
	"github.com/ve-weiyi/ve-admin-store/server/global"
	"time"
)

func OtherInit() {
	global.BlackCache = ecache.NewLRUCache(16, 200, 10*time.Second).LRU2(1024)
}
