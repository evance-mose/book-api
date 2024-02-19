package controllers

import (
	"apiv1/data"
	"apiv1/datatypes"
	"encoding/json"
	"fmt"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	// ...
	var user datatypes.UserCredentials
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	_, ok := data.Users[user.Username]

	if !ok {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User not found"))
		return
	}

	if data.Users[user.Username] == user.Password {
		w.WriteHeader(http.StatusFound)
		w.Write([]byte("User found"))
		return
	}

	w.WriteHeader(http.StatusUnauthorized)
	w.Write([]byte("Invalid credentials"))
}

func Register(w http.ResponseWriter, r *http.Request) {
	// ...

	var newUser datatypes.UserCredentials

	err := json.NewDecoder(r.Body).Decode(&newUser)

	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	//check the user exists
	_, ok := data.Users[newUser.Username]

	if ok {
		w.WriteHeader(http.StatusConflict)
		w.Write([]byte("User already exists"))
		return
	}

	data.Users[newUser.Username] = newUser.Password
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("User created"))
}
