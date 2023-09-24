package main

import (
	"fmt"
	"github.com/eif-courses/golab/api"
	"github.com/eif-courses/golab/services"
	"github.com/joho/godotenv"
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

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("PORT")
	fmt.Println("API Is litenening on port", port)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: api.Routes(),
	}
	return server.ListenAndServe()
}

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	config := Config{
		Port: os.Getenv("PORT"),
	}

	dsn := os.Getenv("DSN")
	dbConnection, err := api.ConnectPostgres(dsn)
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
