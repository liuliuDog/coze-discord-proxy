package discord

import (
	"encoding/json"
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

    // 将字符串写入文件
    return os.WriteFile("config/bot_config.json", []byte(config), 0644)
}

func loadMultiBotConfig() {
	if MultiBotConfig != "" {
		if err := WriteBotConfig(MultiBotConfig); err != nil {
      fmt.Println("无法写入配置文件：", err)
    } else {
	  common.SysLog("Environment variable check passed.")
    }
	}
}
