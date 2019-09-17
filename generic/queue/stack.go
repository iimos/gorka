package queue

import (
	"sync"
)

// Stack is a thread safe LIFO stack
type Stack struct {
	items []interface{}
	lock  sync.RWMutex
}

// Len returns number of nodes on the stack
func (q *Stack) Len() int {
	q.lock.RLock()
	defer q.lock.RUnlock()

	return len(q.items)
}

// Push adds node to the stack. Nils are ignored.
func (q *Stack) Push(n interface{}) int {
	if n == nil {
		return len(q.items)
	}
	q.lock.Lock()
	defer q.lock.Unlock()

	q.items = append(q.items, n)
	return len(q.items)
}

// Pop returns the last node and removes it from the stack. Returns nil only if empty.
func (q *Stack) Pop() interface{} {
	q.lock.Lock()
	defer q.lock.Unlock()

	i := len(q.items) - 1
	if i < 0 {
		return nil
	}

	last := q.items[i]
	q.items = q.items[:i]
	return last
}
