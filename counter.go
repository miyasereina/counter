package counter

import (
	"sync"
	"time"
)

type counter struct {
	key   string
	value int
	mutex sync.RWMutex
}
type counters struct {
	table []*counter
}

//初始化计数器
func Init() *counter {
	return &counter{}
}

func (ct *counter) Flush2broker(t time.Duration, FuncCbFlush FlushCb) {
	ticker := time.NewTicker(t * time.Millisecond)
	go func() {
		for range ticker.C {
			FuncCbFlush()
		}
	}()
}

func (ct *counter) Incr(key string, value int) {
	ct.mutex.Lock()
	if ct.key != "" {
		ct.value += value

	} else {
		ct.key = key
		ct.value = value
	}
	defer ct.mutex.Unlock()

}

func (ct *counter) Get() int {
	ct.mutex.RLock()
	defer ct.mutex.RUnlock()
	return ct.value

}

type FlushCb func()

//储存当前计数值并reset
func (ct *counter) flush() {
	cts := &counters{}
	ct.mutex.Lock()
	defer ct.mutex.Unlock()
	cts.table = append(cts.table, ct)
	ct.value = 0

}
