package gorka

import (
	"fmt"
	"github.com/iimos/gorka/types"
)

// Edge is a graph edge
// type Edge interface {
// 	Dst() Node
// 	Wieght() float32
// 	String() string
// }
type Edge = types.Edge

type edge struct {
	src    types.Node
	dst    types.Node
	weight float32
}

type edgeList []*edge

func (e *edge) From() types.Node {
	return e.src
}

func (e *edge) Dst() types.Node {
	return e.dst
}

func (e *edge) Wieght() float32 {
	return e.weight
}

func (e edge) String() string {
	return fmt.Sprintf("Edge(%s -> %s)", e.src, e.dst)
}
