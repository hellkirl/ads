package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	nums := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&nums[i])
	}

	a, b := nums[0], nums[1]
	if a > b {
		a, b = b, a
	}

	for i := 2; i < n; i++ {
		if nums[i] < a {
			a, b = nums[i], a
		} else if nums[i] < b {
			b = nums[i]
		}
	}

	fmt.Println(a, b)
}
