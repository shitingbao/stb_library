package rabbitmq

import (
	"context"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Rabbitmq struct {
	Connection *amqp.Connection
	Channel    *amqp.Channel
}

// example amqp://guest:guest@localhost:5672/
func NewMqCli(url string) (*Rabbitmq, error) {
	con, err := amqp.Dial(url)
	if err != nil {
		return nil, err
	}
	cn, err := con.Channel()
	if err != nil {
		return nil, err
	}
	return &Rabbitmq{
		Connection: con,
		Channel:    cn,
	}, nil
}

func (r *Rabbitmq) SendMqMes() {
	q, err := r.Channel.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	if err != nil {
		log.Println(err)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	body := "Hello World!"
	err = r.Channel.PublishWithContext(ctx,
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	if err != nil {
		log.Println(err)
		return
	}
	log.Printf(" [x] Sent %s\n", body)
}

func (r *Rabbitmq) Receive() {
	q, err := r.Channel.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	if err != nil {
		log.Println(err)
		return
	}
	msgs, err := r.Channel.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		log.Println(err)
		return
	}

	var forever chan struct{}

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

func (r *Rabbitmq) Close() {
	r.Connection.Close()
	r.Channel.Close()
}
