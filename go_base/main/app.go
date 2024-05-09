package main

import (
	"fmt"
	"time"
)

type user struct {
	name string
	age  int
}

func main() {
	//var users user
	//fmt.Println("输出user", users)

	//getWeekday()

	//GetMonth()

	//getDay()
	//
	fmt.Println(GetDate("-360h"), GetDate("360h"))
}

func GetDate(goBack string) string {
	currentTime := time.Now()
	if goBack != "" {
		t, _ := time.ParseDuration(goBack)
		result := currentTime.Add(t)
		return result.Format("2006-01-02 00:00:00")
	}
	return currentTime.Format("2006-01-02 00:00:00")
}

func getWeekday() {
	// 获取当前时间
	currentTime := time.Now()

	// 获取星期几
	weekdayNum := int(currentTime.Weekday())
	if weekdayNum == 0 {
		weekdayNum = 7
	}

	// 输出星期几
	fmt.Println("Today is", weekdayNum)
}

func GetMonth() string {
	// 获取当前时间
	currentTime := time.Now()

	// 获取月份
	month := currentTime.Month()

	// 输出月份
	fmt.Println("Current month is", MonthInChinese(month))
	return MonthInChinese(month)
}

func MonthInChinese(month time.Month) string {
	switch month {
	case time.January:
		return "一月"
	case time.February:
		return "二月"
	case time.March:
		return "三月"
	case time.April:
		return "四月"
	case time.May:
		return "五月"
	case time.June:
		return "六月"
	case time.July:
		return "七月"
	case time.August:
		return "八月"
	case time.September:
		return "九月"
	case time.October:
		return "十月"
	case time.November:
		return "十一月"
	case time.December:
		return "十二月"
	default:
		return "未知"
	}
}

func getDay() int {
	// 获取当前时间
	currentTime := time.Now()
	// 获取日期
	day := currentTime.Day()
	// 输出日期
	fmt.Println("今天是", day, "号")
	return day
}
