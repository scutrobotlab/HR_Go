package sms

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

type SmsBaoConfig struct {
	SendUrl  string
	Username string
	Password string
	Debug    bool
}

func SendSms(config SmsBaoConfig, phone string, content string) error {
	if config.Debug {
		fmt.Printf("%s: --- send sms to %s: %s\n", time.Now().Format(time.DateTime), phone, content)
		return nil
	}

	encodedContent := url.QueryEscape(content)
	response, err := http.Get(fmt.Sprintf("%s?u=%s&p=%s&m=%s&c=%s", config.SendUrl, config.Username, config.Password, phone, encodedContent))
	if err != nil {
		return err
	}
	defer response.Body.Close()
	respBody, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}

	respStr := string(respBody)
	switch respStr {
	case "0":
		return nil
	case "30":
		return fmt.Errorf("密码错误")
	case "40":
		return fmt.Errorf("账号不存在")
	case "41":
		return fmt.Errorf("余额不足")
	case "43":
		return fmt.Errorf("IP地址限制")
	case "50":
		return fmt.Errorf("内容含有敏感词")
	case "51":
		return fmt.Errorf("手机号码不正确")
	default:
		return fmt.Errorf("未知错误")
	}
}
