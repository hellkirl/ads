package main

import (
	"bufio"
	"fmt"
	"os"
)

func maxWaterLiters(a []int) int64 {
	left, right := 0, len(a)-1
	if left >= right {
		return 0
	}

	leftMax, rightMax := 0, 0
	var water int64

	for left < right {
		if a[left] <= a[right] {
			if a[left] >= leftMax {
				leftMax = a[left]
			} else {
				water += int64(leftMax - a[left])
			}
			left++
		} else {
			if a[right] >= rightMax {
				rightMax = a[right]
			} else {
				water += int64(rightMax - a[right])
			}
			right--
		}
	}

	return water
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n uint32
	if _, err := fmt.Fscan(reader, &n); err != nil {
		return
	}

	a := make([]int, n)
	for i := uint32(0); i < n; i++ {
		if _, err := fmt.Fscan(reader, &a[i]); err != nil {
			return
		}
	}

	fmt.Println(maxWaterLiters(a))
}
