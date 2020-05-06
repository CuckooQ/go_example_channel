package main

import (
	"fmt"
	"time"
)

func main () {
	channel := make(chan bool)
	people := [2]string{"cuckoo", "q"}
	for _, person := range people {
		go isSexy(person, channel)	
	}
	fmt.Println(<- channel)
	fmt.Println(<- channel)
}

func isSexy(person string, channel chan bool) {
	fmt.Println(person)
	time.Sleep(time.Second * 5)
	channel <- true
}