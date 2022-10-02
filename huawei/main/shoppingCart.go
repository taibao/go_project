package main

import (
	"fmt"
)

func main() {
	var money, count int
	fmt.Scan(&money, &count)
	goods := make([][]int, count)
	for i := 0; i < count; i++ {
		goods[i] = make([]int, 3)
	}
	for i := 0; i < count; i++ {
		for j := 0; j < 3; j++ {
			fmt.Scan(&goods[i][j])
		}
	}

	costs := make([][]int, 0)
	prices := make([][]int, 0)
	for i := 0; i < count; i++ {
		//附件暂不计算
		if  goods[i][2] != 0 {
			continue
		}
		//存入主件的价格和满意度
		cost := make([]int, 0)
		price := make([]int, 0)
		cost = append(cost, goods[i][0] * goods[i][1]) //主件的重要度
		price = append(price, goods[i][0]) //主件的价格

		//遍历所有行，加入当前主件的所有附件的满意度和价格
		for j := 0; j < count; j++ {
			//如果是当前主件的附件， 就把价钱和满意度合入到cost
			if goods[j][2]-1 == i {
				cost = append(cost, goods[j][0] * goods[j][1] + cost[0]) //保存当前附件加上主件的价格
				price = append(price, goods[j][0] + price[0]) //价格
			}
		}
		//如果主件和附件达到三件
		if len(cost) == 3 {
			//最后一个元素保存主件加上两个附件的满意度和价格
			cost = append(cost, cost[1] + cost[2] - cost[0])
			price = append(price, price[1] + price[2] - price[0])
		}
		costs = append(costs, cost)
		prices = append(prices, price)
	}
	DpCost(costs, prices, money)
}

func DpCost(cost, price [][]int, money int)  {
	count := len(price)
	packs := make([][]int, count)
	//为每个主件设置money
	for i := 0; i < count; i++ {
		temp := make([]int, 0)
		for j := 0; j <= money; j++ {
			temp = append(temp, -1)
		}
		packs[i] = temp
	}

	packs[0][0] = 0
	len1 := len(price[0]) //第一个主件的价格长度相当于主件数量
	for i := 0; i < len1; i++ {
		if price[0][i] <= money {
			packs[0][price[0][i]] = cost[0][i] //为相应的价格赋予相应的满意度
		}
	}

	for i := 1; i < count; i++ {
		lastPrice := 0
		//逐渐累积上一个主件的价格和满意度
		for j := 0; j <= money; j++ {
			if packs[i-1][j] >= 0 {
				packs[i][j] = packs[i-1][j]
				lastPrice = j
			}
		}

		for j := 0; j <= lastPrice; j++ {
			len2 := len(price[i])
			for k := 0; k < len2; k++ {
				if packs[i-1][j] == -1 {
					continue
				}

				sum := packs[i-1][j] + cost[i][k]
				if j+price[i][k] <= money && sum > packs[i][j+price[i][k]] {
					packs[i][j+price[i][k]] = sum
				}
			}
		}

	}

	//取package最后一行的记录即可
	max := 0
	for i := money; i >= 0; i-- {
		if packs[count-1][i] > 0 && packs[count-1][i] > max {
			max = packs[count-1][i]
		}
	}
	fmt.Println(max)
}
