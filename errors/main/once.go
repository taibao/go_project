package main

import "sync"

//sync包在通过使用类型为存在多个goroutine的情况下提供一种安全的初始化机制once，
//多个线程可以执行once.Do(f)一个特定的f，但只有一个会运行f(), 并且其他调用阻塞直到f()返回

var once sync.Once
var a3 string

func main(){
	twoprint()
}


func setup() {
	a3 = "hello world"

}

func doprint(){
	once.Do(setup) //只有一个线程会执行setup
	print(a3)
}

func twoprint(){
	go doprint()
	go doprint()
}