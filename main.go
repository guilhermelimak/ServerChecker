package main

import (
	"svchk/checker"
	"svchk/config"
)

func main() {

	hosts := config.GetHostsList()

	checker.CheckSites(hosts.([]interface{}))
}
