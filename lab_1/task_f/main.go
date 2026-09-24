package main

import (
	"fmt"
	"slices"
)

func main() {
	var n int
	fmt.Scan(&n)

	nums := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&nums[i])
	}

	fmt.Println(slices.Concat(nums[n-1:], nums[:n-1]))
}
