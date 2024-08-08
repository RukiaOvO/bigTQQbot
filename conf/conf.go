package conf

import (
	"fmt"
	"github.com/spf13/viper"
)

type BasicConfig struct {
	ServeGroups     []string `mapstructure:"serve_groups"`
	AtSelf          string   `mapstructure:"at_self"`
	PostUrl         string   `mapstructure:"post_url"`
	MessagePostPort string   `mapstructure:"message_post_port"`
	ProxyPort       string   `mapstructure:"proxy_port"`
	LocalData       string   `mapstructure:"local_data"`
	LocalPic        string   `mapstructure:"local_pic"`
	OrderHelperData string   `mapstructure:"order_helper_data"`
}
type PluginConfig struct {
	SeSe  SeSeConfig  `mapstructure:"sese"`
	Pixiv PixivConfig `mapstructure:"pixiv"`
	Gpt   GptConfig   `mapstructure:"gpt"`
}
type SeSeConfig struct {
	SeTuApi  string `mapstructure:"setu_api"`
	SeTuRank int    `mapstructure:"setu_rank"`
	SeTuNum  int    `mapstructure:"setu_num"`
}
type PixivConfig struct {
	Cookie    string `mapstructure:"cookie"`
	Host      string `mapstructure:"host"`
	UserAgent string `mapstructure:"user_agent"`
}
type GptConfig struct {
	GptModel  string `mapstructure:"gpt_model"`
	GptApiKey string `mapstructure:"gpt_apikey"`
	GptUrl    string `mapstructure:"gpt_url"`
}

var BotBasicConfig = BasicConfig{}
var BotPluginConfig = PluginConfig{}

func LoadBotConfig() {
	reader := viper.New()
	reader.AddConfigPath("./conf/local")
	reader.SetConfigType("yaml")
	reader.SetConfigName("config")
	if err := reader.ReadInConfig(); err != nil {
		panic(err)
	}
	if err := reader.UnmarshalKey("BasicConfig", &BotBasicConfig); err != nil {
		panic(err)
	}
	if err := reader.UnmarshalKey("PluginConfig", &BotPluginConfig); err != nil {
		panic(err)
	}
	fmt.Println("Config loaded")
}
