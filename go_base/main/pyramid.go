package main

import (
	"fmt"
	"strings"
)

func main(){
	//var str2 string = "hello北"
	//r := []rune(str2)
	//for i :=0; i< len(r);i++{
	//	fmt.Printf("字符=%c\n",r[i])
	//}

	//字符串转整数 n,err := strconv.Atoi("12")
	//n,err := strconv.Atoi("hello")
	//if err != nil{
	//	fmt.Println("转换错误",err)
	//}else{
	//	fmt.Println("转成的结果是",n)
	//}


	//整数转字符串 str = strconv.Itoa(12345)
	//str := strconv.Itoa(12345)
	//fmt.Printf("str=%v,str=%T",str,str)

	//字符串转[]byte： var bytes = []byte("hello go")
	//var bytes = []byte("hello go")
	//fmt.Printf("bytes=%v \n",bytes)

	//byte转字符串
	//str := string([]byte{97,98,99})
	//fmt.Printf("st=%v\n",str)

	//10进制转2,8,16进制
	//str := strconv.FormatInt(123,2)
	//fmt.Printf("123对应的二进制是=%v\n",str)
	//
	//str2 := strconv.FormatInt(123,16)
	//fmt.Printf("123对应的16进制是=%v\n",str2)

	//查找子串是否在指定的字符串中：strings.Contains("seafood","foo") //true
	//b := strings.Contains("seafood","yes")
	//fmt.Printf("b=%v\n",b)

	//统计一个字符串中有几个指定的子串
	//num := strings.Count("cheese","e")
	//fmt.Println(num)

	//不区分大小写的字符串比较（==区分字母大小）
	//判断两个字符串是否相等
	//fmt.Println(strings.EqualFold("abc","Abc"))

	//返回子串在字符串中第一次出现的位置
	//index := strings.Index("NLT_abc","eabc")
	//fmt.Println(index)

	//将指定的子串替换成另一个,用n表示要替换多少个
	//str :=strings.Replace("go go hello ","go","na ",2)
	//fmt.Printf("str=%v\n",str)

	//将字符串拆分成数组
	//strArr := strings.Split("hello,world,ook",",")
	//for i:=0; i<len(strArr);i++{
	//	fmt.Println(strArr[i])
	//}

	//将字符串的字母进行大小写的转换
	//str:= strings.ToLower("Go To Change World")
	//fmt.Println(str)

	//16 将字符串左右两个的空格去掉：string.TrimSpace("")
	//str := strings.Trim("! hello !"," !") //["hello"] 将左右两边指定字符去掉，中间去不掉
	//str := strings.TrimSpace("! hello ") //["hello"] 将左右两边和" "去掉
	//strl := strings.TrimLeft("! hello!","!") //
	//strr := strings.TrimRight("! hello!","!") //
	//fmt.Println(strl,strr)

	//判断字符串是否以指定的字符串开头
	hasStart :=  strings.HasPrefix("ftp://192.168.10.1","ftp") //true
	hasEnd	 := strings.HasSuffix("NLT_abc.jpg","abc") //false
	fmt.Println(hasStart,hasEnd)

}




