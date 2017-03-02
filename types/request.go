package types

// Request : Request object containing the website, the http status code and if it's alive
type Request struct {
	IsAlive         bool
	Website, Status string
}
