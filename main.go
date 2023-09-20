package main

import (
	"fmt"
	"github.com/eif-courses/golab/db"
	"github.com/eif-courses/golab/router"
	"github.com/eif-courses/golab/services"
	"log"
	"net/http"
	"os"
)

type Config struct {
	Port string
}
type Application struct {
	Config Config
	Models services.Models
}

func (app *Application) Serve() error {

	//err := godotenv.Load()
	//if err != nil {
	//	log.Fatal("Error loading .env file")
	//}
	port := os.Getenv("PORT")
	fmt.Println("API Is litenening on port", port)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: router.Routes(),
	}
	return server.ListenAndServe()
}

func main() {

	//err := godotenv.Load()
	//if err != nil {
	//	log.Fatal("Error loading .env file")
	//}
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
		Models: services.New(dbConnection.DB),
	}
	err = app.Serve()
	if err != nil {
		log.Fatal(err)
	}
}
