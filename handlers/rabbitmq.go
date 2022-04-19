package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/GuildGram/Character-Service.git/data"
	"github.com/streadway/amqp"
)

//method for repeated code
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func StartMsgBrokerConnection(n int) (res *data.Character, err error) {

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"guild_rpc", // name
		false,       // durable
		false,       // delete when unused
		false,       // exclusive
		false,       // no-wait
		nil,         // arguments
	)
	failOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	corrId := "getall"
	err = ch.Publish(
		"",          // exchange
		"guild_rpc", // routing key
		false,       // mandatory
		false,       // immediate
		amqp.Publishing{
			ContentType:   "text/plain",
			CorrelationId: corrId,
			ReplyTo:       q.Name,
			Body:          []byte(strconv.Itoa(n)),
		})
	failOnError(err, "Failed to publish a message")

	for d := range msgs {
		if corrId == d.CorrelationId {

			var res data.Character
			err = json.Unmarshal(d.Body, &res)
			data.AddRosterInfo(&res)
			failOnError(err, "Failed to convert body to integer")
			fmt.Print(res)
			break
		}
	}

	return
}
