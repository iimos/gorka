package graph

import (
	"errors"
	"math/rand"
	"time"
)

// NewRandom builds an directed Erdos-Renyi random graph - n-vertices graph each
// vertices pair of which connected by an edge with probability p
func NewRandom(n int, p float32) (Graph, error) {
	if p < 0 || p > 1 {
		return nil, errors.New("p should be in [0, 1] interval")
	}

	g := newGraph()

	for i := n; i > 0; i-- {
		g.NewNode("")
	}

	src := rand.NewSource(time.Now().UnixNano())
	rnd := rand.New(src)

	for i := n - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if i != j && rnd.Float32() < p {
				a, b := g.nodes[i], g.nodes[j]
				g.AddEdge(a, b, 0)
			}
		}
	}

	return g, nil
}

// NewRandomBidir builds an bidirected Erdos-Renyi random graph - n-vertices graph each
// vertices pair of which connected by an edge with probability p
func NewRandomBidir(n int, p float32) (Graph, error) {
	if p < 0 || p > 1 {
		return nil, errors.New("p should be in [0, 1] interval")
	}

	g := newGraph()

	for i := n; i > 0; i-- {
		g.NewNode("")
	}

	src := rand.NewSource(time.Now().UnixNano())
	rnd := rand.New(src)

	for i := n - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if i != j && rnd.Float32() < p {
				a, b := g.nodes[i], g.nodes[j]
				g.AddBiEdge(a, b, 0)
			}
		}
	}

	return g, nil
}
