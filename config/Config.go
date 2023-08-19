package config

import (
	"log"

	"github.com/spf13/viper"
)

var App AppConfig

type AppConfig struct {
	MySqlAddr     string `mapstructure:"MYSQL_ADDRESS"`
	MySqlUsername string `mapstructure:"MYSQL_USERNAME"`
	MySqlPassword string `mapstructure:"MYSQL_PASSWORD"`
	MySqlDbName   string `mapstructure:"MYSQL_DBNAME"`
	WebAddress    string `mapstructure:"WEB_ADDRESS"`
}

func LoadConfig(path string) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app-config")
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Read config failed: ", err)
		return
	}
	err = viper.Unmarshal(&App)
	if err != nil {
		log.Fatal("Read config failed: ", err)
	}
	return
}
