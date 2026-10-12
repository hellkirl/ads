package main

import "fmt"

func main() {
	var n, k int
	if _, err := fmt.Scan(&n, &k); err != nil {
		return
	}

	a := make([]int, n)
	for i := range a {
		if _, err := fmt.Scan(&a[i]); err != nil {
			return
		}
	}

	fmt.Println(MinDistance(n, k, a))
}

func MinDistance(n, k int, a []int) int {
	low := 0
	high := a[n-1] - a[0]
	for low <= high {
		mid := (low + high) / 2
		if CanPlace(a, k, mid) {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return high
}

func CanPlace(a []int, k, distance int) bool {
	last, count := a[0], 1

	for i := 1; i < len(a); i++ {
		if a[i]-last >= distance {
			count++
			last = a[i]
		}
	}

	return count >= k
}
