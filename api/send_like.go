package api

import (
	"bigTQQbot/pkg/routes"
	"encoding/json"
	"strings"
)

func HomePageLikeSender(id int64, times int) error {
	data, err := json.Marshal(map[string]interface{}{
		"user_id": id,
		"times":   times,
	})
	if err != nil {
		return err
	}
	return routes.PostRequest("send_like", strings.NewReader(string(data)))
}
