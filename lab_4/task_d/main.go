package main

import (
	"bufio"
	"fmt"
	"os"
)

type HashTable interface {
	Mark(key string, student int)
	Get(key string) uint8
}

var _ HashTable = &ChainedHashTable{}

type entry struct {
	key  string
	mask uint8
	next *entry
}

type ChainedHashTable struct {
	buckets []*entry
	size    int
}

func NewChainedHashTable(capacity int) *ChainedHashTable {
	return &ChainedHashTable{
		buckets: make([]*entry, capacity),
	}
}

func (h *ChainedHashTable) hash(key string) int {
	x := uint64(0)
	for i := 0; i < len(key); i++ {
		c := key[i]
		x = x*31 + uint64(c)
	}
	return int(x % uint64(len(h.buckets)))
}

func (h *ChainedHashTable) find(key string) *entry {
	for node := h.buckets[h.hash(key)]; node != nil; node = node.next {
		if node.key == key {
			return node
		}
	}
	return nil
}

func (h *ChainedHashTable) Mark(key string, student int) {
	e := h.find(key)
	if e == nil {
		idx := h.hash(key)
		e = &entry{key: key, next: h.buckets[idx]}
		h.buckets[idx] = e
		h.size++
	}
	e.mask |= 1 << student
}

func (h *ChainedHashTable) Get(key string) uint8 {
	e := h.find(key)
	if e != nil {
		return e.mask
	}
	return 0
}

func popcount(mask uint8) int {
	count := 0
	for s := 0; s < 3; s++ {
		if mask&(1<<s) != 0 {
			count++
		}
	}
	return count
}

func score(owners int) int {
	switch owners {
	case 1:
		return 3
	case 2:
		return 1
	default:
		return 0
	}
}

func solve(files [3][]string) [3]int {
	var points [3]int

	table := NewChainedHashTable(60013)
	for s := 0; s < 3; s++ {
		for _, f := range files[s] {
			table.Mark(f, s)
		}
	}
	for s := 0; s < 3; s++ {
		for _, f := range files[s] {
			points[s] += score(popcount(table.Get(f)))
		}
	}

	return points
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	var files [3][]string
	for s := 0; s < 3; s++ {
		files[s] = make([]string, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(reader, &files[s][i])
		}
	}

	points := solve(files)
	fmt.Fprintln(writer, points[0], points[1], points[2])
}
