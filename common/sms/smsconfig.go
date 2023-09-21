package sms

import (
	"fmt"
	"strconv"
	"strings"
)

type Config struct {
	Signature string
	Typ       string
	Condition string
	Template  string
}

func (c Config) GetContent(vars ...string) string {
	content := fmt.Sprintf("【%s】", c.Signature)
	content += c.Template
	for i, s := range vars {
		content = strings.Replace(content, "${"+strconv.Itoa(i)+"}", s, -1)
	}
	return content
}

const Signature = "华工机器人实验室"

var Types = []string{"选择时间", "面试邀请1", "面试邀请2", "录取结果"}

var (
	Configs = map[string]map[string]Config{
		"选择时间": {
			"": {
				Signature: Signature,
				Typ:       "选择时间",
				Condition: "",
				Template: "${0}同学你好，欢迎你参加实验室招新。" +
					"我们看到你尚未选择${1}的面试时间，请看到短信后在报名系统中选择面试时间，以免错过面试。" +
					"放弃面试请忽略该短信。",
			},
		},
		"面试邀请": {
			"online": {
				Signature: Signature,
				Typ:       "面试邀请",
				Condition: "online",
				Template: "${0}同学，欢迎你参加${1}组面试。" +
					"你的面试时间是${2}，形式为飞书会议。" +
					"你的会议号是${3}，你可以下载飞书软件参会，或在面试系统完成页点击启动面试入会。" +
					"请确保摄像头和麦克风正常工作，并提前5分钟进入会议。" +
					"入会时你在等候室，请耐心等候面试官邀请你进入会议，谢谢！" +
					"如因个人原因取消面试，请在面试系统点击放弃面试，确认面试请回复“确认面试”。",
			},
			"offline": {
				Signature: Signature,
				Typ:       "面试邀请",
				Condition: "offline",
				Template: "${0}同学，欢迎你参加${1}组面试。" +
					"你的面试时间是${2}，形式为线下面试。" +
					"请提前10分钟到达${3}签到，谢谢！" +
					"如因个人原因取消面试，请在面试系统点击放弃面试，确认面试请回复“确认面试”。",
			},
		},
		"录取结果": {
			// 通过组别1
			"pass.group1": {
				Signature: Signature,
				Typ:       "录取结果",
				Condition: "pass.group1",
				Template: "${1}同学，你好！" +
					"恭喜你通过华工机器人实验室招新面试，成为${2}实习生。" +
					"希望在2024赛季，我们能够共同进步，突破创新，造出令人惊艳的机器人。" +
					"飞书是发放实习期通知的唯一工具，请登录面试系统查看邀请链接并加入飞书，同时尽快下载飞书客户端登录账号。" +
					"邀请链接于9月20日22:00失效，逾期不加入视为放弃offer。" +
					"收到请回复“华南虎，不要怂，就是干！”。",
			},
			// 通过组别2
			"pass.group2": {
				Signature: Signature,
				Typ:       "录取结果",
				Condition: "pass.group2",
				Template: "${1}同学，你好！" +
					"恭喜你通过华工机器人实验室招新面试，成为${2}实习生。" +
					"希望在2024赛季，我们能够共同进步，突破创新，成为实验室的重要支持力量。" +
					"飞书是发放实习期通知的唯一工具，请登录面试系统查看邀请链接并加入飞书，同时尽快下载飞书客户端登录账号。" +
					"邀请链接于9月20日22:00失效，逾期不加入视为放弃offer。" +
					"收到请回复“华南虎，不要怂，就是干！”。",
			},
			// 不通过不是大一新生
			"fail.not_freshman": {
				Signature: Signature,
				Typ:       "录取结果",
				Condition: "fail.not_freshman",
				Template: "${1}同学，你好！很遗憾，你没有通过本次机器人实验室的面试。" +
					"但这并不代表你不够优秀，而是我们希望你可以去探寻更适合你的天地。" +
					"非常感谢你参加本次面试，如果心存遗憾，期待在下一次招新的时候，可以再见到你的身影。"},
			// 不通过大一新生
			"fail.freshman": {
				Signature: Signature,
				Typ:       "录取结果",
				Condition: "fail.freshman",
				Template: "${1}同学，你好！很遗憾，你没有通过本次机器人实验室的面试。" +
					"但综合考虑你的发展潜力，我们推荐你加入机器人协会，" +
					"招新公告将在“2024赛季华南虎招新群”中同步发布，期待在机器人协会招新时，可以再见到你的身影。",
			},
		},
	}
)
