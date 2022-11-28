//Filename: qapi/cmd/api/main.go

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

// api version
const version = "1.0.0"

// config struct
type config struct {
	port int
	env  string //development, staging, production
}

// application struct
type application struct {
	config config
	logger *log.Logger
}

// main
func main() {
	//configuration flags
	var cfg config

	//Hard coding the server configurations
	cfg.port = 4000
	cfg.env = "development"

	//creating logger
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	//creating instance of app stuct
	app := &application{
		config: cfg,
		logger: logger,
	}

	//creating the HTTP Server
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.port),
		Handler: app.routes(),
	}
	//starting our server
	logger.Printf("starting %s server on %s", cfg.env, srv.Addr)
	err := srv.ListenAndServe()
	logger.Fatal(err)
}