package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func main(){
	//先从pool取出一个连接
	conn := pool.Get()
	defer conn.Close()

	_,err := conn.Do("Set","name","tom")
	if err != nil{
		fmt.Println("conn.Do err = ",err)
		return
	}

	//取出
	user,err := redis.String(conn.Do("Get","name"))
	if err != nil{
		fmt.Println("conn.Do err=",err)
		return
	}

	fmt.Println("user=",user)

	//如果我们要从pool取出连接，一定要保证连接池时没有关闭

}

//定义全局变量的pool
var pool *redis.Pool

//当启动程序时，就初始化连接池

func init(){
	pool = &redis.Pool{
		MaxIdle:8, //最大空闲连接数
		MaxActive:0, //表示和数据库的最大连接数，0表示没有限制
		IdleTimeout:100, //最大空闲时间
		Dial:func()(redis.Conn,error){
			//初始化连接的代码，连接哪个ip
			return redis.Dial("tcp","localhost:6397")
		},
	}
}
