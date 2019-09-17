package alg

import (
	"github.com/iimos/gorka"
)

// ClusteringCoefLocal calculates a local clustering coefficient of the node.
func ClusteringCoefLocal(g graph.Graph, n graph.Node) float32 {
	var triplets, triangles int

	g.NeighbourIter(n, func(u graph.Node) bool {
		g.NeighbourIter(n, func(v graph.Node) bool {
			if u.ID() != v.ID() {
				triplets++
				if g.HasEdgeBetween(u, v) || g.HasEdgeBetween(v, u) {
					triangles++
				}
			}
			return true
		})
		return true
	})

	if triangles == 0 {
		return 0
	}
	// triplets and triangles are counted twice
	return float32(triangles) / float32(triplets)
}

// ClusteringCoef calculates clustering coefficient of the graph
func ClusteringCoef(g graph.Graph) float32 {
	var sum float32
	g.NodeIter(func(n graph.Node) bool {
		sum += ClusteringCoefLocal(g, n)
		return true
	})
	cnt := float32(g.NodesCount())
	return sum / cnt
}
