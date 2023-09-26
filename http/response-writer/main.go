package main

import (
	"encoding/json"
	"net/http"
)

type User struct {
	Name   string
	Habits []string
}

func write(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Custom-Header", "custom")
	w.WriteHeader(201)
	user := &User{
		Name:   "jayleonc",
		Habits: []string{"balls", "running", "hiking"},
	}
	bytes, _ := json.Marshal(user)
	w.Write(bytes)
}

func main() {
	http.HandleFunc("/write", write)
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		return
	}
}
