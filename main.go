package main

import (
	"fmt"
)

func main() {
	var nums []int
	for i := 0; i < 5; i++ {
		var temp int
		fmt.Scan(&temp)
		nums = append(nums, temp)
	}

	max := nums[0]
	for _, v := range nums {
		if v > max {
			max = v
		}
	}

	fmt.Println(max)

}
