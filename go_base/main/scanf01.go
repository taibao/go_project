package main

import (
"fmt"
"math"
"strconv"
)

func main(){
	var num string
	_, _ = fmt.Scan(&num)
	num2, _ := strconv.ParseFloat(num,64)
	if num2-math.Floor(num2) > 0.5{
		fmt.Println(math.Ceil(num2))
	}else{
		fmt.Println(math.Floor(num2))
	}
}
