package notifications

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"os"
	"svcheck/config"
	"svcheck/types"
	"svcheck/util"
)

// Pushbullet : A pushbullet notifier instance
type Pushbullet struct {
	types.NotifierData
}

// Emit : Send push to destination defined in config
func (notifier *Pushbullet) Emit(title, body string) {
	accessToken := notifier.Creds.(map[string]interface{})["access_token"]
	destination := notifier.Destination

	notification := types.Notification{
		Type:  "note",
		Email: destination.(string),
		Title: title,
		Body:  body,
	}

	JSONNotification := bytes.NewReader(util.ConvertNotificationToJSON(notification))

	sendPush(accessToken.(string), JSONNotification)
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

// NewPushbullet : Return a new Pushbullet instance
func NewPushbullet() Pushbullet {
	name := "pushbullet"

	notifierData := config.GetNotifierData(name).(map[string]interface{})

	notifier := types.NotifierData{
		Type:        name,
		Creds:       notifierData["creds"],
		Destination: notifierData["destination"],
	}

	pushbullet := Pushbullet{NotifierData: notifier}

	return pushbullet
}
