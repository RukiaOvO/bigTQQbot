package routes

import (
	"bigTQQbot/conf"
	"bigTQQbot/pkg/utils"
	"bigTQQbot/types"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"time"
)

func PostRequest(api string, data *strings.Reader) error {
	postUrl := conf.BotBasicConfig.PostUrl + api
	req, _ := http.NewRequest("POST", postUrl, data)
	req.Header.Add("Content-Type", "application/json")

	//proxyUrl, _ := url.Parse(fmt.Sprintf("http://127.0.0.1:%s", strconv.Itoa(consts.ProxyPort)))
	//client := &http.Client{
	//	Timeout: 15 * time.Second,
	//	Transport: &http.Transport{
	//		Proxy: http.ProxyURL(proxyUrl),
	//	},
	//}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var respData types.PostMessageResp
	body, _ := io.ReadAll(resp.Body)
	_ = json.Unmarshal(body, &respData)
	if respData.Status != "ok" {
		return errors.New(respData.Message)
	}
	return nil
}

func PostGptRequest(data *strings.Reader) (string, error) {
	req, _ := http.NewRequest("POST", conf.BotPluginConfig.Gpt.GptUrl, data)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+conf.BotPluginConfig.Gpt.GptApiKey)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var respData types.GPT3Dot5Resp
	body, _ := io.ReadAll(resp.Body)
	err = json.Unmarshal(body, &respData)
	if err != nil {
		return "", err
	}
	return respData.Choices[0].Message.Content, nil
}

func PostSetuRequest(data *strings.Reader) ([]string, error) {
	req, _ := http.NewRequest("POST", conf.BotPluginConfig.SeSe.SeTuApi, data)
	req.Header.Add("Content-Type", "application/json")
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var imgUrls []string
	var respData types.SetuResp
	body, _ := io.ReadAll(resp.Body)
	err = json.Unmarshal(body, &respData)
	if err != nil {
		return nil, err
	}
	if len(respData.Data) == 0 {
		return nil, errors.New("image not found")
	}
	for _, v := range respData.Data {
		imgUrls = append(imgUrls, v.Urls.Original)
	}
	return imgUrls, nil
}

func GetPixivListRequest(pageUrl string) ([]string, error) {
	proxyUrl, err := url.Parse(fmt.Sprintf("http://localhost:%s", conf.BotBasicConfig.ProxyPort))
	client := &http.Client{
		Timeout: 20 * time.Second,
		Transport: &http.Transport{
			Proxy: http.ProxyURL(proxyUrl),
		},
	}

	req, err := http.NewRequest("GET", pageUrl, nil)
	req.Header.Set("Host", conf.BotPluginConfig.Pixiv.Host)
	req.Header.Set("User-Agent", conf.BotPluginConfig.Pixiv.UserAgent)
	req.Header.Set("Cookie", conf.BotPluginConfig.Pixiv.Cookie)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data []string
	body, err := ioutil.ReadAll(resp.Body)
	html := strings.Replace(string(body), "\n", "", -1)
	reImg := regexp.MustCompile(`data-src="([^"]*)"data-type`)
	data = reImg.FindAllString(html, 20)
	if len(data) == 0 {
		return nil, errors.New("image not found")
	}
	for c, v := range data {
		data[c] = utils.PixivUrlRebuild(strings.Split(v, `"`)[1])
	}
	return data, nil
}
