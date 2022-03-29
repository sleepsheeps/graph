package graph

const (
	E = 15
	V = 8
)

var (
	TidyEWD = [E][3]float32{
		{4, 5, 0.35},
		{5, 4, 0.35},
		{4, 7, 0.37},
		{5, 7, 0.28},
		{7, 5, 0.28},
		{5, 1, 0.32},
		{0, 4, 0.38},
		{0, 2, 0.26},
		{7, 3, 0.39},
		{1, 3, 0.29},
		{2, 7, 0.34},
		{6, 2, 0.40},
		{3, 6, 0.52},
		{6, 0, 0.58},
		{6, 4, 0.93},
	}
)

func MakeTidyEWDGraph() *Graph {
	g := NewGraph(E, V)
	for _, edge := range TidyEWD {
		g.AddEdge(int(edge[0]), int(edge[1]), edge[2])
	}
	return g
}
