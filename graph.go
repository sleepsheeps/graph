package graph

import "sort"

type Graph struct {
	e     int
	edges [][]Edge
}

func NewGraph(e int) *Graph {
	g := &Graph{
		e: e,
	}
	g.edges = make([][]Edge, 0, e)
	return g
}

func (g *Graph) AddEdge(f, t int, w float32) {
	g.edges[f] = append(g.edges[f], Edge{
		to: t,
		w:  w,
	})
}

func (g *Graph) GetEdges(v int) []Edge {
	return g.edges[v]
}

func (g *Graph) SortedEdges(minToMax bool) {
	for _, edges := range g.edges {
		sort.Slice(edges, func(i, j int) bool {
			if minToMax {
				return edges[i].w < edges[j].w
			}
			return edges[i].w > edges[j].w
		})
	}
}

type Edge struct {
	to int
	w  float32
}

func (e *Edge) GetTo() int {
	return e.to
}

func (e *Edge) GetW() float32 {
	return e.w
}
