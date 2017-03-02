package notify

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"svcheck/config"
	"svcheck/types"
)

// EmitPush : Send push to destination defined in config
func EmitPush(req types.Request) {
	list := config.GetNotifiersList().(map[string]interface{})
	creds := list["creds"].(map[string]interface{})
	accessToken := creds["access_token"]
	destination := list["destination"]

	url := req.Website

	notification := types.Notification{
		Type:  "note",
		Email: destination.(string),
		Body:  "Website health check failed: " + url,
	}

	aa := bytes.NewReader(prepareJSON(notification))

	sendPush(accessToken.(string), aa)
}

func prepareJSON(notification types.Notification) []byte {
	jsonObject, err := json.Marshal(notification)
	if err != nil {
		log.Fatal(err)
	}

	return jsonObject
}

func sendPush(accessToken string, payload io.Reader) {
	apiURL := "https://api.pushbullet.com/v2/pushes"

	client := &http.Client{}

	req, err := http.NewRequest("POST", apiURL, payload)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("Access-Token", accessToken)
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	// Print response body to stdout
	_, error := io.Copy(os.Stdout, res.Body)
	if err != nil {
		log.Fatal(error)
	}

	defer res.Body.Close()
}
