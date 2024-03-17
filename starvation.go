package main

import (
	"fmt"
	"sync"
	"time"
)

type Result struct {
	Value int
	Name  string
}

func main() {
	var wg sync.WaitGroup
	var sharedLock sync.Mutex
	const runtime = 1 * time.Second
	var resultStream chan Result
	resultStream = make(chan Result)

	greedyWorker := func(result chan<- Result) {
		defer wg.Done()

		var count int
		for begin := time.Now(); time.Since(begin) <= runtime; {
			sharedLock.Lock()
			time.Sleep(3 * time.Nanosecond)
			sharedLock.Unlock()
			count++
		}
		rs := Result{
			Value: count,
			Name:  "greedyWorker",
		}
		result <- rs
		fmt.Printf("Greedy worker was able to execute %v work loops\n", count)
	}

	politeWorker := func(result chan<- Result) {
		defer wg.Done()

		var count int
		for begin := time.Now(); time.Since(begin) <= runtime; {
			sharedLock.Lock()
			time.Sleep(1 * time.Nanosecond)
			sharedLock.Unlock()

			sharedLock.Lock()
			time.Sleep(1 * time.Nanosecond)
			sharedLock.Unlock()

			sharedLock.Lock()
			time.Sleep(1 * time.Nanosecond)
			sharedLock.Unlock()

			count++
		}
		rs := Result{
			Value: count,
			Name:  "politeWorker",
		}
		result <- rs
		fmt.Printf("Polite worker was able to execute %v work loops.\n", count)
	}

	wg.Add(2)

	go greedyWorker(resultStream)
	go politeWorker(resultStream)

	go func() {
		wg.Wait()
		close(resultStream)
	}()
	// for value := range resultStream {
	// 	fmt.Println(value)
	// }
	for {
		select {
		case value, ok := <-resultStream:
			if !ok {
				fmt.Println("Received close signal")
				return
			}
			fmt.Println(value)
		}
	}
}
