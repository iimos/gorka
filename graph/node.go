package graph

import (
	"fmt"
	"xxx/complexity/graph/types"
)

// Node is a graph node
// type Node interface {
// 	ID() int
// 	Label() string
// 	String() string
// }

// Node is a graph node
type Node = types.Node

// Node is graph node
type node struct {
	id    int
	index int
	label string
}

// ID returns node id. Id is sequentional and uniq during graph lifetime.
func (n *node) ID() int {
	// fmt.Printf("\nNNN %v\n", n)
	return n.id
}

// Label returns node label
func (n *node) Label() string {
	return n.label
}

// String returns string representation of the node
func (n *node) String() string {
	return fmt.Sprintf("Node(%s)", n.label)
}
