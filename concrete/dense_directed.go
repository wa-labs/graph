// Copyright ©2014 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package concrete

import (
	"github.com/gonum/graph"
	"github.com/gonum/matrix/mat64"
)

// DirectedDenseGraph represents a graph such that all IDs are in a contiguous
// block from 0 to n-1.
type DirectedDenseGraph struct {
	mat *mat64.Dense
}

// NewDirectedDenseGraph creates a directed dense graph with n nodes.
// If passable is true all nodes will have an edge with unit cost, otherwise
// every node will start unconnected (cost of +Inf).
func NewDirectedDenseGraph(n int, passable bool) *DirectedDenseGraph {
	mat := make([]float64, n*n)
	v := 1.
	if !passable {
		v = inf
	}
	for i := range mat {
		mat[i] = v
	}
	return &DirectedDenseGraph{mat64.NewDense(n, n, mat)}
}

func (g *DirectedDenseGraph) Has(n graph.Node) bool {
	id := n.ID()
	r, _ := g.mat.Dims()
	return 0 <= id && id < r
}

func (g *DirectedDenseGraph) Order() int {
	r, _ := g.mat.Dims()
	return r
}

func (g *DirectedDenseGraph) Nodes() []graph.Node {
	r, _ := g.mat.Dims()
	nodes := make([]graph.Node, r)
	for i := 0; i < r; i++ {
		nodes[i] = Node(i)
	}
	return nodes
}

func (g *DirectedDenseGraph) DirectedEdgeList() []graph.Edge {
	var edges []graph.Edge
	r, _ := g.mat.Dims()
	for i := 0; i < r; i++ {
		for j := 0; j < r; j++ {
			if g.mat.At(i, j) != inf {
				edges = append(edges, Edge{Node(i), Node(j)})
			}
		}
	}
	return edges
}

func (g *DirectedDenseGraph) From(n graph.Node) []graph.Node {
	var neighbors []graph.Node
	id := n.ID()
	r, _ := g.mat.Dims()
	for i := 0; i < r; i++ {
		if g.mat.At(id, i) != inf {
			neighbors = append(neighbors, Node(i))
		}
	}
	return neighbors
}

func (g *DirectedDenseGraph) To(n graph.Node) []graph.Node {
	var neighbors []graph.Node
	id := n.ID()
	r, _ := g.mat.Dims()
	for i := 0; i < r; i++ {
		if g.mat.At(i, id) != inf {
			neighbors = append(neighbors, Node(i))
		}
	}
	return neighbors
}

func (g *DirectedDenseGraph) HasEdge(n, succ graph.Node) bool {
	return g.mat.At(n.ID(), succ.ID()) != inf
}

func (g *DirectedDenseGraph) EdgeFromTo(n, succ graph.Node) graph.Edge {
	if g.HasEdge(n, succ) {
		return Edge{n, succ}
	}
	return nil
}

func (g *DirectedDenseGraph) Cost(e graph.Edge) float64 {
	return g.mat.At(e.Head().ID(), e.Tail().ID())
}

func (g *DirectedDenseGraph) SetEdgeCost(e graph.Edge, cost float64, directed bool) {
	g.mat.Set(e.Head().ID(), e.Tail().ID(), cost)
}

func (g *DirectedDenseGraph) RemoveEdge(e graph.Edge, directed bool) {
	g.mat.Set(e.Head().ID(), e.Tail().ID(), inf)
}

func (g *DirectedDenseGraph) Matrix() *mat64.Dense {
	// Prevent alteration of dimensions of the returned matrix.
	m := *g.mat
	return &m
}

func (g *DirectedDenseGraph) Crunch() {}
