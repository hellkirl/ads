package main

import (
	"bufio"
	"fmt"
	"os"
)


func main() {
	reader := bufio.NewReader(os.Stdin)

	var n int
	if _, err := fmt.Fscan(reader, &n); err != nil {
		return
	}

	a := make([]string, n)
	for i := range a {
		if _, err := fmt.Fscan(reader, &a[i]); err != nil {
			return
		}
	}

	const maxLen = 100
	var left, right [maxLen + 1]int
	balanced := 0

	for _, s := range a {
		c, o := 0, 0
		for _, ch := range s {
			if ch == '(' {
				o++
			} else if o > 0 {
				o--
			} else {
				c++
			}
		}

		switch {
		case c == 0 && o == 0:
			balanced++
		case c == 0:
			left[o]++
		case o == 0:
			right[c]++
		}
	}

	pairs := balanced / 2
	for k := 1; k <= maxLen; k++ {
		pairs += min(left[k], right[k])
	}

	fmt.Println(pairs)
}
