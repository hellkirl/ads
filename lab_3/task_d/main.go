package main

import (
	"bufio"
	"fmt"
	"os"
)

type Node struct {
	idx  int
	next *Node
}

type Stack struct {
	top *Node
}

func NewStack() *Stack {
	return &Stack{top: nil}
}

func (s *Stack) Push(idx int) {
	s.top = &Node{idx: idx, next: s.top}
}

func (s *Stack) Pop() (int, bool) {
	if s.top == nil {
		return 0, false
	}
	v := s.top.idx
	s.top = s.top.next
	return v, true
}

func (s *Stack) Peek() (int, bool) {
	if s.top == nil {
		return 0, false
	}
	return s.top.idx, true
}

func maxRectangleArea(a []int) int64 {
	n := len(a)
	if n == 0 {
		return 0
	}

	stack := NewStack()
	var maxArea int64

	for i := 0; i <= n; i++ {
		currHeight := 0
		if i < n {
			currHeight = a[i]
		}

		for {
			if stack.top == nil {
				break
			}

			topIdx, _ := stack.Peek()
			topHeight := a[topIdx]

			if currHeight < topHeight {
				stack.Pop()

				prevIdx := -1
				if stack.top != nil {
					prevIdx, _ = stack.Peek()
				}

				width := i - prevIdx - 1
				area := int64(topHeight * width)
				if area > maxArea {
					maxArea = area
				}
				continue
			}
			break
		}

		if i < n {
			stack.Push(i)
		}
	}

	return maxArea
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

	fmt.Println(maxRectangleArea(a))
}
