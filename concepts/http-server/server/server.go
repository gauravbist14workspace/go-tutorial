package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	models "go_tutorial/concepts/http-server/models"
)

func NewServer() *http.ServeMux {

	handler := http.NewServeMux()
	handler.HandleFunc("GET /hello/{name}", GreetNormal)

	return handler
}

func GreetNormal(w http.ResponseWriter, req *http.Request) {
	name := req.PathValue("name")

	user := models.User{}
	if err := json.NewDecoder(req.Body).Decode(&user); err != nil {
		fmt.Printf("err occured while decoding the request body, err = %v", err)
		return
	}

	w.Write([]byte(fmt.Sprintf("Hello %v who is %v years old", name, user.Age)))
}
