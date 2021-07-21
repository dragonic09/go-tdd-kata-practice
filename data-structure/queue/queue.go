package queue

import (
	"data-structure-kata/node"
)

type IQueue interface {
	Enqueue(value interface{})
	//Dequeue() (interface{}, error)
	Peek() interface{}
	IsEmpty() bool
}

type Queue struct {
	headNode *node.Node
}

func (queue *Queue) Enqueue(value interface{}) {
	for {
		currentNode := queue.headNode
		if currentNode == nil {
			queue.headNode = &node.Node{Value: value}
			break
		}

		if currentNode.Next == nil {
			currentNode.Next = &node.Node{Value: value}
			break
		}
		currentNode = currentNode.Next
	}
}

func (queue *Queue) Peek() interface{} {
	if queue.headNode == nil {
		return nil
	}
	return queue.headNode.Value
}

func (queue *Queue) IsEmpty() bool {
	if queue.headNode == nil {
		return true
	}
	return false
}