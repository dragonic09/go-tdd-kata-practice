package node

import (
	"testing"
)


func TestNodeShoulReceiveNextNodeWhenNextNodeExist(t *testing.T) {
	var node Node = Node{Value: 1}
	var nextNode Node = Node{Value: 2}
	node.Next = &nextNode
	if node.Next != &nextNode {
		t.Errorf("node.next should return the next node when a next node exist.")
	}
}