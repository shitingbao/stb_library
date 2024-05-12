package kafka

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	kafka "github.com/segmentio/kafka-go"
)

func kWriet(ctx context.Context, address, topic string) {
	kWrite := &kafka.Writer{
		Addr:     kafka.TCP(address),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}

	defer kWrite.Close()

	idx := 1

	for {
		k := fmt.Sprintf("%d", idx)
		v := uuid.New().String()
		log.Println(k, "-", v)
		err := kWrite.WriteMessages(ctx, kafka.Message{
			Key:   []byte(k),
			Value: []byte(v),
		})
		if err != nil {
			log.Println("write err:", err)
		}
		log.Println("wirte into~~~~~")
		time.Sleep(time.Second)
	}
}

func kRead(ctx context.Context, address, topic string) {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{address},
		GroupID:  "groupID",
		Topic:    topic,
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
	})

	defer reader.Close()

	for {
		log.Println("read start~")
		mes, err := reader.ReadMessage(ctx)
		log.Println("get start~")
		if err != nil {
			log.Println("read err:", err)
			return
		}

		log.Println("mes:", mes)
	}
}

func getTopics() {
	conn, err := kafka.Dial("tcp", "localhost:9092")
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close()

	partitions, err := conn.ReadPartitions()
	if err != nil {
		panic(err.Error())
	}

	m := map[string]struct{}{}

	for _, p := range partitions {
		m[p.Topic] = struct{}{}
	}
	for k := range m {
		fmt.Println(k)
	}
}

func createTopic(address, topic string) {
	controllerConn, err := kafka.Dial("tcp", address)

	if err != nil {
		panic(err.Error())
	}

	defer controllerConn.Close()

	topicConfigs := []kafka.TopicConfig{
		{
			Topic:             topic,
			NumPartitions:     1,
			ReplicationFactor: 1,
		},
	}

	err = controllerConn.CreateTopics(topicConfigs...)
	if err != nil {
		panic(err.Error())
	}
}
