package main

import (
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	// for i := 0; i < len(links); i++ { // would only loop 5 times, the number of links

	// for { // infinite loop
	// 	// fmt.Println(<-c)
	// 	go checkLink(<-c, c) // arguments are : the value of the channel and the channel itself
	// }

	for l := range c { // clearer syntax for infinite loop
		// go checkLink(l, c) // same arguments as above

		//adding function literal (lambda in python) to fix the bug
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l) // passing l since it's in the outer scope
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)

	if err != nil {
		println(link, "might be down!")
		// c <- "Might be down I think"
		c <- link
		return
	}

	println(link, "is up!")
	// c <- "Yep its up"
	c <- link
}
