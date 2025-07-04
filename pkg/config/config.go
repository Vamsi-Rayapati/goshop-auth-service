package config

import (
	"github.com/spf13/viper"
)

type ApplicationConfig struct {
	Port       int
	DbHost     string
	DbPort     int
	DbUserName string
	DbPassword string
	FaClientId string
	FaSecret   string
	FaApiKey   string
	FaAppId    string
	FaUrl      string
}

var Config *ApplicationConfig

func LoadConfig() *ApplicationConfig {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	viper.ReadInConfig()

	Config = &ApplicationConfig{
		Port:       viper.GetInt("PORT"),
		DbHost:     viper.GetString("DB_HOST"),
		DbPort:     viper.GetInt("DB_PORT"),
		DbUserName: viper.GetString("DB_USER_NAME"),
		DbPassword: viper.GetString("DB_PASSWORD"),
		FaClientId: viper.GetString("FA_CLIENT_ID"),
		FaSecret:   viper.GetString("FA_SECRET"),
		FaApiKey:   viper.GetString("FA_API_KEY"),
		FaAppId:    viper.GetString("FA_APP_ID"),
		FaUrl:      viper.GetString("FA_URL"),
	}

	return Config

}
