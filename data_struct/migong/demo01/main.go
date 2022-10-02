package main
import (
	"fmt"
)

func test(n int) {
	if n > 1 {
		n-- //递归
		test(n)
	}else{
		return
	}
	fmt.Println("n=",n)
}

func main() {
	n := 6
	test(n)
}