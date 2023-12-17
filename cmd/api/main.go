package main

import (
	"iot-messaging/cmd/config"
	"log"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	if err := run(); err != nil {
		log.Fatalf("Error running the application: " + err.Error())
	}
}

func run() error {
	conn, err := config.GetConn()
}
