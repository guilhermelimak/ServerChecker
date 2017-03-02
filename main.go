package main

import (
	"svcheck/checker"
	"svcheck/config"
)

func main() {
	config.GetHostsList()

	hosts := config.GetHostsList()
	checker.CheckSites(hosts.([]interface{}))
}
