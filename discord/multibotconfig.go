package discord

import (
	"context"
	"coze-discord-proxy/common"
	"os"
)

var MultiBotConfig = os.Getenv("MULTI_BOT_CONFIG")

func WriteBotConfig(config string) error {
    // 检查文件是否存在
    _, err := os.Stat("config/bot_config.json")
    if err != nil {
        if os.IsNotExist(err) {
            // 文件不存在，创建文件
            file, err := os.Create("config/bot_config.json")
            if err != nil {
                return err
            }
            defer file.Close()
        } else {
            // 其他错误
            return err
        }
    }

    common.SysLog("文件创建成功.")
    // 将字符串写入文件
    return os.WriteFile("config/bot_config.json", []byte(config), 0644)
}

func loadMultiBotConfig() {
    if MultiBotConfig != "" {
	if err := WriteBotConfig(MultiBotConfig); err != nil {
            common.FatalLog("配置文件写入失败,", err)
        } else {
            common.SysLog("配置文件写入成功.")
        }
    }
}
