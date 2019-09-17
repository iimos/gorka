package gralang

import (
	"errors"
	"fmt"
	"github.com/iimos/gorka/types"
)

const (
	lexNode = iota
	lexEdge
	lexSpace
	lexEol
)

const (
	edgeDir = iota
	edgeBi
)

var lext = [256]uint8{
	'-':  lexEdge,
	'>':  lexEdge,
	'<':  lexEdge,
	'[':  lexEdge,
	']':  lexEdge,
	'\t': lexSpace,
	'\v': lexSpace,
	'\f': lexSpace,
	'\r': lexSpace,
	' ':  lexSpace,
	'\n': lexEol,
	';':  lexEol,
}

// Parse parses Gralang and fill Graph
func Parse(g types.Graph, s string) error {
	if g == nil {
		return errors.New("graph is empty")
	}

	for len(s) > 0 {
		sz := readLn(s)
		s = s[sz:]

		lnodes, sz := readNodeList(s)
		s = s[sz:]

		edge := readSeq(s, lexEdge)
		if edge == "" {
			// case when line consists only of node list
			for _, l := range lnodes {
				obtainNode(g, l)
			}
			continue
		}

		s = s[len(edge):]

		edgelex, err := edgeType(edge)
		if err != nil {
			return err
		}

		if len(lnodes) == 0 {
			return errors.New("empty edge source")
		}

		rnodes, sz := readNodeList(s)
		if len(rnodes) == 0 {
			return errors.New("empty edge destination")
		}

		s = s[sz:]

		for _, ll := range lnodes {
			ln := obtainNode(g, ll)
			for _, rl := range rnodes {
				if ll != rl {
					rn := obtainNode(g, rl)
					switch edgelex {
					case edgeBi:
						g.AddBiEdge(ln, rn, 1)
					case edgeDir:
						g.AddEdge(ln, rn, 1)
					default:
						panic(fmt.Sprintf("gralang.Parse: unknown edge lexem: %d", edgelex))
					}
				}
			}
		}
	}
	return nil
}

func obtainNode(g types.Graph, label string) types.Node {
	n, ok := g.NodeByLabel(label)
	if ok {
		return n
	}
	n, _ = g.NewNode(label)
	return n
}

func readNodeList(s string) ([]string, int) {
	orig := s
	list := make([]string, 0, 8)
	sz := readSpaces(s)
	s = s[sz:]

	for {
		node := readSeq(s, lexNode)
		l := len(node)
		if l == 0 {
			break
		}
		list = append(list, node)
		s = s[l:]

		sz = readSpaces(s)
		s = s[sz:]
	}
	return list, len(orig) - len(s)
}

func readSpaces(s string) int {
	for i := 0; i < len(s); i++ {
		ch := s[i]
		if lext[ch] != lexSpace {
			return i
		}
	}
	return len(s)
}

func readLn(s string) int {
	for i := 0; i < len(s); i++ {
		ch := s[i]
		if lext[ch] != lexEol {
			return i
		}
	}
	return len(s)
}

func readSeq(s string, lextype int) string {
	i, n := 0, len(s)
	for ; i < n; i++ {
		ch := s[i]
		if lext[ch] != uint8(lextype) {
			break
		}
	}
	return s[:i]
}

func edgeType(s string) (int, error) {
	switch s {
	case "--":
		return edgeBi, nil
	case "->":
		return edgeDir, nil
	}
	return 0, fmt.Errorf("Wrong edge syntax: '%s'. Expected '--' or '->'", s)
}
