package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n, k int
	if _, err := fmt.Fscan(reader, &n, &k); err != nil {
		return
	}

	a := make([]int, n)
	for i := range a {
		if _, err := fmt.Fscan(reader, &a[i]); err != nil {
			return
		}
	}

	p, q := FindSmallestFraction(n, k, a)
	fmt.Printf("%d %d\n", p, q)
}

func FindSmallestFraction(n, k int, a []int) (int, int) {
	low, high := 0.0, 1.0
	p, q := 0, 1

	for range 100 {
		mid := (low + high) / 2
		var cnt int64
		cnt, p, q = CountFractions(mid, a)
		if cnt == int64(k) {
			return p, q
		}
		if cnt < int64(k) {
			low = mid
		} else {
			high = mid
		}
	}

	return p, q
}

func CountFractions(mid float64, a []int) (cnt int64, p, q int) {
	n := len(a)
	p, q = 0, 1

	j := 1
	for i := 0; i < n-1; i++ {
		if j <= i {
			j = i + 1
		}
		for j < n && float64(a[i]) > mid*float64(a[j]) {
			j++
		}
		if j == n {
			break
		}

		cnt += int64(n - j)
		if a[i]*q > p*a[j] {
			p, q = a[i], a[j]
		}
	}

	return
}
