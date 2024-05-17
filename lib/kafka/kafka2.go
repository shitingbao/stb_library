package kafka

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

var topic = "hello1"

func Consumer() {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "127.0.0.1:9092",
		"group.id":          "myGroup",
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		panic(err)
	}

	if err := c.SubscribeTopics([]string{topic}, nil); err != nil {
		log.Println("SubscribeTopics:", err)
		return
	}

	for {
		msg, err := c.ReadMessage(time.Second * 2)
		if err != nil {
			log.Println("ReadMessage err:", err)
		} else {
			log.Println("k:v==:", string(msg.Key), string(msg.Value))

		}

	}
}

func producrer() {
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "127.0.0.1:9092"})
	if err != nil {
		panic(err)
	}

	defer p.Close()

	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("Delivered message to %v\n", ev.TopicPartition)
				}
			}
		}
	}()

	// Produce messages to topic (asynchronously)
	idx := 1
	for {
		err := p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          []byte(strconv.Itoa(idx)),
		}, nil)
		if err != nil {
			log.Println("Produce err:", err)
			return
		}
		log.Println("Produce:", idx)
		idx++

		time.Sleep(time.Second)
	}

	// Wait for message deliveries before shutting down
	p.Flush(15 * 1000)
}
