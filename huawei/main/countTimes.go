package main

import(
	"bufio"
	"fmt"
	"os"
	"strings"

	//"strings"
)

func CountScan(str *string){
	input:=bufio.NewScanner(os.Stdin)
	input.Scan()
	*str = input.Text()
}

func main(){
	input:=bufio.NewScanner(os.Stdin)
	input.Scan()
	str := input.Text()
	fmt.Println("字符串",str)
	input.Scan()
	key := input.Text()
	fmt.Println("关键字",key)
	key = strings.ToLower(key)
	keyB1 := int32(key[0])
	var count int
	for _,v := range strings.ToLower(str){
		if v == keyB1{
			count++
		}
	}
	fmt.Println("输出字符串",count)

}