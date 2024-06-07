package discord

import (
	"coze-discord-proxy/common"
	"os"
	"fmt"
)

var MultiBotConfig = os.Getenv("MULTI_BOT_CONFIG")

func loadMultiBotConfig() []byte {
    
    if MultiBotConfig != "" {
        common.SysLog(fmt.Sprintf("读取配置 MultiBotConfig: %+v", MultiBotConfig))
        return []byte(MultiBotConfig)
    }
    common.FatalLog("环境变量 MULTI_BOT_CONFIG 未设置")
    return
}
