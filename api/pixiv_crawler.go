package api

import (
	"bigTQQbot/pkg/routes"
	"bigTQQbot/pkg/utils"
	"errors"
	"fmt"
	"sync"
)

func PixivRankListCrawler(tag string) ([]string, error) {
	var wg sync.WaitGroup
	var concurrencyError error

	if tag != "daily" && tag != "weekly" && tag != "monthly" && tag != "daily_r18" && tag != "weekly_r18" {
		return nil, errors.New("unknown mode param, must be 'daily','weekly','monthly','daily_r18','weekly_r18' or 'monthly'")
	}
	url := fmt.Sprintf("https://www.pixiv.net/ranking.php?mode=%s&content=illust", tag)
	data, err := routes.GetPixivListRequest(url)
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(data); i++ {
		wg.Add(1)
		go func(j int) {
			defer wg.Done()
			imgLocalUrl, err := utils.SaveImageWithProxy(data[j])
			if err != nil {
				concurrencyError = err
			}
			data[j] = imgLocalUrl
		}(i)
	}
	wg.Wait()
	if concurrencyError != nil {
		return nil, concurrencyError
	}
	return data, nil
}
