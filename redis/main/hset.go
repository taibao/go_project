package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func main(){
	c,err := redis.Dial("tcp","localhost:6397")
	if err != nil{
	fmt.Println("conn redis failed",err)
	return
	}

	//2：设置数据
	defer c.Close()
	_,err = c.Do("HSet","user01","name","vitas")
	if err != nil{
		fmt.Println("HSet err=",err)
		return
	}
	_,err = c.Do("HSet","user01","age","19")
	if err != nil{
		fmt.Println("HSet err=",err)
		return
	}

	//3：通过go向redis读取数据 string[key-val],并做类型转换
	name,err := redis.String(c.Do("HGet","user01","name"))
	if err != nil{
		fmt.Println("HGet err=",err)
		return
	}
	fmt.Println("user01名字为",name)

	age,err := redis.String(c.Do("HGet","user01","age"))
	if err != nil{
		fmt.Println("HGet err=",err)
		return
	}
	fmt.Println("user01年龄为",age)
}
