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

		ip := parse(sp[0])
		mask := parse(sp[1])
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

func CheckIp(ip []int) (byte, bool) {
	if !IsOk(ip){
		return 'f', false
	}

	if ip[0] >= 0 && ip[0] <= 126 {
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

func CheckMask(ip []int) bool {
	if ip ==nil{
		return false
	}
	var res uint32
	for i := 0; i < len(ip); i++ {
		res = res<<8 + uint32(ip[i])
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

func parse(s string)[]int{
	sp := strings.Split(s, ".")
	if len(sp) != 4{
		return nil
	}

	res := []int{}
	for i:=0; i<4; i++{
		val, err:=strconv.Atoi(sp[i])
		if err !=nil{
			return nil
		}
		res = append(res, val)
	}
	return res
}