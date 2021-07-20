package list

import (
	"testing"
)

func TestPushValueToList(t *testing.T) {
	var list LinkedList = LinkedList{}
	list.Push(1)
	
	if list.Length != 1 || list.headNode.Value != 1 {
		t.Errorf("List should have length equal to 1 and headNode's value should equal input of Push().")
	}
}

func TestPushMultipleValueToList(t *testing.T) {
	var list LinkedList = LinkedList{}
	list.Push(1)
	list.Push(2)
	if list.Length != 2 || list.headNode.Value != 1 || list.headNode.Next.Value != 2 {
		t.Errorf("List should have length equal to 2 and nodes' value should equal input of Push().")
	}
}

func TestGetValue(t *testing.T) {
	var list LinkedList = LinkedList{}
	list.Push(10)
	list.Push(20)
	result, error := list.Get(1)
	if result != 20 || error != nil {
		t.Errorf("Should get a correct value at index 1.")
	}
}

func TestGetValueWithIndexOutOfBound(t *testing.T) {
	var list LinkedList = LinkedList{}
	list.Push(10)
	list.Push(20)
	_, error := list.Get(3)
	if error == nil {
		t.Errorf("Should return error when target index is out of list's index boundary")
	}
}

func TestRemoveOnEmptyList(t *testing.T) {
	var list LinkedList = LinkedList{}
	error := list.Remove(0)
	if error == nil {
		t.Errorf("Should return Index Out of Bound error")
	}
}

func TestRemoveWithOutsideOfBoundIndex(t *testing.T) {
	var list LinkedList = LinkedList{}
	list.Push(10)
	list.Push(20)
	error := list.Remove(3)
	if error == nil {
		t.Errorf("Should return Index Out of Bound error")
	}
}

func TestRemoveValueSuccessfully(t *testing.T) {
	var list LinkedList = LinkedList{}
	list.Push(10)
	list.Push(20)
	list.Push(30)
	error := list.Remove(1)

	if error != nil {
		t.Errorf("Should remove a value successfully.")
	}
	
	if list.Length != 2 {
		t.Errorf("Remove a value should reduce the target list's length by 1.")
	}

	if list.headNode.Next.Value != 30 {
		t.Errorf("Should connect a node at index 0 with a node at index 2.")
	}
}