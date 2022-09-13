package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main(){
	input := bufio.NewScanner(os.Stdin)

	input.Scan()
	rows, _ := strconv.Atoi(input.Text())
	map1 := make(map[string]int64)
	for i:=0; i<rows;i++{
		input.Scan()
		itemStr := input.Text()
		arr := strings.Fields(itemStr)
		value, _ :=  strconv.Atoi(arr[1])
		if inMAP(map1, arr[0]){
			map1[arr[0]] += int64(value)
		}else{
			map1[arr[0]] = int64(value)
		}
	}
	arr1 := sortMap(map1)
	sort.Ints(arr1)
	for _, v := range arr1{
		value := strconv.Itoa(v)
		fmt.Println(v," ",map1[value])
	}
}

func inMAP(arr map[string]int64, key string) bool{
	_,ok := arr[key]
	return ok
}

func sortMap(map1 map[string]int64) (arr []int) {
	for k,_:=range map1{
		k, _ := strconv.Atoi(k)
		arr = append(arr,k)
	}
	return arr
}