package config

import (
	"log"

	"github.com/spf13/viper"
)

func loadConfig() {
	fileName := "config"

	viper.SetConfigName(fileName)
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()

	if err != nil {
		log.Fatal(err)
	}
}

// GetHostsList : Parse the hosts-list file and return the hosts list
func GetHostsList() interface{} {
	loadConfig()
	return viper.Get("hosts")
}

// GetNotifiersList : Get the notification types to
func GetNotifiersList() interface{} {
	loadConfig()
	return viper.Get("notifiers")
}
