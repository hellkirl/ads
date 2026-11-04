package main

import (
	"bufio"
	"fmt"
	"os"
)

type Operation string

const (
	OperationPush Operation = "+"
	OperationPop  Operation = "-"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n uint32
	if _, err := fmt.Fscan(reader, &n); err != nil {
		return
	}

	stack := NewStack()

	for i := uint32(0); i < n; i++ {
		var op Operation
		if _, err := fmt.Fscan(reader, &op); err != nil {
			return
		}

		switch op {
		case OperationPush:
			var x int32
			if _, err := fmt.Fscan(reader, &x); err != nil {
				return
			}
			stack.Push(x)
		case OperationPop:
			if val, ok := stack.Pop(); ok {
				fmt.Fprintln(writer, val)
			}
		}
	}
}

type Stack interface {
	Push(x int32)
	Pop() (int32, bool)
	Peek() (int32, bool)
	IsEmpty() bool
}

var _ Stack = &StackImpl{}

const minCap = 16

type StackImpl struct {
	nums []int32
	size int
}

func NewStack() Stack {
	return &StackImpl{
		nums: make([]int32, minCap),
		size: 0,
	}
}

func (s *StackImpl) IsEmpty() bool {
	return s.size == 0
}

func (s *StackImpl) Peek() (int32, bool) {
	if s.IsEmpty() {
		return 0, false
	}
	return s.nums[s.size-1], true
}

func (s *StackImpl) Pop() (int32, bool) {
	if s.IsEmpty() {
		return 0, false
	}
	s.size--
	val := s.nums[s.size]
	if len(s.nums) > minCap && s.size <= len(s.nums)/4 {
		s.resize(len(s.nums) / 2)
	}
	return val, true
}

func (s *StackImpl) Push(x int32) {
	if s.size == len(s.nums) {
		s.grow()
	}
	s.nums[s.size] = x
	s.size++
}

func (s *StackImpl) grow() {
	newNums := make([]int32, len(s.nums)*2)
	for i := 0; i < s.size; i++ {
		newNums[i] = s.nums[i]
	}
	s.nums = newNums
}

func (s *StackImpl) resize(capacity int) {
	newNums := make([]int32, capacity)
	for i := 0; i < s.size; i++ {
		newNums[i] = s.nums[i]
	}
	s.nums = newNums
}
