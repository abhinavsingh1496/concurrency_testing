package main

import (
	"fmt"
	"time"
)

func main() {
	chn1 := make(chan string)
	chn2 := make(chan string)

	go func() {
		for {
			chn1 <- "every 2 secs"
			time.Sleep(2 * time.Second)
		}
	}()

	go func() {
		for {
			chn2 <- "every 0.5 secs"
			time.Sleep(500 * time.Millisecond)
		}
	}()

	for {
		select {
		case msg1 := <-chn1:
			{
				fmt.Println(msg1)
			}
		case msg2 := <-chn2:
			{
				fmt.Println(msg2)
			}
		}
	}
}
