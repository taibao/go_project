package main

import (
	murmur3 "github.com/spaolacci/murmur3"
	"strconv"
)


func GetTableName(appId string, tagId int) string {
	num := murmur3.Sum32WithSeed([]byte(appId+strconv.Itoa(tagId)), 0) % 100
	return "t_resource_" + strconv.Itoa(int(num))
}