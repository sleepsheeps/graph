package graph

import "sort"

type Graph struct {
	e     int
	v     int
	edges [][]Edge
}

func NewGraph(e, v int) *Graph {
	g := &Graph{
		e: e,
		v: v,
	}
	g.edges = make([][]Edge, v)
	return g
}

func (g *Graph) AddEdge(f, t int, w float32) {
	g.edges[f] = append(g.edges[f], Edge{
		from: f,
		to:   t,
		w:    w,
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

func (g *Graph) V() int {
	return g.v
}

type Edge struct {
	from int
	to   int
	w    float32
}

func (e Edge) GetFrom() int {
	return e.from
}

func (e Edge) GetTo() int {
	return e.to
}

func (e Edge) GetW() float32 {
	return e.w
}
