package main

import (
	"fmt"
)

func main() {
	var input int
	fmt.Scan(&input)

	result := square(input)
	fmt.Println(result)
}
func square(n int) int {
	return n * n
}
