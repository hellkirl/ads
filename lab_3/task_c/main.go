package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Node struct {
	value [26]int
	next  *Node
}

type LinkedStack struct {
	top *Node
}

func NewLinkedStack() *LinkedStack {
	return &LinkedStack{}
}

func (s *LinkedStack) Push(v [26]int) {
	n := &Node{value: v, next: s.top}
	s.top = n
}

func (s *LinkedStack) Pop() ([26]int, bool) {
	if s.top == nil {
		return [26]int{}, false
	}
	v := s.top.value
	s.top = s.top.next
	return v, true
}

func parseNumber(s string, i int) (int, int) {
	if i >= len(s) || s[i] < '0' || s[i] > '9' {
		return 1, i
	}
	num := 0
	for i < len(s) && s[i] >= '0' && s[i] <= '9' {
		num = num*10 + int(s[i]-'0')
		i++
	}
	return num, i
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	s, _ := reader.ReadString('\n')
	s = strings.TrimSpace(s)

	stack := NewLinkedStack()
	stack.Push([26]int{})

	i := 0
	for i < len(s) {
		ch := s[i]

		switch {
		case 'A' <= ch && ch <= 'Z':
			idx := int(ch - 'A')
			i++

			cnt, nextI := parseNumber(s, i)
			i = nextI

			top, _ := stack.Pop()
			top[idx] += cnt
			stack.Push(top)

		case ch == '(':
			stack.Push([26]int{})
			i++

		case ch == ')':
			inner, _ := stack.Pop()
			i++

			mult, nextI := parseNumber(s, i)
			i = nextI

			for j := 0; j < 26; j++ {
				inner[j] *= mult
			}

			outer, _ := stack.Pop()
			for j := 0; j < 26; j++ {
				outer[j] += inner[j]
			}
			stack.Push(outer)

		default:
			i++
		}
	}

	final, _ := stack.Pop()

	var sb strings.Builder
	for i := 0; i < 26; i++ {
		if final[i] > 0 {
			sb.WriteByte(byte('A' + i))
			if final[i] != 1 {
				sb.WriteString(fmt.Sprintf("%d", final[i]))
			}
		}
	}

	fmt.Println(sb.String())
}
