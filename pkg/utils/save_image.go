package utils

import (
	"bigTQQbot/conf"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

func SaveImage(imgUrl string) (string, error) {
	urlSlice := strings.Split(imgUrl, "/")
	imgName := urlSlice[len(urlSlice)-1]
	localUrl, _ := os.Getwd()
	localUrl = strings.Replace(localUrl, "\\", "/", -1) + conf.BotBasicConfig.LocalPic + imgName
	_, err := os.Stat(localUrl)
	if err == nil {
		return localUrl, nil
	}

	client := &http.Client{}
	resp, err := client.Get(imgUrl)
	if err != nil {
		return "", err
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	fmt.Println(string(data))
	if string(data) == "404 Not Found" {
		return "", nil
	}
	err = ioutil.WriteFile(localUrl, data, 0666)
	if err != nil {
		return "", err

	}
	return localUrl, nil

}

func SaveImageWithProxy(imgUrl string) (string, error) {
	urlSlice := strings.Split(imgUrl, "/")
	imgName := urlSlice[len(urlSlice)-1]
	localUrl, _ := os.Getwd()
	localUrl = strings.Replace(localUrl, "\\", "/", -1) + conf.BotBasicConfig.LocalPic + imgName
	_, err := os.Stat(localUrl)
	if err == nil {
		return localUrl, nil
	}

	proxyUrl, err := url.Parse(fmt.Sprintf("http://localhost:%s", conf.BotBasicConfig.ProxyPort))
	if err != nil {
		fmt.Println("Error parsing")
		return "", err
	}
	client := &http.Client{
		Timeout: 20 * time.Second,
		Transport: &http.Transport{
			Proxy: http.ProxyURL(proxyUrl),
		},
	}
	resp, err := client.Get(imgUrl)
	if err != nil {
		fmt.Println("Error Get")
		return "", err
	}
	if resp.StatusCode != 200 {
		return "", nil
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error Reading")
		return "", err
	}
	err = ioutil.WriteFile(localUrl, data, 0666)
	if err != nil {
		fmt.Println("Error Writing")
		return "", err
	}
	return localUrl, nil
}
