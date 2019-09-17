package alg

import (
	"math"
	"testing"
	"xxx/complexity/graph"
	"xxx/complexity/graph/gralang"
)

const floatEqualityThreshold = 1e-4

func almostEqual(a, b float32) bool {
	return math.Abs(float64(a-b)) <= floatEqualityThreshold
}

func TestClusteringCoefLocal(t *testing.T) {
	type tcase struct {
		s string
		c float32
	}
	cases := [...]tcase{
		tcase{"x", 0},
		tcase{"x -> a b c", 0},
		tcase{"x -> a; a -> b; b -> a", 0},
		tcase{"a b c -> x", 0},
		tcase{"x -> a b; a -> b", 1},
		tcase{"x -> a b c; a -> b", 1.0 / 3.0},
	}

	for i, c := range cases {
		g := graph.New()
		err := gralang.Parse(g, c.s)
		if err != nil {
			t.Errorf("#%d. Parse error: %s", i, err)
			continue
		}
		x, ok := g.NodeByLabel("x")
		if !ok {
			t.Errorf("#%d. Node x is absent", i)
			continue
		}
		res := ClusteringCoefLocal(g, x)
		if res != c.c {
			t.Errorf("#%d. Wrong coef: got %f, expected %f", i, res, c.c)
		}
	}
}

func TestClusteringCoef(t *testing.T) {
	type tcase struct {
		s string
		c float32
	}
	cases := [...]tcase{
		tcase{"x", 0},
		tcase{"x -> a b c", 0},
		tcase{"x -> a; a -> b; b -> a", 0},
		tcase{"a b c -> x", 0},
		tcase{"a -- b c; b c -- d; b -- c", 10.0 / 12.0},
		tcase{"a -- b; b -- c; c -- a", 1},
	}

	for i, c := range cases {
		g := graph.New()
		err := gralang.Parse(g, c.s)
		if err != nil {
			t.Errorf("#%d. Parse error: %s", i, err)
			continue
		}
		res := ClusteringCoef(g)
		if !almostEqual(res, c.c) {
			t.Errorf("#%d. Wrong coef: got %f, expected %f", i, res, c.c)
		}
	}
}
