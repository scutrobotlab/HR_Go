package common

import (
	"github.com/zeromicro/go-zero/core/conf"
	"os"
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
	err := conf.Load(path, &config)
	return config, err
}
