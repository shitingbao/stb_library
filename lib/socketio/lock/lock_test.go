package lock

import (
	"log"
	"sync"
	"testing"
)

func TestLock(t *testing.T) {
	l := NewSLocker()
	log.Println(l.Lock())
	l.Unlock()
}

func BenchmarkLock(b *testing.B) {
	l := NewSLocker()

	wg := sync.WaitGroup{}
	for i := 0; i < b.N; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			l.Lock()
			l.Unlock()
		}()
	}
	wg.Wait()
}
