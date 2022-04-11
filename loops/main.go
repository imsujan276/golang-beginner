package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"myapp/logger"
	"os"
	"time"
)

func main() {
	// whileLoop()
	// infiniteLoopWithChannel()
	nestedLoop()

}

func whileLoop() {
	// seed random number generator
	rand.Seed(time.Now().Unix())
	i := 1000
	for i > 100 {
		i = rand.Intn(1000) + 1
		fmt.Println("i =>", i)

		if i > 100 {
			fmt.Println("i =>", i, "so loop keeps going")
		}
	}
	fmt.Println("Got", i, "and broke out of the loop")
}

func infiniteLoopWithChannel() {
	//read user input 5 times and write it to log

	reader := bufio.NewReader(os.Stdin)
	ch := make(chan string)

	go logger.ListenForLog(ch)

	fmt.Println("Enter Something")

	for i := 0; i < 5; i++ {
		fmt.Print("=> ")
		input, _ := reader.ReadString('\n')
		ch <- input
		time.Sleep(time.Second)
	}
}

func nestedLoop() {
	for i := 0; i <= 10; i++ {
		fmt.Print("i is ", i)
		for j := 1; j <= 3; j++ {
			fmt.Print("   j: ", j)
		}
		fmt.Println()
	}
}
