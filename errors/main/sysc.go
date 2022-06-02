package main

import "sync"

//sync包实现了两种锁数据类型， sync.Mutex和sync.RWMutex
//Lock返回之前就会开始调用unlock

var l sync.Mutex
var a2 string

func f2(){
	a2 = "hello world"
	l.Unlock()
}

//第一次调用l.Unlock()发生在第二次调用l.Lock()返回之前
//l.lock发生在调用l.unlock之后
func main(){
	l.Lock()
	go f2()
	l.Lock()
	print(a2)
}