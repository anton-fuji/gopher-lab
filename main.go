package main

import "fmt"

import "time"

func main() {
	ch := make(chan int) // 容量 0
	go func() {          // goroutine 1
		ch <- 100
	}()

	go func() { // goroutine 2
		v := <-ch
		fmt.Println(v)
	}()

	time.Sleep(2 * time.Second)
}
