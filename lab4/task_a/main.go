package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type BinarySearchTree interface {
	Insert(key int32)
	Search(key int32) bool
	InOrder() []int32
}

var _ BinarySearchTree = &BST{}

type BSTNode struct {
	key   int32
	left  *BSTNode
	right *BSTNode
}

type BST struct {
	root *BSTNode
}

func NewBST() *BST {
	return &BST{}
}

func (b *BST) InOrder() []int32 {
	result := []int32{}

	var walk func(node *BSTNode)
	walk = func(node *BSTNode) {
		if node == nil {
			return
		}
		walk(node.left)
		result = append(result, node.key)
		walk(node.right)
	}

	walk(b.root)
	return result
}

func (b *BST) Insert(key int32) {
	if b.root == nil {
		b.root = &BSTNode{key: key}
	} else {
		var parent *BSTNode
		node := b.root
		for node != nil {
			parent = node
			switch {
			case key == node.key:
				return
			case key < node.key:
				node = node.left
			default:
				node = node.right
			}
		}
		if key < parent.key {
			parent.left = &BSTNode{key: key}
		} else {
			parent.right = &BSTNode{key: key}
		}
	}
}

func (b *BST) Search(key int32) bool {
	for node := b.root; node != nil; {
		switch {
		case key == node.key:
			return true
		case key > node.key:
			node = node.right
		case key < node.key:
			node = node.left
		}
	}
	return false
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n uint32
	if _, err := fmt.Fscan(reader, &n); err != nil {
		return
	}

	a := make([]int32, n)
	for i := uint32(0); i < n; i++ {
		if _, err := fmt.Fscan(reader, &a[i]); err != nil {
			return
		}
	}

	bst := NewBST()

	for i := 0; i < int(n); i++ {
		bst.Insert(a[i])
	}

	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	for i, key := range bst.InOrder() {
		if i > 0 {
			out.WriteByte(' ')
		}
		out.WriteString(strconv.FormatInt(int64(key), 10))
	}
	out.WriteByte('\n')
}
