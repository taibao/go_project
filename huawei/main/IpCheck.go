/**
 * 1. Mask 255.255.255.255 , 0.0.0.0 为非法。
 * 2. IP和Mask必须同时正确，才能被分类到A, B, C, D, E以及私有。
 * 3. IP和Mask同时错误时，只算一次错误
 * 4. 注意0.*.*.*以及127.*.*.*不属于任何类别。
 * 5. 不能把字符串和数字一样做比较: "53" > "123" // true
 */
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)


func main() {
	cnt := map[byte]int{}

	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		s := sc.Text()
		sp := strings.Split(s, "~")
		if strings.HasPrefix(sp[0], "0.") || strings.HasPrefix(sp[0], "127."){
			continue
		}

		ip := parse(sp[0]) //解析ip
		mask := parse(sp[1]) //解析掩码
		kind, private := CheckIp(ip)
		isMask := CheckMask(mask)

		if isMask{
			cnt[kind]++
			if private{
				cnt['g']++
			}
		}else{
			if kind !='n'{
				cnt['f']++
			}
		}
	}

	fmt.Printf("%d %d %d %d %d %d %d", cnt['a'], cnt['b'], cnt['c'], cnt['d'], cnt['e'], cnt['f'], cnt['g'])
}

/*
所有的IP地址划分为 A,B,C,D,E五类
A类地址从1.0.0.0到126.255.255.255
B类地址从128.0.0.0到191.255.255.255
C类地址从192.0.0.0到223.255.255.255
D类地址从224.0.0.0到239.255.255.255
E类地址从240.0.0.0到255.255.255.255

私网IP范围是：
从10.0.0.0到10.255.255.255
从172.16.0.0到172.31.255.255
从192.168.0.0到192.168.255.255
*/
func CheckIp(ip []int)(byte, bool){
	if !IsOk(ip){
		//如果有错则返回f
		return 'f',false
	}
	if ip[0] >= 1 && ip[0] <= 126{
		return 'a' , ip[0] ==10
	}
	if ip[0] >= 128 && ip[0]<= 191{
		return 'b', ip[0] == 172 && ip[1] >= 16 && ip[1] <= 31
	}
	if ip[0] >=192 && ip[0] <= 223{
		return 'c' , ip[0] == 192 && ip[1] == 168
	}
	if ip[0] >= 224 && ip[0] <= 239{
		return 'd' , false
	}
	if ip[0] >= 240 && ip[0] <= 255{
		return 'e' , false
	}
	return 'n', false
}

//判断是否为合法ip
func IsOk(ip []int) bool{
	if ip ==nil {
		return false
	}
	res := true
	for i:=0; i<len(ip); i++{
		res = res && ip[1]>=0 && ip[i]<=255
	}

	return res
}

/*
func CheckIp(ip []int) (byte, bool) {
	if !IsOk(ip){
		return 'f',false
	}

	if ip[0] >= 1 && ip[0] <= 126 {
		return 'a', ip[0] == 10
	}

	if ip[0] >= 128 && ip[0] <= 191 {
		return 'b', ip[0] == 172 && ip[1] >= 16 && ip[1] <= 31
	}

	if ip[0] >= 192 && ip[0] <= 223 {
		return 'c', ip[0] == 192 && ip[1] == 168
	}

	if ip[0] >= 224 && ip[0] <= 239 {
		return 'd', false
	}
	if ip[0] >= 240 && ip[0] <= 255 {
		return 'e', false
	}

	return 'n', false
}
*/
//检查mask 掩码
func CheckMask(ip []int) bool{
	if ip == nil{
		return false
	}
	var res uint32
	for i:=0;i<len(ip);i++{
		res = res << 8 + uint32(ip[i]) //把ip从高到低拼成一个完整数
	}

	if res == 0 || res == 0xffffffff{
			return false
	}

	const high = 0x80000000
	for {
		if (res & high) ==0{
			break //非掩码
		}
		res = res << 1
	}
	return res == 0
}


/*
func CheckMask(ip []int) bool {
	if ip ==nil{
		return false
	}
	var res uint32
	for i := 0; i < len(ip); i++ {
		res = res<<8 + uint32(ip[i]) //
	}

	if res == 0 || res == 0xffffffff{
		return false
	}

	const high = 0x80000000
	for {
		if (res & high) == 0 {
			break
		}
		res = res << 1
	}
	return res == 0
}
*/


/*
func IsOk(ip []int) bool{
	if ip ==nil {
		return false
	}
	res := true
	for i:=0; i<len(ip); i++{
		res = res && ip[1]>=0 && ip[i]<=255
	}

	return res
}
*/

//将ip解析为数组
func parse(s string) []int{
	sp := strings.Split(s,".")
	if len(sp) != 4{
		return nil
	}
	res := []int{}
	for i:=0;i<4;i++{
		val, err := strconv.Atoi(sp[i])
		if err != nil{
			return nil
		}
		res = append(res,val)
	}
	return res
}