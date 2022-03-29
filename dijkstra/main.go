package main

import (
	"fmt"
	"graph"
	"math"
)

//
//  @author: ren hui
//  @date: 2022/3/27
//  @note:
//

const (
	invalidV = -1
)

type dijkstra struct {
	g      *graph.Graph
	s      int
	distTo []float32
	edgeTo []int
	pq     *graph.PriorityQueue
}

func NewDijkstra(g *graph.Graph, s int) *dijkstra {
	v := g.V()
	var (
		distTo []float32
		edgeTo = make([]int, v)
	)
	for i := 0; i < v; i++ {
		distTo = append(distTo, math.MaxFloat32)
		edgeTo[i] = invalidV
	}
	return &dijkstra{
		g:      g,
		s:      s,
		edgeTo: edgeTo,
		distTo: distTo,
		pq:     graph.NewPriorityQueue(),
	}
}

func (d *dijkstra) init() {
	d.distTo[d.s] = 0
	d.edgeTo[d.s] = d.s
	d.pq.Push(d.s, 0)
	for pop := d.pq.Pop(); pop != nil; pop = d.pq.Pop() {
		from := pop.V.(int)
		adjEdges := d.g.GetEdges(from)
		for _, adj := range adjEdges {
			newW := d.distTo[adj.GetFrom()] + adj.GetW()
			if d.distTo[adj.GetTo()] < newW {
				continue
			}
			d.distTo[adj.GetTo()] = newW
			d.edgeTo[adj.GetTo()] = adj.GetFrom()
			d.pq.Push(adj.GetTo(), newW)
		}
	}
}

func (d *dijkstra) findPath(end int) []int {
	if end >= d.g.V() || end < 0 {
		return nil
	}
	if d.edgeTo[end] == invalidV {
		return nil
	}
	var res []int
	res = append(res, end)
	for v := d.edgeTo[end]; v != d.s; v = d.edgeTo[v] {
		res = append(res, v)
	}
	return res
}

func main() {
	g := graph.MakeTidyEWDGraph()
	s := 0
	ds := NewDijkstra(g, s)
	ds.init()
	fmt.Println(ds.findPath(2))
	fmt.Println(ds.findPath(7))
	fmt.Println(ds.findPath(6))
	fmt.Println(ds.findPath(5))
}
