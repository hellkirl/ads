package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)

	nums := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&nums[i])
	}

	var x int
	fmt.Scan(&x)

	closestIdx := 0
	for i := range nums {
		if math.Abs(float64(nums[i]-x)) < math.Abs(float64(nums[closestIdx]-x)) {
			closestIdx = i
		}
	}

	fmt.Println(nums[closestIdx])
}
