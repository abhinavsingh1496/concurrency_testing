package main

import "fmt"

func worker(jobs <-chan int, result chan<- int) {
	for n := range jobs {
		result <- fibo(n)
	}
}

func fibo(num int) int {
	if num <= 1 {
		return num
	}
	return fibo(num-1) + fibo(num-2)
}

func main() {
	jobs := make(chan int, 100)
	result := make(chan int, 100)
	go worker(jobs, result)
	for i := 1; i < 80; i++ {
		jobs <- i
	}
	close(jobs)

	for j := 1; j < 80; j++ {
		fmt.Println(<-result)
	}

}
