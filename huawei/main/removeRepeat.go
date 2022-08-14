package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main(){
	input := bufio.NewScanner(os.Stdin)

	//先输入数据个数
	input.Scan()
	s := input.Text()
	num, _ := strconv.Atoi(s)
	var repeat []int
	for  i :=0;i<num;i++{
		input.Scan()
		randomNum, _ := strconv.Atoi(input.Text())
		if InArr(repeat, randomNum){
			continue
		}
		repeat = append(repeat,randomNum)
	}
	sort.Ints(repeat)
	for _,v := range repeat{
		fmt.Println(v)
	}
}

func InArr(arr []int, key int) bool{
	var inArr bool
	if arr == nil{
		return inArr
	}
	for _, v := range arr {
		if v == key{
			inArr=true
			break
		}
	}
	return inArr
}
