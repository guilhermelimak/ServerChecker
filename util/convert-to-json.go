package util

import (
	"encoding/json"
	"log"
)

// ConvertNotificationToJSON : Get json objet from a notificiation instance
func ConvertNotificationToJSON(notification interface{}) []byte {
	jsonObject, err := json.Marshal(notification)
	if err != nil {
		log.Fatal(err)
	}

	return jsonObject
}
