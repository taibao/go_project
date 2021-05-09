package main

import (
	"fmt"
	"sync"
	"time"
)

//1-200累加
var (
	myMap = make(map[int]int,10)
	//声明一个全局的互斥锁
	//lock是一个全局的互斥锁
	lock sync.Mutex //同步互斥锁
)

//test函数就是计算n!,将结果放入mymap中
func test2(n int){
	res := 1
	for i := 1;i<= n ;i++{
		res += i
	}
	//将res放入map中，加锁
	lock.Lock()
	myMap[n] = res
	//解锁
	lock.Unlock()
}

func main(){
	//开启多个协程完成任务[200个]
	for i :=1;i<=200;i++{
		go test2(i)
	}

	//休眠10秒
	time.Sleep(time.Second*10)
	lock.Lock()
	for i,v :=range myMap {
		fmt.Printf("map[%d]=%d\n",i,v)
	}
	lock.Unlock()
}
