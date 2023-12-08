package internal

import "github.com/Ricerob/skywalk-dsm/internal/types"

type Graph struct {
	Nodes     map[int]types.Node
	Corridors map[int]types.Corridor
}

func (g *Graph) ShortestPath(sourceID, targetID int) []types.Node {
	// Implement Dijkstra's algorithm here
	// Consider opening and closing times when traversing the graph
	// Return the shortest path as a slice of nodes
	return nil
}
