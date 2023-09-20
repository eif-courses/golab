package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Config struct {
	Port string
}
type Application struct {
	Config Config
}

var port = "8080"

func (app *Application) Serve() error {

	server := &http.Server{
		Addr: fmt.Sprintf(":%s", port),
	}
	return server.ListenAndServe()
}

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

	var config Config
	config.Port = port

	http.HandleFunc("/animals", AnimalsHandler)
	http.HandleFunc("/status", HealthHandler)
	log.Println("** Service Started on Port 8080 **")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
