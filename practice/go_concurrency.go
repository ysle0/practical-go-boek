package main

import (
	"fmt"
	"sync"
	"time"
)

func RunConcurrency() {
	// A goroutine is a light-weight thread managed by the GO runtime.
	// The evaluation of "world" happens in the current goroutine
	// and the execution of say happens in the new goroutine.
	//go say("world")
	//say("hello")

	// Goroutines run in the same address space, so access to shared memory
	// must be synchronised.
	//
	// The sync package provides useful primitives, although you won't need
	// then much in Go as there are other primitives.

	//s := []int{7, 2, 8, -9, 4, 0}

	// Channels are a typed conduit through which
	// you can send & receive values with the channel operator <-
	// like map and slice, channels must be created before use:
	//c := make(chan int)

	// by default, sends and receives block until the other side is ready.
	// This allows goroutines to synchronise without explicit locks or condition variables.
	//go sum(s[:len(s)/2], c)
	//go sum(s[len(s)/2:], c)
	//x, y := <-c, <-c

	//fmt.Println(x, y, x+y)

	// buffered channels
	// channels can be buffered. Provide the buffer length as the second argument
	// to make to initialize a buffered channel:
	//ch := make(chan int, 2)
	// sends to a buffered channel block only when the buffer is full.
	// receives block when the buffer is empty.

	//ch <- 1
	//ch <- 2
	//fmt.Println(<-ch)
	//fmt.Println(<-ch)

	// A sender can close a channel to indicate that no more values will be sent.
	// Receiver can test whether a channel has been closed by assigning a second parametre to the
	// receive express: after
	//v. ok := <- ch

	// ok is false if there are no more value to receive and the channel is closed.
	// * Only the sender should close a channel. never the receiver. Sending on a closed channel will cause a panic
	// * Channels are not like files: you do not need to close them. Channel is only necessary when the receiver
	//   must be told that there be no more values coming,
	//   such as to terminate a range loop.

	//c := make(chan int, 10)
	//go fibonacciChan(cap(c), c)
	//for i := range c {
	//	fmt.Println(i)
	//}
	//
	//tick := time.Tick(100 * time.Millisecond)
	//boom := time.After(500 * time.Millisecond)
	//for {
	//	select {
	//	case <-tick:
	//		fmt.Println("tick")
	//	case <-boom:
	//		fmt.Println("boom")
	//	default:
	//		fmt.Println("     ")
	//		time.Sleep(50 * time.Millisecond)
	//	}
	//}

	// what if we do not need communication? What if we just want to make sure
	// only one goroutine access a variable at a time to avoid data racing?
	// This concept is called mutual exclusion, and the conventional name
	// for the data structure that provides it is mutex.

	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Inc("k")
	}
	time.Sleep(time.Second)
	fmt.Println(c.Value("k"))
}

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

func fibonacciChan(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

type SafeCounter struct {
	mu sync.Mutex
	v  map[string]int
}

func (c *SafeCounter) Inc(key string) {
	//defer c.mu.Unlock()
	//c.mu.Lock()

	c.v[key]++
}

func (c *SafeCounter) Value(key string) int {
	//defer c.mu.Unlock()
	//c.mu.Lock()

	return c.v[key]
}

func (c *SafeCounter) Dec(key string) {
	//defer c.mu.Unlock()
	//c.mu.Lock()

	c.v[key]--
}
