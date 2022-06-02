package main

import (
	"fmt"
	"time"
)

func main(){


	timeout := make(chan bool,1)
	go func() {
		time.Sleep(1 * time.Second)
		timeout <- true
	}()

	select {
		case <-timeout:
			fmt.Println("网络超时了哥")

	}
}
