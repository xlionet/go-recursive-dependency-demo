package main

import (
	"strconv"
	"fmt"
)

func main() {
	m := make(map[int]string)
	for i := 0; i < 10; i++ {
		m[i] = strconv.Itoa(i) + "Hello, world~"
	}

	for k, v := range m {
		k += 1
		v += "1"
	}

	for k, v := range m {
		fmt.Println(k,v)
	}
}
