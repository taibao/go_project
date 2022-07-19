package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)


func main(){
		//go func(){
		//	fmt.Println("hello")
		//	panic("这是例子")
		//}()

	//Go(func(){
	//	fmt.Println("hello")
	//	panic("这是例子")
	//})

	//fmt.Println("vim-go")
	//
	//Go(func(){
	//	fmt.Println("hello")
	//	panic("这是例子")
	//})

	//time.Sleep(5*time.Second)

	//
	//var i int
	//
	//go func(){
	//	i = 2
	//}()
	//
	//fmt.Println(i)

		//time.Sleep(5*time.Second)
	//err := 1/0
	//fmt.Errorf("输出错误 %v：%v,%w", "类型","0不能作为除数",err)


	//var res []interface{}
	var wg sync.WaitGroup
	var Chan = make(chan interface{},1)

	for i := 0; i < 10; i++ {
		wg.Add(1) //定义WaitGroup队列长度，长度为1
		go mytest(&wg,Chan)
	}

	var res []interface{}
	for i := 0; i < 10; i++  {
		res = append(res,<-Chan)
	}
	fmt.Println("读出dataChan顺序", res)
	wg.Wait() //主程序最后，等待WaitGroup队列为0时再退出
	return
}



func mytest(wg *sync.WaitGroup, dataChan chan interface{}) {

	rand.Seed(time.Now().UnixNano())
	num := rand.Intn(30000)

	dataChan <- num
	fmt.Println("写入dataChan顺序", num)

	time.Sleep(1*time.Second)
	wg.Done() //相当于wg.Add(-1)，一个goroutine结束，调用wg.Done()，队列减一

}


	func Go(x func()){
		go func(){
			defer func(){
				if err := recover(); err != nil{
					fmt.Println(err)
				}
			}()
			x()
		}()
	}
