package gorka

import (
	"github.com/iimos/gorka/generic/queue"
)

// NodeQueue is a thread safe FIFO queue
type NodeQueue struct {
	q queue.Queue
}

// Len returns number of nodes in the queue
func (q *NodeQueue) Len() int {
	return q.q.Len()
}

// Push adds node to the end of the queue. Nils are ignored.
func (q *NodeQueue) Push(n Node) {
	q.q.Push(n)
}

// Pop returns the first node and removes it from the queue. Returns nil only if empty.
func (q *NodeQueue) Pop() Node {
	item := q.q.Pop()
	if item == nil {
		return nil
	}
	return item.(Node)
}

// NodeStack is a thread safe LIFO stack
type NodeStack struct {
	q queue.Stack
}

// Len returns number of nodes in the stack
func (q *NodeStack) Len() int {
	return q.q.Len()
}

// Push adds node to the end of the queue. Nils are ignored.
func (q *NodeStack) Push(n Node) {
	q.q.Push(n)
}

// Pop returns the first node and removes it from the queue. Returns nil only if empty.
func (q *NodeStack) Pop() Node {
	item := q.q.Pop()
	if item == nil {
		return nil
	}
	return item.(Node)
}
