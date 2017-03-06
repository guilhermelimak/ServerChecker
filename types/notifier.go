package types

// NotifierData : A notifier object, used to communicate the user when a website is apparently down.
type NotifierData struct {
	Type        string
	Creds       interface{}
	Destination interface{}
}

// Emitter : Interface that implements the emit method
type Emitter interface {
	Emit()
}
