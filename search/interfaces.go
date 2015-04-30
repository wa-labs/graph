package search

import (
	"github.com/gonum/graph"
)

// SourceSearchGraph defines the minimal set of methods
// required to "fan out" from a single source node in a search.
type SourceSearchGraph interface {
	graph.Out
}

// CostSearchGraph is the minimal set of methods needed to perform
// a weighted graph search fanning out from a single node.
type CostSearchGraph interface {
	SourceSearchGraph
	graph.DirectedCostGraph
}

// HeuristicSearchGraph is the minimal set of methods needed to perform
// a heuristic search from a single source node.
type HeuristicSearchGraph interface {
	CostSearchGraph
	graph.HeuristicCostGraph
}

// UnitNullGraph is a convenience wrapper for a graph that implements
// a unit cost function and the null heuristic.
type UnitNullGraph struct {
	graph.Graph
}

func (g *UnitNullGraph) Cost(_ graph.Edge) float64 {
	return 1
}

func (g *UnitNullGraph) HeuristicCost(_, _ graph.Node) float64 {
	return 0
}