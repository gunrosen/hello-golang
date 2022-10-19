package main

import (
	"bufio"
	"fmt"
	"os"
)

// https://codeforces.com/problemset/problem/706/B

type Tree struct {
	root *Node
}

type Node struct {
	key   int64
	left  *Node
	right *Node
}

// Tree
func (t *Tree) insert(data int64) {
	if t.root == nil {
		t.root = &Node{key: data}
	} else {
		t.root.insert(data)
	}
}

// Node
func (n *Node) insert(data int64) {
	if data <= n.key {
		if n.left == nil {
			n.left = &Node{key: data}
		} else {
			n.left.insert(data)
		}
	} else {
		if n.right == nil {
			n.right = &Node{key: data}
		} else {
			n.right.insert(data)
		}
	}
}

// find by In Order
func findNode(n *Node, value int64) int64 {
	var count = int64(0)
	if n == nil {
		return 0
	} else {
		if n.key <= value {
			count++
		}
		count = count + findNode(n.left, value)
		count = count + findNode(n.right, value)
	}
	return count
}

func main() {
	var numberOfDay, numberOfShop int64
	var tree Tree
	in := bufio.NewReader(os.Stdin)

	fmt.Fscan(in, &numberOfShop)
	for i := int64(0); i < numberOfShop; i++ {
		var temp int64
		fmt.Fscan(in, &temp)
		tree.insert(temp)
	}

	fmt.Fscan(in, &numberOfDay)
	for i := int64(0); i < numberOfDay; i++ {
		var temp int64
		fmt.Fscan(in, &temp)
		count := findNode(tree.root, temp)
		fmt.Println(count)
	}
}
