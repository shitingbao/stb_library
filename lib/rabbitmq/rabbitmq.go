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

// 带路由的交换器生产者
// 使用 direct 理论上就不用 queue 对象的 name 来作为 key
// 不然和消费者的 name 对应不上，除非 key 也用这个name
func (r *Rabbitmq) MqWithRoute() {
	if err := r.Channel.ExchangeDeclare("log_direct", "direct", true, false, false, false, nil); err != nil {
		return
	}
	ctx, canl := context.WithTimeout(context.Background(), time.Second*5)
	defer canl()
	num := 0
	for {
		r.Channel.PublishWithContext(
			ctx,
			"log_direct",
			"first",
			false,
			false,
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(strconv.Itoa(num)),
			},
		)
		log.Println("publish first:", num)
		num++
		r.Channel.PublishWithContext(
			ctx,
			"log_direct",
			"second",
			false,
			false,
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(strconv.Itoa(num)),
			},
		)
		log.Println("publish second:", num)
		num++
		time.Sleep(time.Second)
	}
}

// 带路由的交换器消费者
// 交换器已经在其他地方定义过就不用重新定义
// 注意，如果用队列，就不能使用同一个channel，会出现Exception (505) Reason: "UNEXPECTED_FRAME - expected content header for class 60, got non content header frame instead"
// 因为不同的线程复用了一个channel，这个指的是队列的channel
// 如果重新定义一个队列会无法接收到消息，因为 name 不对应，如果用队列，需要用相同的 queue 对象的名称来作为 key
func (r *Rabbitmq) MqWithRouteConsume() {
	// if err := r.Channel.ExchangeDeclare("log_direct", "direct", true, false, false, false, nil); err != nil {
	// 	return
	// }
	q, err := r.Channel.QueueDeclare("", true, false, false, false, nil)
	if err != nil {
		return
	}
	if err := r.Channel.QueueBind(q.Name, "first", "log_direct", false, nil); err != nil {
		return
	}
	mes, err := r.Channel.Consume(q.Name, "", true, false, false, false, nil)
	if err != nil {
		return
	}
	for m := range mes {
		log.Println("first:", string(m.Body))
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
