package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)



func main(){
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	s := input.Text()

	input.Scan()
	sub := input.Text()

	var count int

	s = strings.ToLower(s)
	sub  = strings.ToLower(sub)
	b1 := sub[0]

	for i:=0; i<len(s);i++{
		if s[i] == b1{
			count++
		}
	}
	fmt.Println(count)
}