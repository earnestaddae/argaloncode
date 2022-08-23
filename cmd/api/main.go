package main

import (
	"flag"
	"log"
	"os"
)

const version = "0.0.1"

// config contains the configurations
type config struct {
	port    int
	env     string
	version string
}

// application contains the application configurations
type application struct {
	config   config
	infoLog  *log.Logger
	errorLog *log.Logger
}

func main() {
	var cfg config
	// Declare flags for the port, environment and version
	flag.IntVar(&cfg.port, "port", 8080, "Server port to listen on")
	flag.StringVar(&cfg.env, "env", "development", "Application environment (development|staging|production")
	flag.StringVar(&cfg.version, "version", version, "Application version")
	flag.Parse()

	// Declare app logs
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// instantiate the application
	app := &application{
		config:   cfg,
		infoLog:  infoLog,
		errorLog: errorLog,
	}

	err := app.serve()
	if err != nil {
		log.Fatal(err)
	}
}
