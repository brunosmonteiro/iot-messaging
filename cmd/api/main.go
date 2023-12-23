package main

import (
	"iot-messaging/cmd/api/config"
	"log"
)

func main() {
	if err := run(); err != nil {
		log.Fatalf("Error running the application: " + err.Error())
	}
}

func run() error {
	messageConfig, err := config.LoadConfig()
	if err != nil {
		return nil
	}
	conn, err := config.GetConn(messageConfig.Connection)
	if err != nil {
		return err
	}
	channel, err := config.CreateTopicExchange(conn, messageConfig.Exchange)
	if err != nil {
		return err
	}
	err = config.SetupQueues(channel, messageConfig)
	if err != nil {
		return err
	}
	return nil
}
