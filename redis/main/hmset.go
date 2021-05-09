package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"time"
)

func main(){
	c,err := redis.Dial("tcp","localhost:6397")
	if err != nil{
	fmt.Println("conn redis failed",err)
	return
	}

	//2：设置数据
	defer c.Close()
	_,err = c.Do("HMSet","user02","name","vitas","age",19)
	if err != nil{
		fmt.Println("HMSet err=",err)
		return
	}

	_,err = c.Do("expire","user02",3)
	time.Sleep(time.Second*4)

	//3：通过go向redis读取数据 string[key-val],并做类型转换
	user01,err := redis.Strings(c.Do("HMGet","user02","name","age"))
	if err != nil{
		fmt.Println("HMGet err=",err)
		return
	}
	fmt.Println("user01为",user01)

}
