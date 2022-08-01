package lock

import (
	"sync/atomic"
)

type Locker uint32

func (sl *Locker) Lock() bool {
	return atomic.CompareAndSwapUint32((*uint32)(sl), 0, 1)
}

func (sl *Locker) Unlock() {
	atomic.StoreUint32((*uint32)(sl), 0)
}

func NewSLocker() *Locker {
	return new(Locker)
}
