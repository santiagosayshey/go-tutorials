package main

import (
	"fmt"
	"time"
)

func main() {
	var c = make(chan int, 5)
	go process(c)
	// this just means keep going until the channel closes
	for i := range c {
		fmt.Println(i)
		time.Sleep(time.Second * 1)
	}

}

func process(c chan int) {
	// close the pipe when the process exits
	defer close(c)
	for i := 0; i < 5; i++ {
		c <- i
	}
	fmt.Println("Exiting process")
}
