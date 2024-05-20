package initialize

import (
	"time"

	"github.com/orca-zhang/ecache"

	"github.com/ve-weiyi/ve-blog-golang/kit/infra/chatgpt"
	"github.com/ve-weiyi/ve-blog-golang/server/global"
)

func OtherInit() {
	global.BlackCache = ecache.NewLRUCache(16, 200, 10*time.Second).LRU2(1024)

	gpt := chatgpt.NewAIChatGPT(
		chatgpt.WithApiKey(global.CONFIG.ChatGPT.ApiKey),
		chatgpt.WithApiHost(global.CONFIG.ChatGPT.ApiHost),
		chatgpt.WithModel(global.CONFIG.ChatGPT.Model),
	)

	global.AIChatGPT = gpt
}
