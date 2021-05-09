package main

import (
	"encoding/json"
	"fmt"
)

//定义一个结构体

type Monster2 struct{
	Name string `json:"name"`
	Age int
	Birthday string
	Sal float64
	Skill string
}


//演示将json字符串，反序列化成struct
func unmarshalStruct(){
	//说明str在项目开发中，是通过网络传输获取到..或者是读取文件获取到
	str := "{\"name\":\"牛魔王\",\"Age\":500,\"Birthday\":\"2021-01-10\",\"Sal\":9000,\"Skill\":\"法相天地\"}"

	var monster Monster2
	err := json.Unmarshal([]byte(str),&monster)
	if err != nil{
		fmt.Println("反序列化错误")
	}
	fmt.Println(monster,monster.Name)
}

//演示将json字符串，反序列化成map
func unmarshalMap(){
	//说明str在项目开发中，是通过网络传输获取到..或者是读取文件获取到
	str := "{\"name\":\"牛魔王\",\"Age\":500,\"Birthday\":\"2021-01-10\",\"Sal\":9000,\"Skill\":\"法相天地\"}"
	//定义一个map
	var a map[string] interface{}
	//反序列化 ,不需要make，make操作已经封装到unmarshal函数中了
	err := json.Unmarshal([]byte(str),&a)
	if err != nil{
		fmt.Printf("unmarshal err=%v\n",err)
	}
	fmt.Printf("反序列化后，a=%v\n",a)
}

//演示将json字符串，反序列化成Slice
func unmarshalSlice(){
	//说明str在项目开发中，是通过网络传输获取到..或者是读取文件获取到
	str := "[{\"address\":\"火云洞\",\"age\":30,\"name\":\"红孩儿\"}," +
		"{\"address\":\"白鹿洞\",\"age\":26,\"name\":\"白鹿公主\"}]"
	//定义一个slice
	var slice []map[string] interface{}
	//反序列化 ,不需要make，make操作已经封装到unmarshal函数中了
	err := json.Unmarshal([]byte(str),&slice)
	if err != nil{
		fmt.Printf("unmarshal err=%v\n",err)
	}
	fmt.Printf("反序列化后，a=%v\n",slice)
}


func main(){
	unmarshalSlice()
}
