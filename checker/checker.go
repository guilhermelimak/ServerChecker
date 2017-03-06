package checker

import (
	"fmt"
	"log"
	"net/http"
	"svcheck/notifications"
	"svcheck/types"
)

// CheckSites : check website and return status
func CheckSites(url []interface{}) bool {
	allAlive := true

	channel := make(chan types.Request)

	for _, s := range url {
		go isSiteAlive(s.(string), channel)

		req := <-channel
		log.Printf("%s", req.Website)

		if req.IsAlive == false {
			title := fmt.Sprintf("Website is down")
			body := fmt.Sprintf("Website %s is down. Error: %s", req.Website, req.Status)

			notifications.NotifyUser(title, body)
		}

	}

	return allAlive
}

func isSiteAlive(url string, c chan types.Request) {
	res, err := http.Get(url)

	var status string

	isAlive := true

	if err != nil {
		isAlive = false
		status = err.Error()
	} else {
		status = res.Status
	}

	c <- types.Request{
		Website: url,
		Status:  status,
		IsAlive: isAlive,
	}
}
