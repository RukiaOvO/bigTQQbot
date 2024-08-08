package api

import (
	"bigTQQbot/conf"
	"bigTQQbot/pkg/routes"
	"bigTQQbot/pkg/utils"
	"encoding/json"
	"strings"
	"sync"
)

func SendSetu(tag []string) ([]string, error) {
	data, err := json.Marshal(map[string]interface{}{
		"tag": tag,
		"r18": conf.BotPluginConfig.SeSe.SeTuRank,
		"num": conf.BotPluginConfig.SeSe.SeTuNum,
	})
	if err != nil {
		return nil, err
	}
	urls, err := routes.PostSetuRequest(strings.NewReader(string(data)))
	if err != nil {
		return nil, err
	}
	var wg sync.WaitGroup
	var concurrencyError error
	wg.Add(len(urls))
	for i := 0; i < len(urls); i++ {
		go func(j int) {
			defer wg.Done()
			imgLocalUrl, err := utils.SaveImageWithProxy(urls[j])
			if err != nil {
				concurrencyError = err
			}
			urls[j] = imgLocalUrl
		}(i)
	}
	wg.Wait()
	if concurrencyError != nil {
		return nil, concurrencyError
	}
	return urls, nil
}
