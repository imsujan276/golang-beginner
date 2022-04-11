package main

import (
	"fmt"
	"time"
)

//channel of string
var chan1 = make(chan string)
var chan2 = make(chan string)

func main() {
	// start task1 and task2 as go routines
	go task1()
	go task2()

	for i := 0; i < 2; i++ {
		// wait for information to be recived by the particular channel
		select {
		//set the value from channel chan1 to variable msg1
		case msg1 := <-chan1:
			fmt.Println("received from task1", msg1)
		case msg2 := <-chan2:
			fmt.Println("received from task2", msg2)
		}
	}
}

//task1 sleeps for 1 sec and send steinf "one" to the channel
func task1() {
	time.Sleep(1 * time.Second)
	chan1 <- "One"
}

//task2 sleeps for 2 sec and send steinf "two" to the channel
func task2() {
	time.Sleep(2 * time.Second)
	chan1 <- "Two"
}
