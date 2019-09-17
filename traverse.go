package graph

import (
	"container/heap"
	"errors"
)

type pushpoper interface {
	Push(n Node)
	Pop() Node
	Len() int
}

type iterator struct {
	graph   Graph
	queue   pushpoper
	visited []bool
	reverse bool
}

func (iter *iterator) isVisited(n Node) bool {
	return iter.visited[n.ID()]
}

func (iter *iterator) next() (Node, error) {
	n := iter.queue.Pop()
	if n == nil {
		return nil, nil
	}

	if iter.isVisited(n) {
		return iter.next()
	}

	iter.visited[n.ID()] = true

	iter.graph.NodeEdgeIter(n, func(e Edge) bool {
		d := e.Dst()
		if !iter.isVisited(d) {
			iter.queue.Push(d)
		}
		return true
	})

	return n, nil
}

func newIterator(q pushpoper, g Graph, n Node) *iterator {
	iter := iterator{
		graph:   g,
		queue:   q,
		visited: make([]bool, g.MaxNodeID()+1),
	}
	iter.queue.Push(n)
	return &iter
}

func traverse(iter *iterator, fn Callback) error {
	for {
		n, err := iter.next()
		if err != nil {
			return err
		}
		if n == nil {
			return nil
		}

		further := fn(n)
		if !further {
			return nil
		}
	}
}

// Callback is a function that called for each node we visit
type Callback func(n Node) (further bool)

// TraverseBreadthFirst goes throught the graph from the start node and calls fn for each node
func TraverseBreadthFirst(g Graph, start Node, fn Callback) error {
	q := &NodeQueue{}
	iter := newIterator(q, g, start)
	return traverse(iter, fn)
}

// TraverseDepthFirst goes throught the graph from the start node and calls fn for each node
func TraverseDepthFirst(g Graph, start Node, fn Callback) error {
	s := &NodeStack{}
	iter := newIterator(s, g, start)
	iter.reverse = true
	return traverse(iter, fn)
}

// ErrPathNotFound means that path not found
var ErrPathNotFound = errors.New("path not found")

// ShortestPath implements Dijkstra's shortest path algorithm
func ShortestPath(g Graph, a, b Node) (path []Edge, len float32, err error) {
	if a.ID() == b.ID() {
		return []Edge{}, 0, nil
	}

	from := map[int]Edge{}
	dist := map[int]float32{}
	q := &nodeHeap{{node: a, dist: 0}}

	for q.Len() > 0 {
		d := heap.Pop(q).(distance)
		g.NodeEdgeIter(d.node, func(e Edge) bool {
			n := e.Dst()
			w := e.Wieght()
			if w < 0 {
				err = errors.New("Dijkstra's shortest path algorithm doesn't support negative weights")
				return false
			}
			alt := d.dist + w
			curr, visited := dist[n.ID()]
			if !visited || alt < curr {
				heap.Push(q, distance{node: n, dist: alt})
				dist[n.ID()] = alt
				from[n.ID()] = e
			}
			return true
		})
		if err != nil {
			return nil, 0, err
		}
	}

	aid := b.ID()
	bid := b.ID()
	d, ok := dist[bid]
	if !ok {
		return nil, 0, ErrPathNotFound
	}

	path = make([]Edge, 0, 8)
	e, _ := from[bid]
	for e != nil && e.From().ID() != aid {
		path = append(path, e)
		e = from[e.From().ID()]
	}
	reversePath(path)
	return path, d, nil
}

func reversePath(path []Edge) {
	l := len(path)
	for i := l/2 - 1; i >= 0; i-- {
		opp := l - 1 - i
		path[i], path[opp] = path[opp], path[i]
	}
}

// distance to the node
type distance struct {
	node Node
	dist float32
}

// nodeHeap is a min-heap for storing distances to the node
type nodeHeap []distance

func (h nodeHeap) Len() int            { return len(h) }
func (h nodeHeap) Less(i, j int) bool  { return h[i].dist < h[j].dist }
func (h nodeHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *nodeHeap) Push(x interface{}) { *h = append(*h, x.(distance)) }

func (h *nodeHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
