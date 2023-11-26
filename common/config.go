package common

import (
	"fmt"
	"github.com/zeromicro/go-zero/core/conf"
	"os"
	"os/exec"
)

type Config struct {
	Debug bool
	Url   struct {
		BaseUrl   string
		AdminUrl  string
		WechatUrl string
	}
	Groups  []string
	Intents struct {
		Min      int32
		Max      int32
		Parallel bool
	}
	OAuth struct {
		ClientId     string
		ClientSecret string
		RedirectUrl  string
	}
}

func GetConfig(path string) (Config, error) {
	var config Config
	pathEnv := os.Getenv("COMMON_CONFIG_PATH")
	if pathEnv != "" {
		path = pathEnv
	}
	// 如果配置文件不存在，则从配置文件模板复制
	if _, err := os.Stat(path); os.IsNotExist(err) {
		command := exec.Command("cp", fmt.Sprintf("%s.example", path), path)
		err = command.Run()
		if err != nil {
			return config, err
		}
	}
	err := conf.Load(path, &config)
	return config, err
}
