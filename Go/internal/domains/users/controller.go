package users

import (
	"context"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gabszero/recommendation-system-go/internal/infra"
	amqp "github.com/rabbitmq/amqp091-go"
)

func WatchedMovie(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	userId, err := strconv.ParseUint(r.Form["userId"][0], 10, 32)
	if err != nil {
		panic(err)
	}

	movieId, err := strconv.ParseUint(r.Form["movieId"][0], 10, 32)
	if err != nil {
		panic(err)
	}

	infra.SaveMovieWatched(uint(userId), uint(movieId))

	conn, err := amqp.Dial("amqp://guest:guest@rabbitmq:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	body := "Hello World!"
	err = ch.PublishWithContext(ctx,
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	failOnError(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s\n", body)

}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}
