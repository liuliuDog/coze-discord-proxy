package discord

import (
	"coze-discord-proxy/common"
	"os"
)

var MultiBotConfig = os.Getenv("MULTI_BOT_CONFIG")

func loadMultiBotConfig() []byte {
    if MultiBotConfig != "" {
	common.SysLog("读取到配置,", MultiBotConfig)
	return []byte(MultiBotConfig)
    }
}
