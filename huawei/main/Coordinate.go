package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var str string
	fmt.Scan(&str)
	coordinateArr := strings.Split(str,";")
	arrLen := len(coordinateArr)
	var x,y int
	for i:=0;i<arrLen;i++{
		item := coordinateArr[i]
		if len(item) <2{
			continue
		}
		coord := string(item[0])
		number, err := strconv.Atoi(item[1:])
		if err == nil{
			switch coord{
			case "A":
				x -= number
			case "D":
				x += number
			case "S":
				y -= number
			case "W":
				y += number
			default:
			}
		}
	}
	fmt.Print(x,",",y)
}