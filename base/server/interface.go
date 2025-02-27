package server

// ServerInterface is transport server.
type ServerInterface interface {
	Start() error
	Shutdown() error
}
