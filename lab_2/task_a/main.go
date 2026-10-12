package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

func main() {
	var s string
	fmt.Scan(&s)

	if utf8.RuneCountInString(s) > 100 {
		panic("too long")
	}

	fmt.Println(Archiver(s))
}

func Archiver(s string) string {
	n := len(s)

	length := make([][]int, n)
	cut := make([][]int, n)
	block := make([][]int, n)
	for i := 0; i < n; i++ {
		length[i] = make([]int, n)
		cut[i] = make([]int, n)
		block[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		length[i][i] = 1
		cut[i][i] = -1
	}

	for i := 2; i <= n; i++ {
		for j := 0; j <= n-i; j++ {
			end := j + i - 1
			bestK, bestP, bestLen := -1, 0, i
			for k := j; k < end; k++ {
				if l := length[j][k] + length[k+1][end]; l < bestLen {
					bestK, bestLen = k, l
				}
			}

			for p := 1; p < i; p++ {
				if i%p != 0 {
					continue
				}

				repeated := true
				for x := j + p; x <= end; x++ {
					if s[x] != s[x-p] {
						repeated = false
						break
					}
				}

				if repeated {
					if l := len(strconv.Itoa(i/p)) + 2 + length[j][j+p-1]; l < bestLen {
						bestK, bestP, bestLen = -1, p, l
					}
				}
			}

			length[j][end] = bestLen
			cut[j][end] = bestK
			block[j][end] = bestP
		}
	}

	var sb strings.Builder
	var build func(i, j int)
	build = func(i, j int) {
		switch {
		case cut[i][j] != -1:
			build(i, cut[i][j])
			build(cut[i][j]+1, j)
		case block[i][j] != 0:
			p := block[i][j]
			sb.WriteString(strconv.Itoa((j - i + 1) / p))
			sb.WriteByte('(')
			build(i, i+p-1)
			sb.WriteByte(')')
		default:
			sb.WriteString(s[i : j+1])
		}
	}
	build(0, n-1)

	return sb.String()
}
