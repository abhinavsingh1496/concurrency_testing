package main

import (
	"fmt"
	"time"
)

func input(chn chan int) {
	for i := 0; i < 10; i++ {
		time.Sleep(500 * time.Millisecond)
		chn <- i
	}
	close(chn)
}

func main() {
	chn := make(chan int)
	go func() {
		input(chn)
	}()
	for msg:= range chn{
		fmt.Println(msg)
	}
}
