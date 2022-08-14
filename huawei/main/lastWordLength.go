package main

import(
"bufio"
"fmt"
"os"
"strings"
)

//func LastStringLength(str *string) int {
//	newStr := strings.Fields(*str)
//	return len(newStr[len(newStr)-1])
//}
//
//func Scanf(input *string){
//	reader := bufio.NewReader(os.Stdin)
//	data,_,_ := reader.ReadLine()
//	*input = string(data)
//}

func Scanf(input *string){
	reader := bufio.NewReader(os.Stdin)
	data, _, _ := reader.ReadLine()
	*input = string(data)
}

func lastWordLen(input *string) int{
	strArr := strings.Fields(*input)
	return len(strArr[len(strArr)-1])
}


//获取一个输入字符串最后一个字符长度
func main(){
	var input string
	Scanf(&input) //接收字符串
	fmt.Printf("%d",lastWordLen(&input)) //输出长度
}



