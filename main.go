package main

import (
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"golab/db"
	"log"
	"net/http"
	"os"
)

type Config struct {
	Port string
}
type Application struct {
	Config Config
}

func (app *Application) Serve() error {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("PORT")
	fmt.Println("API Is litenening on port", port)

	server := &http.Server{
		Addr: fmt.Sprintf(":%s", port),
		// TODO Add router
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

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	config := Config{
		Port: os.Getenv("PORT"),
	}

	dsn := os.Getenv("DSN")
	dbConnection, err := db.ConnectPostgres(dsn)
	if err != nil {
		log.Fatal("Cannot connect to database!")
	}

	defer dbConnection.DB.Close()

	app := &Application{
		Config: config,
		// todo add models
	}
	err = app.Serve()
	if err != nil {
		log.Fatal(err)
	}

	//
	//http.HandleFunc("/animals", AnimalsHandler)
	//http.HandleFunc("/status", HealthHandler)
	//log.Println("** Service Started on Port 8080 **")
	//err := http.ListenAndServe(":8080", nil)
	//if err != nil {
	//	log.Fatal(err)
	//}
}
