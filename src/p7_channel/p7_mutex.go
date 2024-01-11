package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	x, x2  int64
	wg     sync.WaitGroup
	lock   sync.Mutex
	rwlock sync.RWMutex
)

func add1() {
	for i := 0; i < 5000; i++ {
		lock.Lock()
		x = x + 1
		lock.Unlock()
	}
	wg.Done()
}

func add2() {
	for i := 0; i < 5000; i++ {
		x2 = x2 + 1
	}
	wg.Done()
}

func write() {
	rwlock.Lock()
	x = x + 1
	time.Sleep(10 * time.Millisecond)
	rwlock.Unlock()
	wg.Done()
}

func read() {
	rwlock.RLock()
	time.Sleep(time.Millisecond)
	rwlock.RUnlock()
	wg.Done()
}

func main() {
	start := time.Now()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go write()
	}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go read()
	}

	wg.Wait()
	end := time.Now()
	fmt.Println(end.Sub(start))
	// wg.Add(2)

	// go add1()
	// go add1()

	// // go add2()
	// // go add2()
	// wg.Wait()

	// fmt.Println(x, x2)

}
