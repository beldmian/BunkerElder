package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	TelegramAPIToken string `mapstructure:"telegram_api_token"`	
}

func ProvideConfig() *Config {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}
	conf := Config{}
	log.Println(viper.Get("telegram_api_token"))
	if err := viper.Unmarshal(&conf); err != nil {
		log.Fatal(err)
	}
	return &conf
}
