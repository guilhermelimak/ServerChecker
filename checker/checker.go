package checker

import (
	"net/http"
	"runtime"
	"svcheck/notify"
	"svcheck/types"
)

// CheckSites : check website and return status
func CheckSites(url []interface{}) bool {
	runtime.GOMAXPROCS(10)
	allAlive := true

	channel := make(chan types.Request)

	for _, s := range url {
		go isSiteAlive(s.(string), channel)

		req := <-channel

		if req.IsAlive {
			notify.SendNotification(req)
		} else {
			allAlive = false
			notify.SendNotification(req)
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
