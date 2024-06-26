package trades

import (
	"fmt"
	"math"
	"testing"
)

// // Driver Code
func TestDijkstraAPSP(t *testing.T) {
	var vertices int = 4
	var matrix [][]int = [][]int{
		{0, 0, 2, 0},
		{4, 0, 3, 0},
		{0, 0, 0, 2},
		{0, 1, 0, 0},
	}

	// Initialization
	var graph *DijkstraAPSP = NewDijkstraApspFromMatrix(vertices, matrix)

	// Function Call
	distances, parents := graph.DijkstraApsp()
	_ = parents

	// The code fragment below outputs
	// an formatted distance matrix.
	// Its first row and first
	// column represent vertices
	fmt.Println("Distance matrix:")

	fmt.Printf("   \t")
	for i := 0; i < vertices; i++ {
		fmt.Printf("%3d\t", i)
	}

	for i := 0; i < vertices; i++ {
		fmt.Println()
		fmt.Printf("%3d\t", i)
		for j := 0; j < vertices; j++ {
			if distances[i][j] == math.MaxInt {
				fmt.Printf(" X\t")
			} else {
				fmt.Printf("%3d\t",
					distances[i][j])
			}
		}
	}
}

func TestDijkstraAPSPWithGraph(t *testing.T) {
	graph := NewGameGraph()
	graph.SetEdge("a", "b", 5)
	graph.SetEdge("a", "d", 10)
	graph.SetEdge("b", "c", 3)
	graph.SetEdge("c", "d", 1)
	johnson := NewDijkstraApspFromGraph(graph)
	dist, parents := johnson.DijkstraApsp()

	fmt.Println("solved matrix")
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if dist[i][j] == math.MaxInt {
				fmt.Printf("%7s", "INF")
			} else {
				fmt.Printf("%7d", dist[i][j])
			}
		}
		fmt.Println()
	}

	fmt.Println("get dists and paths")
	fmt.Println("a -> c = ", GetDist(graph, dist, "a", "c"), "path=", graph.GetPaths(parents, dist, "a", "c"))
	fmt.Println("a -> b = ", GetDist(graph, dist, "a", "b"), "path=", GetPath(graph, parents, dist, "a", "b"))
	fmt.Println("a -> d = ", GetDist(graph, dist, "a", "d"), "path=", GetPath(graph, parents, dist, "a", "d"))

	fmt.Println("detailed path reconstruction, useful for debug")
	paths := graph.GetPaths(parents, dist, "a", "c")
	for index, path := range paths {
		if path.Dist == 0 && (index != 0 || index != len(paths)-1) {
			continue
		}
		fmt.Println(
			fmt.Sprintf("prev= %20s", path.PrevName),
			fmt.Sprintf("next= %20s", path.NextName),
			fmt.Sprintf("prev_node= %5d", path.PrevNode),
			fmt.Sprintf("next_node= %5d", path.NextNode),
			fmt.Sprintf("dist= %5d", path.Dist),
		)
	}
}
