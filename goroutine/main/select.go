package main

import "fmt"

func main(){
	intChan := make(chan int,10)
	for i:=0;i<10;i++{
		intChan<-i
	}

	stringChan := make(chan string,5)
	for i:=0;i<5;i++{
		stringChan <- "hello" + fmt.Sprintf("%d",i)
	}

	//传统的方法在遍历管道时，如果不关闭会阻塞导致deadlock
	label :
	for{
		//使用select可以自动跳转到下一个case匹配
		select {
			case v := <- intChan:
				fmt.Printf("从intChan读取数据%d\n",v)
			case v := <- stringChan:
				fmt.Printf("从stringChan读取数据%s\n",v)
			default:
				fmt.Printf("都取不到了，不玩了\n")
				break label
		}
	}
}
