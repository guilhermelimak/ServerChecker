package config

import (
	"log"

	"github.com/spf13/viper"
)

func loadConfig() {
	fileName := ".srvchkrc"

	viper.SetConfigName(fileName)

	viper.AddConfigPath(".")
	viper.AddConfigPath("$HOME")

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

// GetNotifierData : Get the notification types to
func GetNotifierData(name string) interface{} {
	loadConfig()

	notifiers := viper.Get("notifiers")
	notifier := notifiers.(map[string]interface{})[name]

	return notifier
}
