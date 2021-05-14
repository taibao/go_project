package utils

import "github.com/garyburd/redigo/redis"

//定义一个全局的pool
var pool *redis.Pool
func initPool(address string,maxIdle,maxActive int,){
	pool = &redis.Pool{

	}
}