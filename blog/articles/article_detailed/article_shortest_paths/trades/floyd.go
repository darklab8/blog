package trades

import (
	"fmt"
	"math"
)

/*
floydWarshall algirthm
*/
type Floyder struct {
	*GameGraph
	dist [][]int
}

var FloydMax = int(math.MaxInt / 2)

func (f *Floyder) mapMatrixEdgeToFloyd(keya VertexName, keyb VertexName, distance int) {
	f.dist[f.IndexByNick[keya]][f.IndexByNick[keyb]] = distance
}

func NewFloyder(graph *GameGraph) *Floyder {
	f := &Floyder{GameGraph: graph}
	return f
}

func (f *Floyder) Calculate() *Floyder {

	len_vertexes := len(f.matrix)

	f.dist = make([][]int, len_vertexes)

	for i := 0; i < len_vertexes; i++ {
		f.dist[i] = make([]int, len_vertexes)
		for j := 0; j < len_vertexes; j++ {
			f.dist[i][j] = FloydMax
		}
	}
	for i := 0; i < len_vertexes; i++ {
		f.dist[i][i] = 0
	}

	index := 0
	for vertex, _ := range f.matrix {
		f.IndexByNick[vertex] = index
		index++
	}

	for vertex_source, vertex_targets := range f.matrix {
		for vertex_target_name, vertex_target_dist := range vertex_targets {
			f.mapMatrixEdgeToFloyd(vertex_source, vertex_target_name, int(vertex_target_dist))
		}
	}

	for k := 0; k < len_vertexes; k++ {
		if k%100 == 0 {
			fmt.Println("starting, k=", k)
		}
		for i := 0; i < len_vertexes; i++ {
			for j := 0; j < len_vertexes; j++ {
				if f.dist[i][k]+f.dist[k][j] < f.dist[i][j] {
					f.dist[i][j] = f.dist[i][k] + f.dist[k][j]
				}
			}
		}
	}
	return f
}
