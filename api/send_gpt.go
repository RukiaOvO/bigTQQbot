package api

import (
	"bigTQQbot/conf"
	"bigTQQbot/pkg/routes"
	"encoding/json"
	"fmt"
	"strings"
)

func GptSender(q string) (string, error) {
	data, _ := json.Marshal(map[string]interface{}{
		"model": conf.BotPluginConfig.Gpt.GptModel,
		"messages": []map[string]interface{}{{
			"role":    "system",
			"content": "You are a helpful assistant.",
		}, {
			"role":    "user",
			"content": q,
		}},
	})
	fmt.Println(string(data))
	return routes.PostGptRequest(strings.NewReader(string(data)))
}
