package main

import (
	"bufio"
	"fmt"
	"math/big"
	"math/bits"
	"os"
)

type Heap[T any] struct {
	data []T
	less func(a, b T) bool
}

func NewHeap[T any](less func(a, b T) bool) *Heap[T] {
	return &Heap[T]{less: less}
}

func (h *Heap[T]) Len() int {
	return len(h.data)
}

func (h *Heap[T]) swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}

func (h *Heap[T]) siftUp(i int) {
	for i > 0 {
		parent := (i - 1) / 2
		if !h.less(h.data[i], h.data[parent]) {
			break
		}
		h.swap(i, parent)
		i = parent
	}
}

func (h *Heap[T]) siftDown(i int) {
	n := h.Len()
	for {
		left, right := 2*i+1, 2*i+2
		smallest := i
		if left < n && h.less(h.data[left], h.data[smallest]) {
			smallest = left
		}
		if right < n && h.less(h.data[right], h.data[smallest]) {
			smallest = right
		}
		if smallest == i {
			return
		}
		h.swap(i, smallest)
		i = smallest
	}
}

func (h *Heap[T]) Push(x T) {
	h.data = append(h.data, x)
	h.siftUp(h.Len() - 1)
}

func (h *Heap[T]) Pop() T {
	root := h.data[0]
	last := h.Len() - 1
	h.data[0] = h.data[last]
	var zero T
	h.data[last] = zero
	h.data = h.data[:last]
	if last > 0 {
		h.siftDown(0)
	}
	return root
}

func (h *Heap[T]) Peek() T {
	return h.data[0]
}

type Job struct {
	finish int
	salary int
}

func readInt(scanner *bufio.Scanner) int {
	scanner.Scan()
	n := 0
	for _, c := range scanner.Bytes() {
		n = n*10 + int(c-'0')
	}
	return n
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1024*1024), 1024*1024)
	scanner.Split(bufio.ScanWords)

	n, m := readInt(scanner), readInt(scanner)

	free := NewHeap(func(a, b int) bool { return a < b })
	busy := NewHeap(func(a, b Job) bool { return a.finish < b.finish })

	for i := 0; i < n; i++ {
		free.Push(readInt(scanner))
	}

	var hi, lo uint64

	for i := 0; i < m; i++ {
		t, f := readInt(scanner), readInt(scanner)

		for busy.Len() > 0 && busy.Peek().finish <= t {
			free.Push(busy.Pop().salary)
		}

		if free.Len() == 0 {
			continue
		}

		salary := free.Pop()
		pHi, pLo := bits.Mul64(uint64(salary), uint64(f))
		var carry uint64
		lo, carry = bits.Add64(lo, pLo, 0)
		hi += pHi + carry

		busy.Push(Job{finish: t + f, salary: salary})
	}

	total := new(big.Int).SetUint64(hi)
	total.Lsh(total, 64)
	total.Or(total, new(big.Int).SetUint64(lo))
	fmt.Println(total)
}
