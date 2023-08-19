package config

import (
	"log"

	"github.com/spf13/viper"
)

var App AppConfig

type AppConfig struct {
	// MySQL connnection address in format {host}:{port}
	MySqlAddr string `mapstructure:"MYSQL_ADDRESS"`
	// MySQL username
	MySqlUsername string `mapstructure:"MYSQL_USERNAME"`
	// MySQL password
	MySqlPassword string `mapstructure:"MYSQL_PASSWORD"`
	// MySQL database name
	MySqlDbName string `mapstructure:"MYSQL_DBNAME"`
	// Web connection address in format {host}:{port}
	WebAddress string `mapstructure:"WEB_ADDRESS"`
}

// Load program configuration from the local file into the global configuration struct
func LoadConfig(path string) {
	// set the config file path
	viper.AddConfigPath(path)
	// set the config file name
	viper.SetConfigName("app-config")
	// set the config file type
	viper.SetConfigType("env")
	// load environment variable if available
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
