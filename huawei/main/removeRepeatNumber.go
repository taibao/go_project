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

	for i:=len(str)-1;i>=0;i--{
		if !inArr(arr,int(str[i])){
			arr = append(arr,int(str[i]))
		}
	}
	for _,v:=range arr{
		fmt.Print(string(v))
	}
	fmt.Println()
}

func inArr(arr []int, value int) bool{
	var isInArr = false
	for _,v:=range arr{
		if v==value{
			isInArr = true
		}
	}
	return isInArr
}
