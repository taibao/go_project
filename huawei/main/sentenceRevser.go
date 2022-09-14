package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func main(){

	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	s := input.Text()
	strArr := strings.Fields(s)

	for i:=0;i<int(math.Ceil(float64(len(strArr)/2)));i++{
		var temp string
		pos := len(strArr)-1-i //i的对称位置
		temp = strArr[pos]
		strArr[pos] = strArr[i]
		strArr[i] = temp
	}

	for i:=0;i<len(strArr);i++{
		fmt.Print(strArr[i]," ")
	}
}