package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true

}

func main() {

	var wg sync.WaitGroup
	wg.Add(1)
	done := make(chan bool, 1)
	go func() {
		worker(done)
		wg.Done()
	}()

	wg.Wait()

	//<-done
	//fmt.Scanln()
}
