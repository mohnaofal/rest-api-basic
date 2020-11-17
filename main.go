package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// User struct model
type User struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Gender      string `json:"gender"`
}

var users []User

func main() {
	r := mux.NewRouter()

	fmt.Println("Listening port 9090")
	http.ListenAndServe(":9090", r)
}
