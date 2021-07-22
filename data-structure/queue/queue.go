package queue

import (
	"data-structure-kata/node"
	"errors"
)

type IQueue interface {
	Enqueue(value interface{})
	Dequeue() (interface{}, error)
	Peek() interface{}
	IsEmpty() bool
}

type Queue struct {
	headNode *node.Node
}

func (queue *Queue) Enqueue(value interface{}) {
	currentNode := queue.headNode
	for {
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

func (queue *Queue) Dequeue() (interface{}, error) {
	if queue.headNode == nil {
		return nil, errors.New("Queue is empty")
	}
	firstValue := queue.headNode.Value
	queue.headNode = queue.headNode.Next

	return firstValue, nil
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