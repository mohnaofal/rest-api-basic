package main

import (
	"encoding/json"
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

func findUserExist(form *User) bool {
	var exist bool
	for _, v := range users {
		if v.ID == form.ID || v.Email == form.Email {
			exist = true
			break
		}
	}

	return exist
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/user", Create).Methods("POST")
	r.HandleFunc("/users", ViewAll).Methods("GET")

	fmt.Println("Listening port 9090")
	http.ListenAndServe(":9090", r)
}

// Create new user
func Create(w http.ResponseWriter, r *http.Request) {
	form := new(User)

	json.NewDecoder(r.Body).Decode(form)

	if findUserExist(form) {
		w.WriteHeader(400)
		w.Write([]byte("User already exist"))
		return
	}

	users = append(users, *form)

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(form)
}

// ViewAll view all user
func ViewAll(w http.ResponseWriter, r *http.Request) {
	if len(users) == 0 {
		w.WriteHeader(404)
		w.Write([]byte("Data Not Found"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}
