package main

import (
	"encoding/json"
	"net/http"
)

type Animal struct {
	Name string `json:"Name"`
	Type string `json:"Type"`
}

func AnimalsHandler(w http.ResponseWriter, r *http.Request) {
	animals := []Animal{
		{"Alice", "Cat"},
		{"Bob", "Cat"},
		{"Trinity", "Dog"},
	}
	err := json.NewEncoder(w).Encode(animals)
	if err != nil {
		return
	}
}

func HealthHandler(w http.ResponseWriter, r *http.Request) {

	data := Animal{"Liutas", "Kate"}
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		return
	}
}

func main() {
	http.HandleFunc("/animals", AnimalsHandler)
	http.HandleFunc("/status", HealthHandler)
	err := http.ListenAndServe(":1234", nil)
	if err != nil {
		return
	}
}
