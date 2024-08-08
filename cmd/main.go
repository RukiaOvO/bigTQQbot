package main

import (
	"bigTQQbot/conf"
	"bigTQQbot/database"
	"bigTQQbot/pkg/utils"
	"bigTQQbot/service"
	"bigTQQbot/types"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
)

func main() {
	conf.LoadBotConfig()
	database.InitLocalData()
	ReceivePostMessage()
}

func ReceivePostMessage() {
	http.HandleFunc("/", BotPostMessageController)
	fmt.Printf("Listen and serve%s\n", conf.BotBasicConfig.MessagePostPort)

	if err := http.ListenAndServe(conf.BotBasicConfig.MessagePostPort, nil); err != nil {
		fmt.Println("Error: ", err)
	}
}

func BotPostMessageController(_ http.ResponseWriter, r *http.Request) {
	var newMsg types.BotPostMessageData
	data, _ := io.ReadAll(r.Body)
	fmt.Println(string(data))
	if err := json.Unmarshal(data, &newMsg); err != nil {
		log.Fatal(err)
	}
	if newMsg.MessageType == "private" {
		service.PrivateMessageResponse(newMsg)
	} else if newMsg.MessageType == "group" && utils.StrIsContains(strconv.FormatInt(newMsg.GroupID, 10), conf.BotBasicConfig.ServeGroups) {
		service.GroupMessageResponse(newMsg)
	}
}
