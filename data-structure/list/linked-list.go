package list

import (
	"data-structure-kata/node"
	"errors"
)

type ILinkedList interface {
	Get(index int) (interface{}, error)
	Remove(index int)
	Push(value interface{})
}

type LinkedList struct {
	headNode *node.Node
	Length int
}

func (list *LinkedList) Push(value interface{}) {
	if list.Length == 0 {
		list.headNode = &node.Node{Value: value}
	} else {
		currentNode := list.headNode
		for {
			if currentNode.Next == nil {
				currentNode.Next = &node.Node{Value: value}
				break
			}
			currentNode = currentNode.Next
		}
	}
	list.Length += 1
}

func (list *LinkedList) Get(targetIndex int) (interface{}, error) {
	if targetIndex > list.Length {
		return nil, errors.New("List index out of bound")
	}
	listIndex := 0
	currentNode := list.headNode
	for {
		if listIndex == targetIndex {
			break
		}
		currentNode = currentNode.Next
		listIndex++
	}
	return currentNode.Value, nil
}

func (list *LinkedList) Remove(targetIndex int) error {
	if list.Length - 1 < targetIndex {
		return errors.New("List index out of bound")
	}
	listIndex := 0
	var currentNode *node.Node = list.headNode
	var previousNode *node.Node = nil
	for {
		if listIndex == targetIndex - 1 {
			previousNode = currentNode
		}
		if listIndex == targetIndex {
			if currentNode.Next != nil { 
				previousNode.Next = currentNode.Next
			}
			list.Length -= 1
			break
		}
		currentNode = currentNode.Next
		listIndex++
	}
	return nil
}

