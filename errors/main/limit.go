package main

//该程序为工作列表中的每个条目启动了一个goroutine，
//这些goroutine使用limit通道进行协调，以确保一次最多运行三个工作函数

var limit  = make(chan int,3)

func main(){
	for _,w := range work{
		go func(w func()){
			limit <- 1
			w()
			<-limit
		}(w)
	}
	select{}
}