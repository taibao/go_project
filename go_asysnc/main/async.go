package main

import (
	"fmt"
)

func DoneAsync() {
	r := make(chan int)
	fmt.Println("Warming up ...")
	go func() {
		r <- 1
		fmt.Println("Done ...")
	}()

	<-r
	return
}

func main() {

	fmt.Println("Let's start ...")

	//var wg *sync.WaitGroup
	//wg.Add(2)
	DoneAsync()

	fmt.Println("Done is running ...")
	//wg.Wait()
}
