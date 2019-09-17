package graph

import (
	"fmt"
	"strings"
	"sync"
	"github.com/iimos/gorka/types"
)

// Graph is a graph
type Graph = types.Graph

// New graph
func New() Graph {
	return &graph{
		edgesOut:    make(map[int]map[int]*edge),
		edgesIn:     make(map[int]map[int]*edge),
		labelToNode: make(map[string]*node),
		nodeMap:     make(map[int]*node),
	}
}

func newGraph() *graph {
	return &graph{
		edgesOut:    make(map[int]map[int]*edge),
		edgesIn:     make(map[int]map[int]*edge),
		labelToNode: make(map[string]*node),
		nodeMap:     make(map[int]*node),
	}
}

// Graph is graph
type graph struct {
	nodes       []*node
	nodeMap     map[int]*node
	labelToNode map[string]*node
	edgesOut    map[int]map[int]*edge // output edges
	edgesIn     map[int]map[int]*edge // input edges
	lock        sync.RWMutex
	lastNodeID  int
}

// NewNode creates a graph node
func (g *graph) NewNode(label string) (Node, error) {
	g.lock.Lock()
	defer g.lock.Unlock()

	if label != "" {
		_, exists := g.labelToNode[label]
		if exists {
			return nil, fmt.Errorf("node with label '%s' exists", label)
		}
	}

	g.lastNodeID++
	id := g.lastNodeID
	i := len(g.nodes)
	n := &node{index: i, id: id, label: label}
	g.nodes = append(g.nodes, n)
	g.nodeMap[id] = n
	g.edgesOut[id] = make(map[int]*edge)

	if label != "" {
		g.labelToNode[label] = n
	}
	return n, nil
}

// NodeByLabel returns node with given label or nil
func (g *graph) NodeByLabel(label string) (n Node, ok bool) {
	n, ok = g.labelToNode[label]
	return n, ok
}

// AddEdge adds weighted edge between two nodes into the graph
func (g *graph) AddEdge(src, dst Node, weight float32) Edge {
	g.lock.Lock()
	defer g.lock.Unlock()

	sid, did := src.ID(), dst.ID()
	e := &edge{src: src, dst: dst, weight: weight}

	out, ok := g.edgesOut[sid]
	if !ok {
		out = make(map[int]*edge)
		g.edgesOut[sid] = out
	}
	out[did] = e

	in, ok := g.edgesIn[did]
	if !ok {
		in = make(map[int]*edge)
		g.edgesIn[did] = in
	}
	in[sid] = e

	return e
}

// AddBiEdge adds bidirectional edge.
func (g *graph) AddBiEdge(src, dst Node, weight float32) {
	g.AddEdge(src, dst, weight)
	g.AddEdge(dst, src, weight)
}

// NodeIter calls cb for each node of the graph. Stops when cb returns false.
func (g *graph) NodeIter(cb func(n Node) bool) {
	for _, n := range g.nodes {
		if !cb(n) {
			break
		}
	}
}

// NodeEdgeIter calls cb for each edge of the node. Stops when cb returns false.
func (g *graph) NodeEdgeIter(n Node, cb func(e Edge) bool) {
	for _, e := range g.edgesOut[n.ID()] {
		if !cb(e) {
			break
		}
	}
}

// NeighbourIter calls cb for each neighbor of the the given node. Stops when cb returns false.
func (g *graph) NeighbourIter(n Node, cb func(n Node) bool) {
	for _, e := range g.edgesOut[n.ID()] {
		d := e.Dst()
		if !cb(d) {
			break
		}
	}
}

// HasEdgeBetween returns true if there ia a [a->b] edge in graph.
func (g *graph) HasEdgeBetween(a, b Node) bool {
	aid, bid := a.ID(), b.ID()
	_, ok := g.edgesOut[aid][bid]
	return ok
}

// OutDegree returns number of outgoing edges from the node.
func (g *graph) OutDegree(n Node) int {
	g.lock.RLock()
	defer g.lock.RUnlock()

	return len(g.edgesOut[n.ID()])
}

// InDegree returns number of ingoing edges for the node.
func (g *graph) InDegree(n Node) int {
	g.lock.RLock()
	defer g.lock.RUnlock()

	return len(g.edgesIn[n.ID()])
}

// NodesCount returns nodes count in the graph.
func (g *graph) NodesCount() int {
	g.lock.RLock()
	defer g.lock.RUnlock()

	return len(g.nodes)
}

// EdgesCount returns edges count in the graph
func (g *graph) EdgesCount() (count int) {
	g.lock.RLock()
	defer g.lock.RUnlock()

	for _, mp := range g.edgesOut {
		count += len(mp)
	}
	return count
}

// MaxNodeID returns max node id in the graph
func (g *graph) MaxNodeID() int {
	return g.lastNodeID
}

func (g *graph) String() string {
	g.lock.RLock()
	defer g.lock.RUnlock()

	var s string
	var b strings.Builder

	for _, n := range g.nodes {
		s = fmt.Sprintf("%d%s -> [", n.id, n.label)
		b.WriteString(s)
		i := 0
		for _, d := range g.edgesOut[n.ID()] {
			if i != 0 {
				b.WriteByte(' ')
			}
			l := d.dst.Label()
			if len(l) == 0 {
				s = fmt.Sprintf("%d", d.dst.ID())
			} else {
				s = fmt.Sprintf("%d:%s", d.dst.ID(), l)
			}
			b.WriteString(s)
			i++
		}
		b.WriteString("]\n")
	}
	return b.String()
}

// IsConnected checks whether all graph nodes are connected
func (g *graph) IsConnected() bool {
	g.lock.RLock()
	defer g.lock.RUnlock()

	if len(g.nodes) == 0 {
		return false
	}

	start := g.nodes[0]
	visited := 0
	TraverseBreadthFirst(g, start, func(n Node) bool {
		visited++
		return false
	})
	return g.NodesCount() == visited
}
