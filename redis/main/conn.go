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
	_,err = c.Do("Set","name","vitas")
	if err != nil{
		fmt.Println("数据设置失败",err)
	}

	//3：通过go向redis读取数据 string[key-val],并做类型转换
	r,err := redis.String(c.Do("Get","name"))
	if err != nil{
		fmt.Println("set err=",err)
		return
	}
	//因为返回的r是interface{}
	//因为name对应的是string，因此我们需要转换
	fmt.Println("操作ok",r)
}
