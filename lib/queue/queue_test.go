package queue

import (
	"log"
	"testing"
)

func TestQueue(t *testing.T) {
	q := Constructor(3)
	q.EnQueue(1)
	q.EnQueue(2)
	q.EnQueue(3)
	log.Println(q.Front())
	q.DeQueue()
	log.Println(q.Rear())
}
