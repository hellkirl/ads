package main

import (
	"bufio"
	"os"
	"strconv"
)

type PriorityQueue interface {
	Insert(x int)
	ExtractMin() (int, bool)
	DecreaseKey(x, y int)
	Len() int
}

var _ PriorityQueue = &MinHeap{}

type MinHeap struct {
	data []int
}

func NewMinHeap() *MinHeap {
	return &MinHeap{
		data: []int{},
	}
}

func MergeHeaps(a, b *MinHeap) *MinHeap {
	data := make([]int, len(a.data)+len(b.data))
	copy(data, a.data)
	copy(data[len(a.data):], b.data)

	h := &MinHeap{data: data}
	for i := len(h.data)/2 - 1; i >= 0; i-- {
		h.siftDown(i)
	}
	return h
}

func parent(i int) int {
	return (i - 1) / 2
}

func left(i int) int {
	return 2*i + 1
}

func right(i int) int {
	return 2*i + 2
}

func (h *MinHeap) swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}

func (h *MinHeap) siftUp(i int) {
	for i > 0 && h.data[i] < h.data[parent(i)] {
		h.swap(i, parent(i))
		i = parent(i)
	}
}

func (h *MinHeap) siftDown(i int) {
	n := h.Len()
	for {
		l := left(i)
		if l >= n {
			return
		}

		smallest := l
		if r := right(i); r < n && h.data[r] < h.data[l] {
			smallest = r
		}

		if h.data[i] <= h.data[smallest] {
			return
		}

		h.swap(i, smallest)
		i = smallest
	}
}

func (h *MinHeap) Insert(x int) {
	h.data = append(h.data, x)
	h.siftUp(len(h.data) - 1)
}

func (h *MinHeap) ExtractMin() (int, bool) {
	if h.Len() == 0 {
		return 0, false
	}

	min := h.data[0]
	last := h.Len() - 1
	h.data[0] = h.data[last]
	h.data = h.data[:last]
	if h.Len() > 0 {
		h.siftDown(0)
	}
	return min, true
}

func (h *MinHeap) DecreaseKey(x, y int) {
	for i, v := range h.data {
		if v == x {
			h.data[i] = y
			h.siftUp(i)
			return
		}
	}
}

func (h *MinHeap) Len() int {
	return len(h.data)
}

func main() {
	in := bufio.NewScanner(os.Stdin)
	in.Buffer(make([]byte, 1024*1024), 1024*1024)
	in.Split(bufio.ScanWords)

	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	readInt := func() int {
		in.Scan()
		n, _ := strconv.Atoi(in.Text())
		return n
	}

	queues := []*MinHeap{}

	for in.Scan() {
		switch in.Text() {
		case "create":
			queues = append(queues, NewMinHeap())
		case "insert":
			k := readInt()
			x := readInt()
			queues[k].Insert(x)
		case "extract-min":
			k := readInt()
			if v, ok := queues[k].ExtractMin(); ok {
				out.WriteString(strconv.Itoa(v))
				out.WriteByte('\n')
			} else {
				out.WriteString("*\n")
			}
		case "merge":
			k := readInt()
			m := readInt()
			queues = append(queues, MergeHeaps(queues[k], queues[m]))
		case "decrease-key":
			k := readInt()
			x := readInt()
			y := readInt()
			queues[k].DecreaseKey(x, y)
		}
	}
}
