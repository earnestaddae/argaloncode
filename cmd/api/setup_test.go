package main

import (
	"log"
	"os"
	"testing"
)

var testApp application

func TestMain(m *testing.M) {
	testApp = application{
		config:   config{},
		infoLog:  log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime),
		errorLog: log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile),
	}

	os.Exit(m.Run())
}
