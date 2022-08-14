package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	input := bufio.NewScanner(os.Stdin)
	//先输入数据个数
	for{
		input.Scan()
		s := input.Text()
		strSlice(s)

		break
	}
}

func strSlice(str string){
	size := 8
	for i:=0;i<len(str);i=i+size{
		if i+8>len(str){
			suffix := ""
			for j:=0;j<(i+8-len(str));j++{
				suffix += "0"
			}
			fmt.Println(str[i:]+suffix)
		}else{
			fmt.Println(str[i:i+8])
		}
	}
}

