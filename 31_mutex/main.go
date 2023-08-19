package main

import (
	"fmt"
	"net/http"
	"sync"
)

var test = []string{"test"}
var waitgrp sync.WaitGroup
var mut sync.Mutex

func main() {
	fmt.Println("Mutex Groups in Golang")

	urls := []string{
		"https://www.google.com",
		"https://www.facebook.com",
		"https://seervitax.com",
		"https://www.google.com",
	}

	for _, web := range urls {
		waitgrp.Add(1)
		go getStatus(&waitgrp, web)
	}
	waitgrp.Wait()
}

func getStatus(wg *sync.WaitGroup, endpoint string) {
	defer wg.Done()
	res, err := http.Get(endpoint)
	if err != nil {
		panic(err)
	} else {
		mut.Lock()
		test = append(test, endpoint)
		mut.Unlock()
		fmt.Println(res.StatusCode,test)

	}

}
