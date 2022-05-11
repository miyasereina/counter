package V2

import (
	"counter/myLog"
	"encoding/json"

	"log"

	"sync"
	"time"
)

type counter struct {
	created_time time.Time
	key          string
	value        int
	ch           chan func()
	wg           sync.WaitGroup
}

type counters struct {
	table map[string]*counter
}

var Cts = InitCts()

func InitCts() *counters {
	return &counters{
		table: make(map[string]*counter),
	}
}

//初始化计数器
func Init(ctKey string) *counter {
	ct, ok := Cts.table[ctKey]
	if ok {
		return ct
	}
	ct = &counter{
		created_time: time.Now(),
		key:          ctKey,
		value:        0,
		ch:           make(chan func()),
	}
	//加入计数器管理表
	Cts.table[ct.key] = ct
	go start(ct.ch)
	return ct
}

func start(ch chan func()) {
	for f := range ch {
		f()
	}

}

func (ct *counter) Incr(value int) {
	ct.wg.Add(1)
	ct.ch <- func() {
		defer func() {
			ct.wg.Done()
		}()
		ct.value = ct.value + value
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
func (ct *counter) Reset() {
	ct.wg.Add(1)
	ct.ch <- func() {
		go flush(ct.value, ct.key, ct.created_time)
		ct.value = 0
	}
}

func (cts *counters) IncrWithIndex(key string, value int) {
	ct, ok := cts.table[key]
	if !ok {
		log.Println("there is no +" + key + " in this counters")
		return
	}
	ct.Incr(value)

}

func (cts *counters) GetWithIndex(key string) int {
	ct, ok := cts.table[key]
	if !ok {
		log.Println("there is no +" + key + " in this counters")
		return 0
	}
	return ct.Get()
}

func (cts *counters) ResetIndex(key string) {
	ct, ok := cts.table[key]
	if !ok {
		log.Println("there is no " + key + " in this counters")
		return
	}
	ct.Reset()

}

func (cts *counters) ResetAll() {

}

type record struct {
	Name  string    `json:"name"`
	T     time.Time `json:"t"`
	Value int       `json:"value"`
}

func flush(value int, name string, start time.Time) {
	info := record{
		Name:  name,
		Value: value,
		T:     start,
	}
	s, err := json.Marshal(info)
	if err != nil {
		panic(err)
	}
	myLog.Logfile(string(s[1 : len(s)-1]))
}
