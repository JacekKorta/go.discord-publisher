package main

import (
	"bytes"
	"discord-publisher/msgs"
	"discord-publisher/settings"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

var wg = sync.WaitGroup{}

func main() {

	settings := &settings.Settings{}
	settings.GetSettings()
	amqpAddress := settings.GetRabbitmqUrl()

	httpClient := &http.Client{}

	conn, err := amqp.Dial(amqpAddress)
	log.Println(amqpAddress)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	chConsume, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	queue, err := ch.QueueDeclarePassive(
		settings.Rabbit.InputQueue, // name
		true,                       // durable
		false,                      // delete when unused
		false,                      // exclusive
		false,                      // no-wait
		nil,                        // arguments
	)
	failOnError(err, "Failed to declare a queue")

	messages, err := chConsume.Consume(
		queue.Name, // queue
		"Tagger",   // consumer
		true,       // auto-ack
		false,      // exclusive
		false,      // no-local
		false,      // no-wait
		nil,        // args
	)
	failOnError(err, "Failed to register a consumer")

	var forever chan struct{}

	go func() {
		for d := range messages {
			var msg msgs.DiscordMessageOut
			var emb msgs.Embed
			q := &msgs.QuestionIn{}
			json.Unmarshal(d.Body, q)
			reasons := strings.Join(q.Reasons, " ")
			currentTime := time.Now()
			msg.Content = currentTime.Format("2006.01.02 15:04:05")
			emb.Title = q.Title
			emb.Description = fmt.Sprintf("ID: %v\nReasons: %v\nLink: %v", q.QuestionID, reasons, q.Link)
			msg.Embeds = append(msg.Embeds, emb)
			jsonBody, _ := json.Marshal(msg)
			req, err := http.NewRequest("POST", settings.DiscordWebhook, bytes.NewReader(jsonBody))
			failOnError(err, "Unable to prepare request")
			req.Header.Add("Content-Type", "application/json")
			resp, reqErr := httpClient.Do(req)
			failOnError(reqErr, "Sending discord message fail")
			log.Println("message publish status code: ", resp.StatusCode)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever

	wg.Wait()
}
