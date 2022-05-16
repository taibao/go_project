package main

import (
	"fmt"
	"github.com/pkg/errors"
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

	fmt.Println("vim-go")

	Go(func(){
		fmt.Println("hello")
		panic("一路向西")
	})

	//return errors.Wrap(ErrNotFound,"蜜桃成熟时4")
	//return errors.WithStack(ErrNotFound


	//time.Sleep(5*time.Second)
	//err := 1/0
	//fmt.Errorf("输出错误 %v：%v,%w", "类型","0不能作为除数",err)

}




var ErrNotFound error

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
