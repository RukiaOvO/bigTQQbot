package service

import (
	"bigTQQbot/api"
	"bigTQQbot/conf"
	"bigTQQbot/types"
	"fmt"
	"io/ioutil"
	"strings"
)

func LikeOrderController(msg string, botMsg types.BotPostMessageData) {
	err := api.HomePageLikeSender(botMsg.Sender.UserID, 10)
	if err != nil {
		msg += err.Error()
	} else {
		msg += "已点赞十次![每日上限10个赞]"
	}
	api.GroupMessageSender(botMsg, msg)
}
func HelpOrderController(msg string, botMsg types.BotPostMessageData) {
	ans, err := ioutil.ReadFile("./database/local/order_helper.txt")
	if err != nil {
		msg += err.Error()
	} else {
		msg += string(ans)
	}
	api.GroupMessageSender(botMsg, msg)
}

func GptOrderController(msg string, botMsg types.BotPostMessageData) {
	orderMsgList := strings.Split(botMsg.RawMessage, " ")
	if len(orderMsgList) == 2 {
		msg += "请给出想问的问题"
	} else {
		ans, err := api.GptSender(strings.Join(orderMsgList[2:], " "))
		if err != nil {
			msg += err.Error()
		} else {
			msg += ans
		}
	}
	api.GroupMessageSender(botMsg, msg)
}

func SetuOrderController(msg string, botMsg types.BotPostMessageData) {
	orderMsgList := strings.Split(botMsg.RawMessage, " ")
	var tag []string
	if len(orderMsgList) > 2 {
		tag = orderMsgList[2:]
	}
	urls, err := api.SendSetu(tag)
	if err != nil {
		msg += err.Error()
	} else {
		for _, v := range urls {
			if v != "" {
				msg += fmt.Sprintf("[CQ:image,file=file://%s]\n", v)
			} else {
				msg += fmt.Sprintf("该图片已丢失\n")
			}
		}
	}
	fmt.Println(msg)
	api.GroupMessageSender(botMsg, msg)
}

func SetuRankOrderController(msg string, botMsg types.BotPostMessageData) {
	if conf.BotPluginConfig.SeSe.SeTuRank == 1 {
		conf.BotPluginConfig.SeSe.SeTuRank = 0
		msg += "已转换，目前为非R18"
	} else {
		conf.BotPluginConfig.SeSe.SeTuRank = 1
		msg += "已转换，目前为R18"
	}
	api.GroupMessageSender(botMsg, msg)
}

func PixivOrderController(msg string, botMsg types.BotPostMessageData) {
	orderMsgList := strings.Split(botMsg.RawMessage, " ")
	var urls []string
	var err error
	if len(orderMsgList) == 2 {
		urls, err = api.PixivRankListCrawler("daily")
	} else {
		urls, err = api.PixivRankListCrawler(orderMsgList[2])
	}
	if err != nil {
		msg += err.Error()
	} else {
		for _, v := range urls {
			if v != "" {
				msg += fmt.Sprintf("[CQ:image,file=file://%v]\n", v)
			}
		}
	}
	fmt.Println(msg)
	api.GroupMessageSender(botMsg, msg)
}
