package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Race Condition in Golang")
	wg := &sync.WaitGroup{}
	mutex := &sync.Mutex{}

	data := []int{0}

	wg.Add(3)
	go func(wg *sync.WaitGroup, mut *sync.Mutex) {
		mut.Lock()
		data = append(data, 1)
		mut.Unlock()
		fmt.Println("One")
		wg.Done()
	}(wg, mutex)
	go func(wg *sync.WaitGroup, mut *sync.Mutex) {
		mut.Lock()
		data = append(data, 2)
		mut.Unlock()
		fmt.Println("Two")
		wg.Done()
	}(wg, mutex)
	go func(wg *sync.WaitGroup, mut *sync.Mutex) {
		mut.Lock()
		data = append(data, 3)
		mut.Unlock()
		fmt.Println("Three")
		wg.Done()
	}(wg, mutex)
	wg.Wait()
	fmt.Println(data)
}

// Type go run --race .
