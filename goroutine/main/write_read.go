package main

import (
	"fmt"
)

//write data
func writeData(intChan chan int){
	for i:=1;i<=50;i++{
		//放入数据
		intChan <- i
		fmt.Printf("writeData 写入数据=%v\n",i)
	}
	close(intChan)
}

//read data
func readData(intChan chan int,exitChan chan bool){
	for{
		v,ok := <-intChan
		if !ok{
			break
		}
		//time.Sleep(time.Second)
		fmt.Printf("readData 读到数据=%v\n",v)
	}

	//readData 读取完数据后，即认为任务完成
	exitChan<-true
	close(exitChan)
}

func main(){
	//创建两个管道
	intChan := make(chan int ,10)
	exitChan := make(chan bool,1)

	go writeData(intChan)
	go readData(intChan,exitChan)

	for{
		v,_ := <-exitChan
		if v{
			break
		}
	}
}
