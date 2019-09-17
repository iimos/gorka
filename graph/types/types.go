package types

// Node is a graph node
type Node interface {
	ID() int
	Label() string
	String() string
}

// Edge is a graph edge
type Edge interface {
	From() Node
	Dst() Node
	Wieght() float32
	String() string
}

// Graph is graph
type Graph interface {
	NewNode(label string) (Node, error)
	NodeByLabel(label string) (n Node, ok bool)

	AddEdge(src, dst Node, weight float32) Edge
	AddBiEdge(src, dst Node, weight float32)
	HasEdgeBetween(a, b Node) bool
	
	NodeIter(cb func(n Node) bool)
	NodeEdgeIter(n Node, cb func(e Edge) bool)
	NeighbourIter(n Node, cb func(n Node) bool)
	
	OutDegree(n Node) int
	InDegree(n Node) int
	
	NodesCount() int
	EdgesCount() (count int)
	
	MaxNodeID() int
	String() string
	IsConnected() bool
}
