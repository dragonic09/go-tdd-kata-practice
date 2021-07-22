package queue

import "testing"


func TestEnQueuedAddValueToEmptyQueue(t *testing.T) {
	queue := Queue{}
	queue.Enqueue(10)
	if queue.headNode == nil {
		t.Errorf("An Empty queue should have a head node after enqueue")
	}

	if queue.headNode.Value != 10 {
		t.Errorf("Added node should have the same value as Enqueue() input")
	}
}

func TestEnQueueAddValuesToQueue(t *testing.T) {
	queue := Queue{}
	queue.Enqueue(10)
	queue.Enqueue(20)
	queue.Enqueue(30)
	if queue.headNode == nil || queue.headNode.Next == nil {
		t.Errorf("Queue should have 2 elements after enqueue to times")
	}

	if queue.headNode.Value != 10 || queue.headNode.Next.Value != 20 {
		t.Errorf("Added node should have the same value as Enqueue() input")
	}
}

func TestPeekOnEmptyQueue(t *testing.T) {
	queue := Queue{}
	value := queue.Peek()
	if value != nil {
		t.Errorf("Using Peek() on an empty queue should return nil")
	}
}

func TestPeekOnQueue(t *testing.T) {
	queue := Queue{}
	queue.Enqueue(10)
	queue.Enqueue(20)
	value := queue.Peek()
	
	if value != 10 {
		t.Errorf("Peek() should return the first value from a queue.")
	}

	if queue.headNode == nil || queue.headNode.Next == nil {
		t.Errorf("Queue should have the same number of nodes after Peek().")
	}

	if queue.headNode.Value != 10 || queue.headNode.Next.Value != 20 {
		t.Errorf("Queue should have the same values after Peek(). ")
	}
}

func TestIsEmptyOnEmptyQueue(t *testing.T) {
	queue := Queue{}
	isEmpty := queue.IsEmpty()
	if !isEmpty {
		t.Errorf("Using IsEmpty() on an empty queue should return true.")
	}
}

func TestIsEmptyOnQueue(t *testing.T) {
	queue := Queue{}
	queue.Enqueue(10)
	queue.Enqueue(20)
	isEmpty := queue.IsEmpty()
	
	if isEmpty {
		t.Errorf("IsEmpty() should return true when a queue has values.")
	}

	if queue.headNode == nil || queue.headNode.Next == nil {
		t.Errorf("Queue should have the same number of nodes after IsEmpty().")
	}

	if queue.headNode.Value != 10 || queue.headNode.Next.Value != 20 {
		t.Errorf("Queue should have the same values after IsEmpty(). ")
	}
}

func TestDequeueOnEmptyQueue(t *testing.T) {
	queue := Queue{}
	_, err  := queue.Dequeue()
	if err == nil {
		t.Errorf("Dequeue() on an empty queue should return an error.")
	}
}

func TestDequeueOnQueueShouldReturnTheFirstValue(t *testing.T) {
	queue := Queue{}
	queue.Enqueue(10)
	queue.Enqueue(20)
	queue.Enqueue(30)
	value, err := queue.Dequeue()
	if err != nil {
		t.Errorf("Dequeue() on a non-empty queue should return error as nil.")
	}

	if value != 10 {
		t.Errorf("Dequeue() on a non-empty queue should return the first value.")
	}
}

func TestDequeueOnQueueShouldRemoveFirstValueAndMoveNextValueToFirstValue(t *testing.T) {
	queue := Queue{}
	queue.Enqueue(10)
	queue.Enqueue(20)
	queue.Enqueue(30)
	queue.Dequeue()
	if queue.headNode.Value != 20 {
		t.Errorf("Queue should pop 10 out of the queue and have 20 as the first value after Dequeue().")
	}
}


func TestDequeueOnQueueWithSingleValueQueueShouldBeEmpty(t *testing.T) {
	queue := Queue{}
	queue.Enqueue(10)
	queue.Dequeue()
	if queue.headNode != nil {
		t.Errorf("Queue should be empty if Dequeque() is called on a queue with a single value.")
	}
}

