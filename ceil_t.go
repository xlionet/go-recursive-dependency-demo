package main

import (
	"fmt"
	"math"
)

func main() {
	 ex := int64(1491656492)
	now := int64(1481288820)
	f := (ex - now)
	c :=int (math.Ceil(float64(f/60)))
	fmt.Println("c is:", c,"f:",f/60)
}
