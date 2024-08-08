package database

import (
	"bigTQQbot/conf"
	"fmt"
	"os"
)

func InitLocalData() {
	filePath, _ := os.Getwd()
	if _, err := os.Stat(filePath + "\\database\\local"); os.IsNotExist(err) {
		err := os.Mkdir(filePath+"\\database\\local", 0755)
		if err != nil {
			panic(err)
		}
	}
	if _, err := os.Stat(filePath + conf.BotBasicConfig.LocalData + "photos"); os.IsNotExist(err) {
		err := os.Mkdir(filePath+conf.BotBasicConfig.LocalData+"photos", 0755)
		if err != nil {
			panic(err)
		}
	}
	if _, err := os.Stat(filePath + conf.BotBasicConfig.LocalData + "order_helper.txt"); os.IsNotExist(err) {
		file, err := os.Create(filePath + conf.BotBasicConfig.LocalData + "order_helper.txt")
		if err != nil {
			panic(err)
		}
		defer file.Close()

		err = os.WriteFile("."+conf.BotBasicConfig.LocalData+"order_helper.txt", []byte(conf.BotBasicConfig.OrderHelperData), 0644)
		if err != nil {
			panic(err)
		}
	}

	fmt.Println("LocalDB loaded")
}
