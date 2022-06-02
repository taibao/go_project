package main

//var a3 string
//func f(){
//	print(a3)
//}
//
//func hello(){
//	a3 = "hello wordl"
//	time.Sleep(1*time.Second)
//
//	print("谁先执行")
//	 go f()
//}
//
//
//func hello1(){
//	//go内的变量没有同步事件， 碰上激进的编译器可能会删除整个go语句
//	go func(){a3="hello"}()
//	print(a3)
//}

//如果一个goroutine的效果必须被另一个goroutine观察到，
//可以使用锁或者通道通信等同步机制来建立相对顺序

// channel 通信

var a string
//var c = make(chan int ,10) //有缓冲
var c = make(chan int) //无缓冲


func f(str string){
	a = str
	<-c //开始发送
}

//来自无缓冲通道的接收发生在通道发送完成之前

func main(){
	go f("发送完成之前已经赋值了")
	c<-0 //开始接收 发送完成之前开始接收, 此时可以保证a的赋值已经完成。 操作非常骚
	println(a)

	//hello1()

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
	//	panic("一路向西")
	//})

	//return errors.Wrap(ErrNotFound,"蜜桃成熟时4")
	//return errors.WithStack(ErrNotFound


	//time.Sleep(5*time.Second)
	//err := 1/0
	//fmt.Errorf("输出错误 %v：%v,%w", "类型","0不能作为除数",err)

	//go func(){
	//	done<-serverApp()
	//}()
	//
	//for i:=0;i<cap(done);i++{
	//	<-done
	//	close(stop)
	//}
}





//func serverApp(stop chan struct{}) error{
//	go func(){
//		<-stop
//		http.Shutdown()
//	}()
//	return http.listen()
//}

//var ErrNotFound error
//
//func Go(x func()){
//	go func(){
//		defer func(){
//			if err := recover(); err != nil{
//				fmt.Println(err)
//			}
//		}()
//		x()
//	}()
//}
