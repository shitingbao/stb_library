package spinlock

import (
	"log"
	"testing"
)

func TestLock(t *testing.T) {
	lock := NewSpinLock()
	lock.Lock()
	defer lock.Unlock()
}

func BenchmarkLock(b *testing.B) {
	lock := NewSpinLock()

	for i := 0; i < b.N; i++ {
		lock.Lock()
		log.Println("123")
		lock.Unlock()
	}

}
