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

	queue := NewLinkedQueue()

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
			queue.Enqueue(x)
		case OperationPop:
			if val, ok := queue.Dequeue(); ok {
				fmt.Fprintln(writer, val)
			}
		}
	}
}

type LinkedQueue struct {
	front *Node
	rear  *Node
	size  int
}

func NewLinkedQueue() *LinkedQueue {
	return &LinkedQueue{
		front: nil,
		rear:  nil,
		size:  0,
	}
}

type Node struct {
	value int32
	next  *Node
}

func (q *LinkedQueue) IsEmpty() bool {
	return q.size == 0
}

func (q *LinkedQueue) Size() int {
	return q.size
}

func (q *LinkedQueue) Enqueue(val int32) {
	newNode := &Node{value: val, next: nil}

	if q.IsEmpty() {
		q.front = newNode
		q.rear = newNode
	} else {
		q.rear.next = newNode
		q.rear = newNode
	}
	q.size++
}

func (q *LinkedQueue) Dequeue() (int32, bool) {
	if q.IsEmpty() {
		return 0, false
	}

	val := q.front.value
	q.front = q.front.next
	q.size--

	if q.front == nil {
		q.rear = nil
	}

	return val, true
}

func (q *LinkedQueue) Peek() (int32, bool) {
	if q.IsEmpty() {
		return 0, false
	}
	return q.front.value, true
}
