package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Creating a cache, each book has an ID.
var cache = make(map[int]Book)
var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

func RunConcurrency2() {
	wg := sync.WaitGroup{}
	rwm := sync.RWMutex{}

	cacheCh := make(chan Book)
	dbCh := make(chan Book)

	for i := 0; i < 10; i++ {
		id := rnd.Intn(10) + 1
		wg.Add(2) // we wait for 2 goroutines

		go queryCacheGo(id, &wg, &rwm, cacheCh)
		go queryDbGo(id, &wg, &rwm, dbCh)

		go func(cacheCh, dbCh <-chan Book) {
			select {
			case b := <-cacheCh:
				fmt.Println("source [cache]")
				fmt.Println(b)
				<-dbCh
			case b := <-dbCh:
				fmt.Println("source [db]")
				fmt.Println(b)
			}
		}(cacheCh, dbCh)

		time.Sleep(150 * time.Millisecond)
	}
	wg.Wait()
}

func queryCacheGo(
	id int,
	wg *sync.WaitGroup,
	mx *sync.RWMutex,
	bookCh chan<- Book,
) {
	if b, ok := queryCache(id, mx); ok {
		bookCh <- b
	}
	wg.Done()
}

func queryDbGo(id int, wg *sync.WaitGroup, mx *sync.RWMutex, bookCh chan<- Book) {
	if b, ok := queryDb(id, mx); ok {
		bookCh <- b
	}
	wg.Done()
}

func queryCache(id int, mx *sync.RWMutex) (Book, bool) {
	mx.RLock()
	defer mx.RUnlock()
	b, ok := cache[id]
	return b, ok
}

func mockDbQueryIoLatency(ms time.Duration) {
	time.Sleep(ms * time.Millisecond)
}

func queryDb(id int, rwm *sync.RWMutex) (Book, bool) {
	mockDbQueryIoLatency(100)

	for _, b := range books {
		if b.ID == id {
			rwm.Lock()
			defer rwm.Unlock()
			cache[id] = b
			return b, true
		}
	}

	return Book{}, false
}
