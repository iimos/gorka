package graph

import (
	"errors"
)

// NewRegular builds a graph where each node has k neighbors
func NewRegular(n, k int) (Graph, error) {
	if n <= k {
		if n == 0 && k == 0 {
			return New(), nil
		}
		return nil, errors.New("k should be less than n")
	}
	if k < 0 {
		return nil, errors.New("k should be positive")
	}

	g := newGraph()

	// 1. make n nodes
	for i := n; i > 0; i-- {
		g.NewNode("")
	}

	// 2. Connect each node to its k/2 nearest neighbors on ether side
	var v1, v2 Node
	r := k / 2
	odd := k&1 == 1

	for i := n - 1; i >= 0; i-- {
		v1 = g.nodes[i]

		for j := r; j > 0; j-- {
			v2 = g.nodes[(n+i+j)%n]
			g.AddEdge(v1, v2, 0)

			v2 = g.nodes[(n+i-j)%n]
			g.AddEdge(v1, v2, 0)
		}

		// if k is odd then connect nodes with the directly opposite node
		if odd {
			v2 = g.nodes[(i+n/2)%n]
			g.AddEdge(v1, v2, 0)
		}
	}

	return g, nil
}
