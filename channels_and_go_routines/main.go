package main

import (
	"fmt"
	"net/http"
	"sync"
)

func main() {
	// links := []string{
	// 	"http://google.com",
	// 	"http://facebook.com",
	// 	"http://stackoverflow.com",
	// 	"http://golang.com",
	// 	"http://amazon.com",
	// 	"http://dexonline.ro",
	// }

	// c := make(chan string)

	// for _, link := range links {
	// 	// go = run this function inside a new routine
	// 	go checkLink(link, c)
	// }

	// for l := range c {
	// 	go func(l string) {
	// 		time.Sleep(5 * time.Second)
	// 		checkLink(l, c)
	// 	}(l)
	// }

	testAnonymousFunctions()

	testSyncWaitGroup()
}

func testSyncWaitGroup() {
	fmt.Println("in testSyncWaitGroup")

	// a WaitGroup is just a counter with a “wait until it hits zero” feature
	var wg sync.WaitGroup

	go func() {
		// does not get printed
		fmt.Println("in go routine")
	}()

	// Counter goes to 1.
	// You’re saying: “I’m about to start 1 goroutine, and I need to wait for it.”
	wg.Add(1)

	go func() {
		defer wg.Done() // subtracts 1
		// gets printed
		fmt.Println("in go routine 2")
	}()

	wg.Wait() // waits until counter is 0
}

func checkLink(link string, c chan (string)) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, "might be donw!")
		c <- link
		return
	}
	fmt.Println(link, "is up!")
	c <- link
}

func testAnonymousFunctions() {

	fmt.Println("in testAnonymousFunctions...")

	// go routine anonymous function
	go func() {
		// does not get printed
		// hasn't got the time to finish
		fmt.Println("running from goroutine")
	}()

	// anonymous function
	func() {
		// gets printed
		fmt.Println("running from anonymous function!")
	}()

	// if uncommented, gives enough time for go routine anonymous function
	// to run
	// time.Sleep(time.Second)
}
