package datastructures

import "fmt"

type ValueTypes interface {
	int | uint | float32
}

type BNode[T ValueTypes] struct {
	Value int
	Right *BNode[T]
	Left	*BNode[T]
}

type BinarySearchTree[T ValueTypes] struct {
	Root	*BNode[T]
}

func (b *BinarySearchTree[T]) Insert(value int) bool {
	newNode := &BNode[T]{Value: value}
	if b.Root == nil {
		b.Root = newNode
		return true
	}

	current := b.Root
	for  {
		if current.Value == value {
			return false
		}

		if value < current.Value {
			if current.Left == nil {
				current.Left = newNode
				return true
			}
			current = current.Left
		} else {
			if current.Right == nil {
				current.Right = newNode
				return true
			}
			current = current.Right
		}
	}
}

func (b *BinarySearchTree[T]) Contains(value int) bool {
	if b.Root == nil {
		return false
	}

	if b.Root.Value == value {
		return true
	}
	
	for current := b.Root; current != nil; {
		if value < current.Value {
			current = current.Left
		} else if value > current.Value {
			current = current.Right
		} else {
			return true
		}
	}
	return false
}

func (b *BinarySearchTree[T]) InOrderTraversal(node BNode[T]) {
	if &node.Value != nil {
		b.InOrderTraversal(*node.Left)
		fmt.Println(node.Value)
		b.InOrderTraversal(*node.Right)
	}
}
