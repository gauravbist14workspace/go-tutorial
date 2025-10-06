package server

import (
	"fmt"
	"net/http"
)

func NewServer() *http.ServeMux {

	handler := http.NewServeMux()
	handler.HandleFunc("GET /hello/{name}", GreetNormal)

	return handler
}

func GreetNormal(w http.ResponseWriter, req *http.Request) {
	name := req.PathValue("name")
	w.Write([]byte(fmt.Sprintf("Hello %v", name)))
}
