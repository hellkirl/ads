package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var s string
	if _, err := fmt.Fscan(reader, &s); err != nil {
		return
	}
	n := len(s)

	var total int64

	d1 := make([]int, n)
	for i, l, r := 0, 0, -1; i < n; i++ {
		k := 1
		if i <= r {
			k = min(d1[l+r-i], r-i+1)
		}
		for i-k >= 0 && i+k < n && s[i-k] == s[i+k] {
			k++
		}
		d1[i] = k
		total += int64(k)
		if i+k-1 > r {
			l, r = i-k+1, i+k-1
		}
	}

	d2 := make([]int, n)
	for i, l, r := 0, 0, -1; i < n; i++ {
		k := 0
		if i <= r {
			k = min(d2[l+r-i+1], r-i+1)
		}
		for i-k-1 >= 0 && i+k < n && s[i-k-1] == s[i+k] {
			k++
		}
		d2[i] = k
		total += int64(k)
		if i+k-1 > r {
			l, r = i-k, i+k-1
		}
	}

	fmt.Println(total)
}
