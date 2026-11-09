package main

import (
	"bufio"
	"os"
	"strconv"
)

type BinarySearchTree interface {
	Insert(key int64)
	Search(key int64) bool
	InOrder() []int64
	RightView(n int) []int64
}

var _ BinarySearchTree = &BST{}

type BSTNode struct {
	key   int64
	left  *BSTNode
	right *BSTNode
}

type BST struct {
	root *BSTNode
}

func NewBST() *BST {
	return &BST{}
}

func (b *BST) InOrder() []int64 {
	result := []int64{}

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

func (b *BST) Insert(key int64) {
	if b.root == nil {
		b.root = &BSTNode{key: key}
		return
	}

	var parent *BSTNode
	node := b.root
	for node != nil {
		parent = node
		if key == node.key {
			return
		}
		if key < node.key {
			node = node.left
		} else {
			node = node.right
		}
	}
	if key < parent.key {
		parent.left = &BSTNode{key: key}
	} else {
		parent.right = &BSTNode{key: key}
	}
}

func (b *BST) Search(key int64) bool {
	for node := b.root; node != nil; {
		switch {
		case key == node.key:
			return true
		case key > node.key:
			node = node.right
		default:
			node = node.left
		}
	}
	return false
}

func (b *BST) RightView(n int) []int64 {
	result := []int64{}
	if b.root == nil {
		return result
	}
	queue := make([]*BSTNode, 0, n)
	queue = append(queue, b.root)
	head := 0

	for head < len(queue) {
		levelEnd := len(queue)
		for ; head < levelEnd; head++ {
			node := queue[head]
			if head == levelEnd-1 {
				result = append(result, node.key)
			}

			if node.left != nil {
				queue = append(queue, node.left)
			}
			if node.right != nil {
				queue = append(queue, node.right)
			}
		}
	}

	return result
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1024*1024), 1024*1024)
	scanner.Split(bufio.ScanWords)

	if !scanner.Scan() {
		return
	}
	n, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return
	}

	bst := NewBST()

	for i := 0; i < n && scanner.Scan(); i++ {
		key, err := strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			continue
		}
		bst.Insert(key)
	}

	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	for i, key := range bst.RightView(n) {
		if i > 0 {
			out.WriteByte(' ')
		}
		out.WriteString(strconv.FormatInt(key, 10))
	}
	out.WriteByte('\n')
}
