package V2

import (
	"sync"
)

type counter struct {
	key   string
	value int
	ch    chan func()
	wg    sync.WaitGroup
}

//type counters struct {
//	table []*counter
//}

//初始化计数器
func Init(ctKey ...string) *counter {
	ct := &counter{
		key:   "",
		value: 0,
		ch:    make(chan func()),
	}
	if ctKey != nil {
		ct.key = ctKey[0]
	}

	go start(ct.ch)
	return ct
}

func start(ch chan func()) {
	for f := range ch {
		f()
	}

}

//func (ct *counter) Flush2broker(t time.Duration, FuncCbFlush FlushCb) {
//	ticker := time.NewTicker(t * time.Millisecond)
//	go func() {
//		for range ticker.C {
//			FuncCbFlush()
//		}
//	}()
//}

func (ct *counter) Incr(key string, value int) {
	ct.wg.Add(1)
	ct.ch <- func() {
		defer func() {
			ct.wg.Done()
		}()
		if ct.key != "" {
			ct.value += value
		} else {
			ct.key = key
			ct.value = value

		}
	}

}

func (ct *counter) Get() int {
	n := make(chan int)
	ct.wg.Add(1)
	ct.ch <- func() {
		n <- ct.value
		close(n)
		defer ct.wg.Done()
	}
	return <-n

}

//type FlushCb func()

//储存当前计数值并reset
//func (ct *counter) flush() {
//	ct.wg.Add(1)
//	ct.ch
//	cts := &counters{}
//	cts.table = append(cts.table, ct)
//	ct.value = 0
//
//}
