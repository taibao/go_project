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
		if  goods[i][2] != 0 {
			continue
		}
		cost := make([]int, 0)
		price := make([]int, 0)
		cost = append(cost, goods[i][0] * goods[i][1])
		price = append(price, goods[i][0])
		for j := 0; j < count; j++ {
			if goods[j][2]-1 == i {
				cost = append(cost, goods[j][0] * goods[j][1] + cost[0])
				price = append(price, goods[j][0] + price[0])
			}
		}
		if len(cost) == 3 {
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
	for i := 0; i < count; i++ {
		temp := make([]int, 0)
		for j := 0; j <= money; j++ {
			temp = append(temp, -1)
		}
		packs[i] = temp
	}

	packs[0][0] = 0
	len1 := len(price[0])
	for i := 0; i < len1; i++ {
		if price[0][i] <= money {
			packs[0][price[0][i]] = cost[0][i]
		}
	}

	for i := 1; i < count; i++ {
		lastPrice := 0
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

	//取出各种情况的最大值
	max := 0
	for i := money; i >= 0; i-- {
		if packs[count-1][i] > 0 && packs[count-1][i] > max {
			max = packs[count-1][i]
		}
	}
	fmt.Println(max)
}
