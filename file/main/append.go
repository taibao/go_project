package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	filePath := "2.txt"
	//O_TRUNC 打开文件时会把文件清空
	file,err := os.OpenFile(filePath,os.O_WRONLY|os.O_APPEND,0666)
	if err != nil{
		fmt.Printf("open file err =%v\n",err)
		return
	}

	//及时关闭file句柄
	defer file.Close()

	//准备写入5句 “hello ，gardon”
	str := "hello,南京\r\n"
	//写入时，使用带缓存的*Writer
	writer := bufio.NewWriter(file)
	for i:=0;i<5;i++{
		writer.WriteString(str)
	}

	//因为writer是带缓存的，因此在调用writeString方法时，其实是将内容先写入到缓存的
	//所以需要调用Flush方法，将缓冲的数据
	//真正的写入到文件中， 否则文件会丢失数据
	writer.Flush()
}
