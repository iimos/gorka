package gralang

import (
	"testing"
	"github.com/iimos/gorka"
)

func TestGralang(t *testing.T) {
	type e struct{ a, b string }
	type testcase struct {
		text      string
		nodeCount int
		edgeCount int
		edges     []e
	}

	cases := []testcase{
		testcase{"", 0, 0, []e{}},
		testcase{"a", 1, 0, []e{}},
		testcase{"a b", 2, 0, []e{}},
		testcase{"a -> b", 2, 1, []e{
			e{"a", "b"},
		}},
		testcase{"  a->b  ", 2, 1, []e{
			e{"a", "b"},
		}},
		testcase{"\n\na -> b\n\n\n", 2, 1, []e{
			e{"a", "b"},
		}},
		testcase{"a -- b", 2, 2, []e{
			e{"a", "b"},
			e{"b", "a"},
		}},
		testcase{"a -> b\nb -> a", 2, 2, []e{
			e{"a", "b"},
			e{"b", "a"},
		}},
		testcase{"a b -> c d", 4, 4, []e{
			e{"a", "c"}, e{"a", "d"},
			e{"b", "c"}, e{"b", "d"},
		}},
		testcase{"a b -- c d", 4, 8, []e{
			e{"a", "c"}, e{"a", "d"},
			e{"b", "c"}, e{"b", "d"},
			e{"c", "a"}, e{"c", "b"},
			e{"d", "a"}, e{"d", "b"},
		}},
		testcase{"a b c -- a b c", 3, 6, []e{
			e{"a", "b"}, e{"a", "c"},
			e{"b", "a"}, e{"b", "c"},
			e{"c", "a"}, e{"c", "b"},
		}},
		testcase{"aaa -- bbb", 2, 2, []e{
			e{"aaa", "bbb"}, e{"aaa", "bbb"},
		}},
		testcase{`
			a -> b c
			b -> c
			c -- d
			d -> a c
		`, 4, 6, []e{
			e{"a", "b"}, e{"a", "c"},
			e{"b", "c"}, e{"c", "d"}, e{"d", "a"}, e{"d", "c"},
		}},
		testcase{`
			a a -> a a a
			a -- a a
		`, 1, 0, []e{}},
		testcase{"йö -> Ы 漢字", 3, 2, []e{
			e{"йö", "Ы"}, e{"йö", "漢字"},
		}},
		testcase{"😞-> 😀", 2, 1, []e{
			e{"😞", "😀"},
		}},
		testcase{";;;;a -> b;;;", 2, 1, []e{
			e{"a", "b"},
		}},
		testcase{"a -> b c ; b -> c ; c -- d \n d -> a c", 4, 6, []e{
			e{"a", "b"}, e{"a", "c"},
			e{"b", "c"}, e{"c", "d"}, e{"d", "a"}, e{"d", "c"},
		}},
	}

	for i, c := range cases {
		g := graph.New()
		err := Parse(g, c.text)
		if err != nil {
			t.Errorf("#%d: parse error: %s", i, err)
			return
		}

		if g.NodesCount() != c.nodeCount {
			t.Errorf("#%d: wrong nodes count - %d, expected %d", i, g.NodesCount(), c.nodeCount)
		}

		if g.EdgesCount() != c.edgeCount {
			t.Errorf("#%d: wrong edges count - %d, expected %d", i, g.EdgesCount(), c.edgeCount)
		}

		for _, e := range c.edges {
			a, ok := g.NodeByLabel(e.a)
			if !ok {
				t.Errorf("#%d: no node with label %s", i, e.a)
			}
			b, ok := g.NodeByLabel(e.b)
			if !ok {
				t.Errorf("#%d: no node with label %s", i, e.b)
			}
			if !g.HasEdgeBetween(a, b) {
				t.Errorf("#%d: no edge %s->%s", i, e.a, e.b)
			}
		}
	}
}

func TestGralangSyntaxErrors(t *testing.T) {
	cases := []string{
		"a <-> b",
		"a ->",
		"a --",
		"-- a",
		"-> a",
	}

	for i, text := range cases {
		g := graph.New()
		err := Parse(g, text)
		if err == nil {
			t.Errorf("#%d: '%s' parsed without error", i, text)
			return
		}
	}
}

func TestBadGraph(t *testing.T) {
	err := Parse(nil, "")
	if err == nil {
		t.Errorf("parse should not accept nil graphs")
	}
}

func BenchmarkParse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		g := graph.New()
		b.StartTimer()
		Parse(g, `
			a -- b c d
			b -> c d
			Й -> 漢
		`)
	}
}
