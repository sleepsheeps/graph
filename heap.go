package graph

import "sort"

type PriorityNode struct {
	V interface{}
	W float32
}

type PriorityQueue struct {
	d      []*PriorityNode
	m      map[interface{}]int
	offset int
}

func NewPriorityQueue() *PriorityQueue {
	return &PriorityQueue{
		m: make(map[interface{}]int),
	}
}

func (pq *PriorityQueue) Len() int {
	return len(pq.d)
}

func (pq *PriorityQueue) Less(i, j int) bool {
	return pq.d[i].W < pq.d[j].W
}

func (pq *PriorityQueue) Swap(i, j int) {
	pq.d[i], pq.d[j] = pq.d[j], pq.d[i]
	pq.m[pq.d[i]] = j
	pq.m[pq.d[j]] = i
}

func (pq *PriorityQueue) Pop() *PriorityNode {
	if len(pq.d) > 0 && pq.offset < len(pq.d) {
		old := pq.offset
		pq.offset++
		d := pq.d[old]
		delete(pq.m, d.V)
		return d
	}
	return nil
}

func (pq *PriorityQueue) Push(v int, w float32) {
	if pq.InQueue(v) {
		index := pq.m[v]
		pq.d[index].W = w
	} else {
		pq.d = pq.d[pq.offset:]
		pq.offset = 0
		pq.d = append(pq.d, &PriorityNode{
			V: v,
			W: w,
		})
		pq.m[v] = len(pq.d) - 1
	}
	sort.Sort(pq)
}

func (pq *PriorityQueue) InQueue(key interface{}) bool {
	_, ok := pq.m[key]
	return ok
}

func (pq *PriorityQueue) ChangePriority(key interface{}, w float32) {
	index, ok := pq.m[key]
	if !ok {
		return
	}
	pq.d[index].W = w
	sort.Sort(pq)
}
