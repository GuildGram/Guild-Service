package handlers

import (
	"encoding/json"
	"fmt"

	"github.com/GuildGram/Character-Service.git/data"
	"github.com/streadway/amqp"
)

func StartMsgBrokerConnection() {

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer ch.Close()

	msgs, err := ch.Consume(
		"TestQ",
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	forever := make(chan bool)
	go func() {
		for d := range msgs {
			var body data.Character
			err := json.Unmarshal(d.Body, &body)
			if err != nil {
				fmt.Println("failed to receive json msg")
			}

			fmt.Println(body)
			//do something with json data

			data.AddRosterInfo(body.UserID, &body)
			fmt.Println(data.GetGuild(1))
		}
	}()

	fmt.Println("Successfully connected to our RabbitMQ instance")
	fmt.Println(" [*] - waiting for messages")
	<-forever
}
