package rabbitmq

import (
	"context"
	"log"
	"strconv"
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
		false,   // durable 是否持久化，需要生产者和消费者同时开启，并使用新的通道
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

// 无名称的临时队列
// 当方法返回时，队列实例包含一个由 RabbitMQ 生成的随机队列名称。例如，它可能看起来像amq.gen-JzTY20BRgKO-HjmUJj0wLg。
// 当声明它的连接关闭时，队列将被删除，因为它被声明为独占。

func (r *Rabbitmq) MqMesWithTemporary() {
	if err := r.Channel.ExchangeDeclare("logs", "fanout", true, false, false, false, nil); err != nil {
		return
	}
	q, err := r.Channel.QueueDeclare(
		"",    // name
		false, // durable 是否持久化，需要生产者和消费者同时开启，并使用新的通道
		false, // delete when unused
		true,  // exclusive , 因为他是独占，临时队列需要该值为 true，因为无名称需要唯一
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		log.Println(err)
		return
	}
	if err := r.Channel.QueueBind(
		q.Name, // queue name
		"",     // routing key
		"logs", // exchange
		false,
		nil,
	); err != nil {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	go func() {

		body := 1
		for {
			err = r.Channel.PublishWithContext(ctx,
				"",     // exchange
				q.Name, // routing key
				false,  // mandatory
				false,  // immediate
				amqp.Publishing{
					ContentType: "text/plain",
					Body:        []byte(strconv.Itoa(body)),
				})
			if err != nil {
				log.Println(err)
				return
			}
			log.Println("send mes:", body)
			body++
			time.Sleep(time.Second)
		}
	}()

	mes, err := r.Channel.Consume(q.Name, "", true, false, false, false, nil)
	if err != nil {
		log.Println()
		return
	}
	for {
		for m := range mes {
			log.Println("get mes:", string(m.Body))
		}
	}
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
		false,  // auto-ack false 关闭自动回复消息，并在处理完消息后使用 Ack 手动回收消息
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
			d.Ack(false)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

func (r *Rabbitmq) Close() {
	r.Connection.Close()
	r.Channel.Close()
}
