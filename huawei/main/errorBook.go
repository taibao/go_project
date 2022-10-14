package main

import (
	"fmt"
	"strings"
)

//维护一个大小为8的队列， 两个栈拼成一个队列
type queue struct{
	nums []string
	hh,tt int
}

func (q *queue) push(val string) {
	n := len(q.nums)
	if q.isFull() {
		//出栈
		if q.tt == n-1 {
			q.tt = 0
		} else {
			q.tt++
		}
	}
	//入栈
	q.nums[q.hh] = val
	if q.hh == n-1 {
		q.hh = 0
	} else {
		q.hh++
	}
}
func (q *queue) isFull() bool {
	n := len(q.nums)
	return (q.tt+n-1)%n == q.hh //满的状态下hh永远在tt前一位
}

//从尾部弹出数据
func (q *queue) pop() string {
	n := len(q.nums)
	ans := q.nums[q.tt]
	if q.tt == n-1 {
		q.tt = 0
	} else {
		q.tt++
	}
	return ans
}
func (q *queue) isEmpty() bool {
	return q.hh == q.tt
}
func main() {
	q := queue{nums: make([]string, 9)}
	s := ""
	a := 0
	mp := make(map[string]int)
	for {
		n, _ := fmt.Scanln(&s, &a)
		if n == 0 {
			break
		}
		index := strings.LastIndex(s, "\\")
		if index >= 0 {
			s = s[index+1:]
		}
		if len(s) > 16 {
			s = s[len(s)-16:]
		}
		s = fmt.Sprintf("%s %d", s, a) //将文件名和行数重新拼在一起
		if _, ok := mp[s]; !ok {
			q.push(s) //入队列
		}
		mp[s]++
	}

	fmt.Println("输出队列容量", len(q.nums))

	//相等的时候不弹出数据，导致空间虽然为9但是只弹出8条
	for !q.isEmpty() {
		st := q.pop()
		fmt.Printf("%s %d\n", st, mp[st])
	}
}
