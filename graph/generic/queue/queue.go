package queue

import (
	"sync"
)

// Queue is a thread safe LIFO queue
type Queue struct {
	items []interface{}
	lock  sync.RWMutex
}

// Len returns number of items in the queue
func (q *Queue) Len() int {
	q.lock.RLock()
	defer q.lock.RUnlock()

	return len(q.items)
}

// Push adds item to the end of the queue. Nils are ignored.
func (q *Queue) Push(n interface{}) int {
	if n == nil {
		return len(q.items)
	}

	q.lock.Lock()
	defer q.lock.Unlock()

	q.items = append(q.items, n)
	return len(q.items)
}

// Pop returns the first item and removes it from the queue. Returns nil only if empty.
func (q *Queue) Pop() interface{} {
	q.lock.Lock()
	defer q.lock.Unlock()

	if len(q.items) == 0 {
		return nil
	}

	first := q.items[0]
	q.items = q.items[1:]
	return first
}
