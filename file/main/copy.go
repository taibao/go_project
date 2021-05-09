package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

//将一个文件内容写入到另一个文件中
func main(){
	file1Path := "1.txt"
	file2Path := "2.txt"

	//O_TRUNC 打开文件时会把文件清空
	data,err := ioutil.ReadFile(file1Path)
	if err != nil{
		fmt.Println("read file err =%v\n",err)
		return
	}

	err = ioutil.WriteFile(file2Path,data,0666)
	if err != nil{
		fmt.Println("write file error=%v\n",err)
	}
}

//判断文件是否存在
func PathExists(path string) (bool,error){
	_,err := os.Stat(path)
	if err == nil{ //文件或目录存在
		return true,nil
	}
	if os.IsNotExist(err){
		return false,nil
	}
	return false,err
}





