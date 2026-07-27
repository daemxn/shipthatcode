package main

import (
	"fmt"
)

func main() {
	var nums []int
	for {
		var temp int
		_, err := fmt.Scan(&temp)
		if err != nil {
			break
		}
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
