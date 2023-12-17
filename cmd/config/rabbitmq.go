package config

import (
	"fmt"
	"github.com/streadway/amqp"
	"iot-messaging/cmd/model"
	"log"
)

const (
	TopicExchange = "topic"
)

func GetConn(connection model.ConnectionConfig) (*amqp.Connection, error) {
	connectionString := fmt.Sprintf("amqp://%s:%s@%s:%d/", connection.Username, connection.Password, connection.Host, connection.Port)
	conn, err := amqp.Dial(connectionString)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func CreateTopicExchange(conn *amqp.Connection, name string) (*amqp.Channel, error) {
	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}
	err = ch.ExchangeDeclare(
		name,
		TopicExchange,
		true,
		false,
		false,
		false,
		nil,
	)
	log.Printf("Exchange set up successfully")
	return ch, err
}

func SetupQueues(channel *amqp.Channel, queueConfig model.MessageConfig) error {
	for _, config := range queueConfig.Queues {
		queue, err := channel.QueueDeclare(
			config.Name,
			false,
			false,
			false,
			false,
			nil,
		)
		if err != nil {
			return err
		}

		for _, routingKey := range config.RoutingKey {
			err = channel.QueueBind(
				queue.Name,
				routingKey,
				queueConfig.Exchange,
				false,
				nil,
			)
			if err != nil {
				return err
			}
		}
	}
	log.Printf("Queues set up successfully")
	return nil
}
