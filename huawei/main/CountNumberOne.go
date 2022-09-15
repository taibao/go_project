package main

import "fmt"

func main(){
	n :=0
	_, _ = fmt.Scan(&n)

	num:=0
	for i:=0;i<32;i++{
		if n>>i&1==1{
			num++
		}
	}
	fmt.Println(num)
}