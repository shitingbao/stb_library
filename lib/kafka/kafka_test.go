package kafka

import (
	"context"
	"testing"
)

func TestMain(m *testing.M) {
	// getTopics()
	ctx := context.Background()

	address := "localhost:9092"
	topic := "hello"

	go kRead(ctx, address, topic)

	kWriet(ctx, address, topic)
}
