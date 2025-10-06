package main

import "fmt"

// FunctionalOptions is a design pattern that allows for flexible configuration of functions or methods.
type FunctionalOptions func(*serverOptions)

type serverOptions struct {
	tls                 bool
	id                  string
	numberOfConnections int
}

type server struct {
	serverOptions
}

func defaultOptions() serverOptions {
	return serverOptions{
		tls:                 false,
		id:                  "default",
		numberOfConnections: 10,
	}
}

func NewServer(opts ...FunctionalOptions) *server {

	o := defaultOptions()

	for _, fn := range opts {
		fn(&o)
	}

	return &server{
		serverOptions: o,
	}
}

func WithTLS(o *serverOptions) {
	o.tls = true
}

func SetNumberOfConnections(count int) FunctionalOptions {
	return func(o *serverOptions) {
		o.numberOfConnections = count
	}
}

func SetID(id string) FunctionalOptions {
	return func(o *serverOptions) {
		o.id = id
	}
}

func main() {
	server := NewServer(WithTLS, SetNumberOfConnections(19))
	if server != nil {
		fmt.Printf("%+v\n", server)
	}
}
