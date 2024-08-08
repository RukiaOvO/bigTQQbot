package api

import (
	"bigTQQbot/pkg/routes"
	"bigTQQbot/types"
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

func PrivateMessageSender(id int64, msg string) {
	data, err := json.Marshal(map[string]interface{}{
		"user_id":     id,
		"message":     msg,
		"auto_escape": false,
	})
	if err != nil {
		log.Fatal(err)
	}
	err = routes.PostRequest("send_private_msg", strings.NewReader(string(data)))
	if err != nil {
		log.Fatal(err)
	}
}

func GroupMessageSender(botMsg types.BotPostMessageData, msg string) {
	data, _ := json.Marshal(map[string]interface{}{
		"group_id":    botMsg.GroupID,
		"message":     msg,
		"auto_escape": false,
	})
	err := routes.PostRequest("send_group_msg", strings.NewReader(string(data)))
	if err != nil {
		errData, _ := json.Marshal(map[string]interface{}{
			"group_id":    botMsg.GroupID,
			"message":     fmt.Sprintf("[CQ:reply,id=%v][CQ:at,qq=%v]%v", botMsg.MessageID, botMsg.UserID, err.Error()),
			"auto_escape": false,
		})
		err := routes.PostRequest("send_group_msg", strings.NewReader(string(errData)))
		fmt.Println("Error :", err)
	}
}
