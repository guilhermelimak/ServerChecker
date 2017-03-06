package types

// Notification : Pushbullet notification
type Notification struct {
	Body  string `json:"body"`
	Email string `json:"email"`
	Type  string `json:"type"`
	Title string `json:"title"`
}
