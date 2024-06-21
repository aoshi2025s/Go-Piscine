package main

import (
	"fmt"
	"piscine"
)

// 構造体のマクロ定義が必要そう?
// point型

func main() {
	points := &point{}
	
	setPoint(points)

	fmt.Printf("x = %d, y = %d\n", points.x, points.y)
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

// go run .
// x = 42, y = 21
