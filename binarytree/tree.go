package binarytree

import (
	"fmt"
	"sync"
)

type Node struct {
	Key   int
	Value string
	Left  *Node
	Right *Node
}

type BinaryTree struct {
	Root *Node
	lock sync.RWMutex
}

func (bst *BinaryTree) Insert(key int, value string) {
	bst.lock.Lock()
	defer bst.lock.Unlock()
	n := &Node{key, value, nil, nil}
	if bst.Root == nil {
		bst.Root = n
	} else {
		insertNode(bst.Root, n)
	}
}

func insertNode(node, newNode *Node) {
	if newNode.Key < node.Key {
		if node.Left == nil {
			node.Left = newNode
		} else {
			insertNode(node.Left, newNode)
		}
	} else {
		if node.Right == nil {
			node.Right = newNode
		} else {
			insertNode(node.Right, newNode)
		}
	}
}

func (bst *BinaryTree) InOrderTraverse(f func(string)) {
	bst.lock.RLock()
	defer bst.lock.RUnlock()
	inOrderTraverse(bst.Root, f)
}

func inOrderTraverse(n *Node, f func(string)) {
	if n != nil {
		inOrderTraverse(n.Left, f)
		f(n.Value)
		inOrderTraverse(n.Right, f)
	}
}

func (bst *BinaryTree) String() {
	bst.lock.Lock()
	defer bst.lock.Unlock()
	fmt.Println("-------------------------")
	stringify(bst.Root, 0)
	fmt.Println("-------------------------")
}

func stringify(n *Node, level int) {
	if n != nil {
		format := ""
		for i := 0; i < level; i++ {
			format += "     "
		}
		format += "---["
		level++
		stringify(n.Left, level)
		fmt.Printf(format+"%d\n", n.Key)
		stringify(n.Right, level)
	}
}
