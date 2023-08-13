package main

import (
	"fmt"
	"net/http"
	"sync"
)

func main() {
	fmt.Println("Wait Groups in Golang")

	urls := []string{
		"https://www.google.com",
		"https://www.facebook.com",
		"https://seervitax.com",
		"https://www.google.com",
		"https://www.facebook.com",
		"https://seervitax.com",
		"https://www.google.com",
		"https://www.facebook.com",
		"https://seervitax.com",
		"https://www.google.com",
		"https://www.facebook.com",
		"https://seervitax.com",
		"https://www.google.com",
		"https://www.facebook.com",
		"https://seervitax.com",

	}
	var waitgrp sync.WaitGroup
	
	for _, web := range urls {
		waitgrp.Add(1)
		go getStatus(&waitgrp,web)
	}
	waitgrp.Wait()
}

func getStatus(wg *sync.WaitGroup, endpoint string) {
	defer wg.Done()
	res, err := http.Get(endpoint)
	if err != nil {
		panic(err)
	}
	fmt.Println(res.StatusCode)
}
