package main

import (
	"github.com/streadway/amqp"
	"log"
)

func main() {
	amqpServerURL := "amqp://guest:guest@10.10.0.151:5672/"

	// Create a new RabbitMQ connection.
	connectRabbitMQ, err := amqp.Dial(amqpServerURL)
	if err != nil {
		panic(err)
	}
	defer connectRabbitMQ.Close()

	// Opening a channel to our RabbitMQ instance over
	// the connection we have already established.
	channelRabbitMQ, err := connectRabbitMQ.Channel()
	if err != nil {
		panic(err)
	}
	defer channelRabbitMQ.Close()

	messages, err := channelRabbitMQ.Consume(
		"QueueService1", // queue name
		"",              // consumer
		true,            // auto-ack
		false,           // exclusive
		false,           // no local
		false,           // no wait
		nil,             // arguments
	)
	if err != nil {
		log.Println(err)
	}

	log.Println("Подключение успешно")
	log.Println("Ожидание сообщений")

	forever := make(chan bool)

	go func() {
		for message := range messages {
			log.Println("> Пересылка сообщения: ", string(message.Body))
		}
	}()
	<-forever
}
