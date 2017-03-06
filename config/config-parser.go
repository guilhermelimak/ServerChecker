package config

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

func createDefaultConfig() {
	exampleFile := "./.svchk.example.yml"
	configName := ".svchk.yml"
	homeDir, homeDirErr := homedir.Dir()

	if homeDirErr != nil {
		log.Fatal(homeDirErr)
	}

	destination := fmt.Sprintf("%s/%s", homeDir, configName)
	cpCmd := exec.Command("cp", exampleFile, destination)

	var out bytes.Buffer
	var stderr bytes.Buffer

	cpCmd.Stdout = &out
	cpCmd.Stderr = &stderr

	cpErr := cpCmd.Run()
	if cpErr != nil {
		fmt.Println(fmt.Sprint(cpErr) + ": " + stderr.String())
		return
	}
}

func loadConfig() {
	fileName := ".svchk"

	viper.SetConfigName(fileName)

	viper.AddConfigPath(".")
	viper.AddConfigPath("$HOME")

	configErr := viper.ReadInConfig()
	if configErr != nil {
		createDefaultConfig()
		log.Fatal(configErr)
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
