package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){

	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	str := input.Text()

	var arr []int
	var countNum int

	for i:=0;i<len(str);i++{
		if !inArr02(arr,int(str[i])){
			arr = append(arr,int(str[i]))
			countNum++
		}
	}
	fmt.Println(countNum)
}


func inArr02(arr []int, value int) bool{
	var isInArr = false
	for _,v:=range arr{
		if v==value{
			isInArr = true
		}
	}
	return isInArr
}
