package graph

import (
	"reflect"
	"strings"
	"testing"
	"github.com/iimos/gorka/gralang"
)

func TestBFSDFS(t *testing.T) {
	type tcase struct {
		s     string
		bfs   bool
		count int
	}
	cases := [...]tcase{
		// BFS
		tcase{"", true, 0},
		tcase{"1", true, 1},
		tcase{"1 -> 1", true, 1},
		tcase{"1 -- 11", true, 2},
		tcase{"1 -> 11 12; 11 -> 111 112; 12 -> 121 122", true, 7},
		tcase{"1 -> 11 12; 11 -> 111 112 12; 12 -> 121 122 1", true, 7},
		// DFS
		tcase{"", false, 0},
		tcase{"1", false, 1},
		tcase{"1 -> 1", false, 1},
		tcase{"1 -- 11", false, 2},
		tcase{"1 -> 11 12; 11 -> 111 112; 12 -> 121 122", false, 7},
		tcase{"1 -> 11 12; 11 -> 111 112 12; 12 -> 121 122 1", false, 7},
	}

	for i, c := range cases {
		g := New()
		gralang.Parse(g, c.s)

		visited := make(map[string]bool)
		order := []string{}

		start, ok := g.NodeByLabel("1")
		if ok {
			handler := func(n Node) bool {
				l := n.Label()
				order = append(order, l)

				if visited[l] {
					t.Errorf("#%d. node %s visited twice: order = %v", i, l, order)
				}
				if len(l) > 1 {
					parent := l[:len(l)-1]
					if !visited[parent] {
						t.Errorf("#%d. node %s should be visited before %s: order = %v", i, parent, l, order)
					}
				}
				if c.bfs {
					if len(l) == 3 {
						if !visited["1"] || !visited["11"] || !visited["12"] {
							t.Errorf("#%d. 3rd level node visited before 2nd level: order = %v", i, order)
						}
					}
				} else {
					if l == "12" && visited["11"] {
						if !visited["111"] && !visited["112"] {
							t.Errorf("#%d. sibling branch was not fully visited: order = %v", i, order)
						}
					}
					if l == "11" && visited["12"] {
						if !visited["121"] && !visited["122"] {
							t.Errorf("#%d. sibling branch was not fully visited: order = %v", i, order)
						}
					}
				}
				visited[l] = true
				return true
			}

			if c.bfs {
				err := TraverseBreadthFirst(g, start, handler)
				if err != nil {
					t.Errorf("#%d. TraverseBreadthFirst error: %s", i, err)
				}
			} else {
				err := TraverseDepthFirst(g, start, handler)
				if err != nil {
					t.Errorf("#%d. TraverseDepthFirst error: %s", i, err)
				}
			}
		}

		if len(visited) != c.count {
			t.Errorf("#%d. not all nodes are visited: %v", i, visited)
		}
	}
}

func TestShortestPath(t *testing.T) {
	type tcase struct {
		s    string
		from string
		to   string
		dist float32
		path []string
	}
	cases := [...]tcase{
		tcase{
			"a -> b",
			"a", "a", 0,
			[]string{},
		},
		tcase{
			"a -> b",
			"a", "b", 1,
			[]string{"a-b"},
		},
		tcase{
			"a -> b c",
			"a", "c", 1,
			[]string{"a-c"},
		},
		tcase{
			"a -> b; b -> c",
			"a", "c", 2,
			[]string{"a-b", "b-c"},
		},
	}
	for i, tc := range cases {
		g := New()
		gralang.Parse(g, tc.s)

		expectedPath := []Edge{}
		from, _ := g.NodeByLabel(tc.from)
		to, _ := g.NodeByLabel(tc.to)

		for _, pair := range tc.path {
			p := strings.SplitN(pair, "-", 2)
			src, _ := g.NodeByLabel(p[0])
			dst, _ := g.NodeByLabel(p[1])
			e := &edge{src: src, dst: dst, weight: 1}
			expectedPath = append(expectedPath, e)
		}

		path, dist, err := ShortestPath(g, from, to)
		if err != nil {
			t.Errorf("#%d. error %s", i, err)
			return
		}

		if dist != tc.dist {
			t.Errorf("#%d. wrong distance: got %f, expected %f", i, dist, tc.dist)
		}
		if !reflect.DeepEqual(path, expectedPath) {
			t.Errorf("#%d. wrong path: got %v, expected %v", i, path, expectedPath)
		}
	}
}
