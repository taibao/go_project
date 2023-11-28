package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

const (
	width  = 20
	height = 20
)

type point struct {
	x, y int
}

func main() {
	snake := []point{{10, 10}, {10, 9}, {10, 8}}
	food := point{rand.Intn(width), rand.Intn(height)}
	direction := point{1, 0}

	for {
		os.system("clear")
		draw(snake, food)
		time.Sleep(100 * time.Millisecond)

		newHead := point{snake[0].x + direction.x, snake[0].y + direction.y}
		if newHead == food {
			snake = append([]point{newHead}, snake...)
			food = point{rand.Intn(width), rand.Intn(height)}
		} else {
			snake = append([]point{newHead}, snake[:len(snake)-1]...)
		}

		if newHead.x < 0 || newHead.x >= width || newHead.y < 0 || newHead.y >= height || contains(snake, newHead) {
			break
		}
	}

	fmt.Println("Game Over!")
}

func draw(snake []point, food point) {
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if contains(snake, point{x, y}) {
				fmt.Print("■")
			} else if x == food.x && y == food.y {
				fmt.Print("★")
			} else {
				fmt.Print("□")
			}
		}
		fmt.Println()
	}
}

func contains(snake []point, p point) bool {
	for _, segment := range snake {
		if segment == p {
			return true
		}
	}
	return false
}
