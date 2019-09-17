package graph

import (
	"testing"
)

func TestRandomGraph(t *testing.T) {
	table := [][]float32{
		// n, p, minEdges, maxEdges
		[]float32{10, 0.0, 0, 0},
		[]float32{10, 1.0, 90, 90},
		[]float32{5, 0.99, 18, 20},
		[]float32{10, 0.5, 30, 60},
		[]float32{0, 0.5, 0, 0},
	}

	for i, test := range table {
		n, p, minEdges, maxEdges := int(test[0]), test[1], int(test[2]), int(test[3])

		g, err := NewRandom(n, p)
		if err != nil {
			t.Errorf("#%d: can't make graph: %s", i, err)
			return
		}

		if g.NodesCount() != n {
			t.Errorf("#%d: wrong nodes count: %d", i, g.NodesCount())
		}

		edges := g.EdgesCount()
		if edges < minEdges || edges > maxEdges {
			t.Errorf("#%d: wrong edges count: %d not in [%d, %d]", i, edges, minEdges, maxEdges)
		}
	}
}

func TestRandomGraphBidir(t *testing.T) {
	table := [][]float32{
		// n, p, minEdges, maxEdges
		[]float32{10, 0.0, 0, 0},
		[]float32{10, 1.0, 90, 90},
		[]float32{5, 0.99, 18, 20},
		[]float32{10, 0.5, 30, 60},
		[]float32{0, 0.5, 0, 0},
	}

	for i, test := range table {
		n, p, minEdges, maxEdges := int(test[0]), test[1], int(test[2]), int(test[3])

		g, err := NewRandom(n, p)
		if err != nil {
			t.Errorf("#%d: can't make graph: %s", i, err)
			return
		}

		if g.NodesCount() != n {
			t.Errorf("#%d: wrong nodes count: %d", i, g.NodesCount())
		}

		edges := g.EdgesCount()
		if edges < minEdges || edges > maxEdges {
			t.Errorf("#%d: wrong edges count: %d not in [%d, %d]", i, edges, minEdges, maxEdges)
		}
	}
}

func BenchmarkRandomGraph(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewRandom(1000, 0.5)
	}
}
