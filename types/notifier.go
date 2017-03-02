package types

// Notifier : A notifier object, used to communicate the user when a website is apparently down.
type Notifier struct {
	Type  string
	Creds interface{}
	Emitter
}

// Emitter : Interface that implements the emit method
type Emitter interface {
	emit(Request)
}
