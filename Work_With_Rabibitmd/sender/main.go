package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/streadway/amqp"
	"log"
)

//
//var ErrorConnectRabbitMQ = errors.New("Ошибка подключения к RabbitMQ")
//var ErrorCHannelRabbitMQ = errors.New("Ошибка канала RabbitMQ")

func main() {
	amqpServerURL := "amqp://guest:guest@10.10.0.151:5672/" //получаем параметр из env

	//создание подключения
	connectRabbitMQ, err := amqp.Dial(amqpServerURL)
	if err != nil {
		//return panic(err)
		panic(err)
	}
	defer connectRabbitMQ.Close()

	channelRabbitMQ, err := connectRabbitMQ.Channel()
	if err != nil {
		panic(err)
	}
	defer channelRabbitMQ.Close()

	_, err = channelRabbitMQ.QueueDeclare(
		"QueueService1", // queue name
		true,            // durable
		false,           // auto delete
		false,           // exclusive
		false,           // no wait
		nil,             // arguments
	)

	if err != nil {
		panic(err)
	}

	//создание фибера
	app := fiber.New()

	app.Use(
		logger.New(), //логирование
	)

	app.Get("/send", func(ctx *fiber.Ctx) error {
		message := amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(ctx.Query("msg")),
		}
		if err := channelRabbitMQ.Publish(
			"",              // exchange
			"QueueService1", // queue name
			false,           // mandatory
			false,           // immediate
			message,         // message to publish
		); err != nil {
			return err
		}
		return nil
	})
	log.Fatal(app.Listen(":3000"))
}
