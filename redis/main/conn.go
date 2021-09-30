package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func main(){
	c,err := redis.Dial("tcp","localhost:6379")
	if err != nil{
	fmt.Println("conn redis failed",err)
	return
	}

	//输入链接密码
	//if _, err := c.Do("AUTH", "123"); err != nil {
	//	c.Close()
	//	fmt.Println(err)
	//	return
	//}

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


	//获取redis返回值
	 i2,err2 := redis.Int(c.Do("Get","i2"))
	i2++
	_, _ = c.Do("Set", "i2", i2)
	_,err2 = c.Do("expire","i2",3)
	if err2 != nil{
		fmt.Println("数据设置失败",err2)
	}

	 fmt.Println(i2)


}
