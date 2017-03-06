package notifications

// NotifyUser : Send a notification to user
func NotifyUser(title, body string) {
	notifier := NewPushbullet()

	notifier.Emit(title, body)
}
