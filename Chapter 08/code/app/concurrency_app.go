package main

import (
	"fmt"
	"net/http"
)


func main() {

	channel := make(chan string)

	links := []string{
		"https://www.golangbot.com",
		"https://www.golang.org",
		"https://www.changelog.com/gotime",
		"http://qvault.io/",
		"https://golang.ch/",
		"https://gosamples.dev/",
		"https://golangcode.com",
		"https://appliedgo.net",
	}

	for _, link := range links {
		go accessLink(link, channel)
	}


	for message := range channel {
		fmt.Println(message)
	}


}
func accessLink(link string, channel chan string) {
	if _, error := http.Get(link); error != nil {
		channel <- link + " is not accessible"

	} else {
		channel <- link + " is accessible"
	}

}
