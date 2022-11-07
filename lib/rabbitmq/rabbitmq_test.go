package rabbitmq

import (
	"log"
	"testing"
)

func TestRabbit(m *testing.T) {
	mq, err := NewMqCli("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Println(err)
		return
	}
	mq.MqMesWithTemporary()

}
