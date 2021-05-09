package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Monster struct{
	Name string
	Age int
	Skill string
}

//给monster绑定方法store，可以将一个Monster变量（对象），序列化后保存到文件中
func (this *Monster) Store() bool{
	//直接将序列化字符串保存到文件中
	data,err := json.Marshal(this)
	if err != nil{
		fmt.Println("marshal err = ",err)
		return false
	}
	//保存到文件
	filePath := "3.txt"
	err = ioutil.WriteFile(filePath,data,0666) //写入文件
	if err != nil{
		fmt.Println("write file err=",err)
		return false
	}
	return true
}

//给Monster绑定方法restore
func (this *Monster) Restore() bool{
	//从文件中读取序列化的字符串
	filePath := "3.txt"
	data,err := ioutil.ReadFile(filePath) //读取文件
	if err != nil{
		fmt.Println("read file err=",err)  //读取文件失败
		return false
	}
	//使用读取到的data []byte,对反序列化
	err = json.Unmarshal(data,this)
	if err != nil{
		fmt.Println("UnMarshal err =",err) //反序列化失败
		return false
	}
	return true
}