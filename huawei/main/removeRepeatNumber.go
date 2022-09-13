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

	var arr []string

	for _,v := range str{
		if !inArr(arr,string(v)){
			fmt.Print(string(v))
		}else{
			arr = append(arr,string(v))
		}
	}
	fmt.Println()
}

func inArr(arr []string, value string) bool{
	fmt.Println(arr)
	var isInArr = false
	for _,v:=range arr{
		if v==value{
			isInArr = true
		}
	}

	fmt.Println(isInArr)
	return isInArr
}
