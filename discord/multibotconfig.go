package discord

import (
	"coze-discord-proxy/common"
	"os"
	"fmt"
)

var MultiBotConfig = os.Getenv("MULTI_BOT_CONFIG")

func loadMultiBotConfig() []byte {
    common.SysLog(fmt.Sprintf("读取配置 MultiBotConfig: %+v", MultiBotConfig))
    return []byte(MultiBotConfig)
}
