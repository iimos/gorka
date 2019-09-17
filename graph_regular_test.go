package graph

import "testing"

func TestRegularGraphCreation(t *testing.T) {
	table := [][]int{
		[]int{5, 3},
		[]int{3, 1},
		[]int{4, 3},
		[]int{4, 2},
		[]int{2, 0},
		[]int{2, 1},
		[]int{0, 0},
	}

	for i, nk := range table {
		n, k := nk[0], nk[1]
		g, err := NewRegular(n, k)
		if err != nil {
			t.Errorf("#%d: error: %s", i, err)
		}

		nc := g.NodesCount()
		if nc != n {
			t.Errorf("#%d: wrong nodes count - %d, expected %d", i, nc, n)
		}

		ec := g.EdgesCount()
		ecExpected := n * k
		if ec != ecExpected {
			t.Errorf("#%d: wrong edges count - %d, expected %d", i, ec, ecExpected)
		}

		g.NodeIter(func(n Node) bool {
			d := g.OutDegree(n)
			if d != k {
				t.Errorf("#%d: node %s has wrong degree - %d", i, n, d)
			}
			return true
		})
	}
}

func TestRegularGraphWithWrongParams(t *testing.T) {
	table := [][]int{
		[]int{5, 5},
		[]int{1, 10},
		[]int{-5, 5},
		[]int{5, -1},
		[]int{0, 1},
	}

	for _, pair := range table {
		n, k := pair[0], pair[1]
		_, err := NewRegular(n, k)
		if err == nil {
			t.Errorf("NewRegular(%d, %d) returns no error", n, k)
			return
		}
	}
}
