package main

import (
	"bufio"
	"fmt"
	"os"
)

type node struct {
	next  *node
	state uint8
}

func readInt(r *bufio.Reader) (int, bool) {
	c, err := r.ReadByte()
	for err == nil && (c == ' ' || c == '\n' || c == '\r' || c == '\t') {
		c, err = r.ReadByte()
	}
	if err != nil {
		return 0, false
	}
	neg := false
	if c == '-' {
		neg = true
		c, err = r.ReadByte()
	}
	x := 0
	for err == nil && c >= '0' && c <= '9' {
		x = x*10 + int(c-'0')
		c, err = r.ReadByte()
	}
	if neg {
		x = -x
	}
	return x, true
}

func loopLength(list []node) int {
	for i := range list {
		start := &list[i]
		if start.state != 0 {
			continue
		}

		cur := start
		for cur != nil && cur.state == 0 {
			cur.state = 1
			cur = cur.next
		}

		if cur != nil && cur.state == 1 {
			length := 1
			for p := cur.next; p != cur; p = p.next {
				length++
			}
			return length
		}

		for p := start; p != nil && p.state == 1; p = p.next {
			p.state = 2
		}
	}
	return -1
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)

	n, ok := readInt(reader)
	if !ok || n <= 0 {
		fmt.Println(-1)
		return
	}

	list := make([]node, n)
	for i := 0; i < n; i++ {
		v, ok := readInt(reader)
		if !ok {
			break
		}
		j := v - 1
		if v < 1 || v > n || j == i {
			continue
		}
		list[i].next = &list[j]
	}

	fmt.Println(loopLength(list))
}
