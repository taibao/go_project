package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main(){

	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	s := input.Text()

	str := []byte(s)
	for i:=0;i<int(math.Ceil(float64(len(str)/2)));i++{
		var temp byte
		pos := len(str)-1-i //i的对称位置
		temp = str[pos]
		str[pos] = str[i]
		str[i] = temp
	}
	fmt.Println(string(str))
}