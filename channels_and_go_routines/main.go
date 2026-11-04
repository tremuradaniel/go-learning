package main

import (
	"fmt"
	"net/http"
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
}
