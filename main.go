package main

import (
	"svcheck/checker"
	"svcheck/config"
)

func main() {
	hosts := config.GetHostsList()

	checker.CheckSites(hosts.([]interface{}))
}
