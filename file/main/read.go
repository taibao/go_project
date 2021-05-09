package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	//打开文件
	//概念说明：file的叫法
	//1.file 叫file对象
	//2.file 叫file指针
	//3.file叫file文件句柄
	//file,err := os.Open("1.txt")
	//当函数退出时，要及时关闭file句柄，否则会有内存泄露
	//defer file.Close()
	//
	//if err != nil{
	//	fmt.Println("read file err=",err)
	//}
	//输出文件内容 通过输出内容可知file就是一个指针 *File
	//fmt.Printf("file=%v",file)
	//
	////关闭文件
	//err = file.Close()
	//if err != nil{
	//	fmt.Println("close file err=",err)
	//}


	//创建一个*Reader，是带缓冲的
	/*
	const (
		defaultBufSize = 4096 //默认的缓冲区为4096字节
	)
	*/

	//reader := bufio.NewReader(file)
	////循环的读取文件的内容
	//for{
	//	str,err := reader.ReadString('\n') //读到一个换行符就结束
	//	//输出内容
	//	fmt.Print(str)
	//
	//	//如果读到文件末尾
	//	if err == io.EOF{
	//		break
	//	}
	//}
	//fmt.Println("文件读取结束。。。")

	//一次性读取文件

	//把读取到的内容显示到终端
	//file := "1.txt"
	//content ,err := ioutil.ReadFile(file)
	//if err != nil{
	//	fmt.Printf("read file err = %v",err)
	//}
	//fmt.Printf("%v",string(content)) //[]byte 需要手动转换为字符串
	//因为我们没有显式的open文件，因此不需要显式的close文件
	//因为文件的open和close被封装到ReadFile函数内部

	//创建一个新文件，写入内容
	filePath := "2.txt"
	file,err := os.OpenFile(filePath,os.O_WRONLY|os.O_CREATE,0666)
	if err != nil{
		fmt.Printf("open file err =%v\n",err)
		return
	}

	//及时关闭file句柄
	defer file.Close()

	//准备写入5句 “hello ，gardon”
	str := "hello,beijing\r\n"
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



