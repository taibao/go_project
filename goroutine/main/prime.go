package main

import (
	"fmt"
	"time"
)

func putNum(intChan chan int) {
	for i := 1; i <= 300000; i++ {
		intChan <- i
	}
	//关闭intChan
	close(intChan)
}

//从intChan取出数据，并判断是否为素数，如果是就放入到primeChan
func primeNum(intChan chan int , primeChan chan int ,exitChan chan bool){
	//使用for循环
	var flag bool
	for{
		num,ok := <-intChan
		if !ok{ //取不到值则退出
			break
		}
		flag = true //假设是素数
		//判断num是不是素数
		for i:=2;i<num;i++{
			if num %i == 0{
				//说明该num不是素数
				flag = false
				break
			}
		}

		if flag{
			primeChan <- num
		}
	}

	fmt.Println("有一个primeNum协程因为取不到数据退出")
	//这里不能关闭primeChan
	//向exitChan写入true
	exitChan <- true

}


func main(){

	start := time.Now().Unix()
	intChan := make(chan int,10000)
	primeChan := make(chan int,20000) //放入结果

	//标识退出的管道
	exitChan := make(chan bool,4) //4个

	//开启一个协程，向intChan放入1-8000个数
	go putNum(intChan)
	//开启4个协程，从intChan取出数据，并判断是否为素数，如果是就放入到primeChan
	for i := 0;i<30;i++{
		go primeNum(intChan,primeChan,exitChan)
	}

	//当我们从exitChan取出了4个结果，皆可以放心的关闭primeNum
	go func (){
		for i:=0;i<30;i++{
			<-exitChan
		}
		close(primeChan)
	}()

	i:=0
	for{
		_,ok := <-primeChan
		i++
		if !ok{
			break
		}
	}
	fmt.Printf("素数有%v个\n",i)

	end := time.Now().Unix()
	fmt.Printf("main线程退出,耗时%v秒\n",end-start)



}
