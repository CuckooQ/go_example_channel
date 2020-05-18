package main

import (
	"fmt"
	"time"
)

func isSexy(person string, channel chan string) {
	time.Sleep(time.Second * 1)
	channel <- person + " is sexy"
}

func main() {
	channel := make(chan string)
	people := [3]string{"cuckoo", "q", "midory"}

	// go routine operation x 2
	for _, person := range people {
		go isSexy(person, channel)
	}

	// 2 channel results
	for i := 0; i < len(people); i++ {
		fmt.Println(<-channel)
	}
}
