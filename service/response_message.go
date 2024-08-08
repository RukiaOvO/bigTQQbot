package service

import (
	"bigTQQbot/api"
	"bigTQQbot/conf"
	"bigTQQbot/pkg/utils"
	"bigTQQbot/types"
	"fmt"
	"strings"
)

func PrivateMessageResponse(msg types.BotPostMessageData) {
	return
}

func GroupMessageResponse(msg types.BotPostMessageData) {
	orderMsgList := strings.Split(msg.RawMessage, " ")
	if orderMsgList[0] == conf.BotBasicConfig.AtSelf {
		order, ok := utils.OrderCheck(strings.Join(orderMsgList[1:], " "))
		if ok {
			OrderMessageController(msg, order)
		} else {
			WrongOrderMessageController(msg)
		}
	} else {
		NormalMessageController()
	}
}

func OrderMessageController(msg types.BotPostMessageData, op string) {
	respMsgHead := fmt.Sprintf("[CQ:reply,id=%v][CQ:at,qq=%v]", msg.MessageID, msg.UserID)

	if op == "/like" {
		LikeOrderController(respMsgHead, msg)
	} else if op == "/gpt" {
		GptOrderController(respMsgHead, msg)
	} else if op == "/help" {
		HelpOrderController(respMsgHead, msg)
	} else if op == "/setu" {
		SetuOrderController(respMsgHead, msg)
	} else if op == "/setu_rank" {
		SetuRankOrderController(respMsgHead, msg)
	} else if op == "/pixivlist" {
		PixivOrderController(respMsgHead, msg)
	} else {
		WrongOrderMessageController(msg)
	}
}

func WrongOrderMessageController(botMsg types.BotPostMessageData) {
	msg := fmt.Sprintf("[CQ:reply,id=%v][CQ:at,qq=%v]", botMsg.MessageID, botMsg.UserID)
	if botMsg.RawMessage == conf.BotBasicConfig.AtSelf {
		msg += "¿"
	} else {
		msg += "泥在嗦甚么，使用/help获取指令帮助"
	}
	api.GroupMessageSender(botMsg, msg)
}

func NormalMessageController() {
}
