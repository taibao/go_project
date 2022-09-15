package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	n, _ := strconv.Atoi(input.Text())

	var arr []string
	for i:=0;i<n;i++{
		input.Scan()
		s := input.Text()
		arr = append(arr,s)
	}

	arr = sortArr(arr)

	for i:=0;i<n;i++{
		fmt.Println(arr[i])
	}
}

func sortArr(arr []string) []string{
	for i:=0;i<len(arr)-1;i++{
		for j:=i;j<len(arr);j++{
			if arr[i]>arr[j]{
				var str string
				str = arr[j]
				arr[j] = arr[i]
				arr[i] = str
			}
		}
	}
	return arr
}