package models

const (
	CONNECTION_TIMEOUT = 5
)

type APIResponse struct {
	Data interface{}
	Err  error
}
