package trades

import (
	"container/heap"
	"math"
)

// Written based on https://www.geeksforgeeks.org/implementation-of-johnsons-algorithm-for-all-pairs-shortest-paths/

type Neighbour struct {
	destination int
	weight      int
}

func NewNeighbour(destination int, weight int) *Neighbour {
	return &Neighbour{destination: destination,
		weight: weight,
	}
}

type Johnson struct {
	vertices      int
	adjacencyList [][]*Neighbour
}

// On using the below constructor,
// edges must be added manually
// to the graph using addEdge()
func NewJohnson(vertices int) *Johnson {
	g := &Johnson{
		vertices: vertices,
	}

	g.adjacencyList = make([][]*Neighbour, vertices)
	for i := 0; i < vertices; i++ {
		g.adjacencyList[i] = make([]*Neighbour, 0)
	}

	return g
}

// // On using the below constructor,
// // edges will be added automatically
// // to the graph using the adjacency matrix
func NewJohnsonFromMatrix(vertices int, adjacencyMatrix [][]int) *Johnson {
	g := NewJohnson(vertices)

	for i := 0; i < vertices; i++ {
		for j := 0; j < vertices; j++ {
			if adjacencyMatrix[i][j] != 0 {
				g.addEdge(i, j, adjacencyMatrix[i][j])
			}
		}
	}
	return g
}

func NewJohnsonFromGraph(graph *GameGraph) *Johnson {
	vertices := len(graph.matrix)
	g := NewJohnson(vertices)

	index := 0
	for vertex, _ := range graph.matrix {
		graph.IndexByNick[vertex] = index
		index++
	}

	for vertex_name, vertex := range graph.matrix {
		for vertex_target, weight := range vertex {
			i := graph.IndexByNick[vertex_name]
			j := graph.IndexByNick[vertex_target]

			g.addEdge(i, j, int(weight))
		}
	}
	return g
}

func (g *Johnson) addEdge(source int, destination int, weight int) {
	g.adjacencyList[source] = append(g.adjacencyList[source], NewNeighbour(destination, weight))
}

func ArraysFill[T any](array []T, value T) {
	for i := 0; i < len(array); i++ {
		array[i] = value
	}
}

// Dijkstra which is faster than regular one with O(|V|^2)
// This is Dijkstra with heap priority queue. Compltexity O(|E|+|V|log|V|)
// According to wiki page https://en.wikipedia.org/wiki/Dijkstra's_algorithm
func (g *Johnson) dijkstra(source int) []int {
	var distance []int = make([]int, g.vertices)

	pq := make(PriorityQueue, 0)
	item := &Item{
		value_weight: 0,
		priority:     source,
	}
	pq.Push(item)

	ArraysFill(distance, math.MaxInt)
	distance[source] = 0

	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		node := item.priority
		dist := item.value_weight

		for _, neighbour := range g.adjacencyList[node] {
			if dist+neighbour.weight < distance[neighbour.destination] {
				distance[neighbour.destination] = dist + neighbour.weight
				pq.Push(&Item{
					value_weight: distance[neighbour.destination],
					priority:     neighbour.destination,
				})
			}

		}

	}

	return distance
}

// // Returns null if
// // negative weight cycle is detected
func (g *Johnson) bellmanford(source int) []int {
	var distance []int = make([]int, g.vertices)

	ArraysFill(distance, math.MaxInt)
	distance[source] = 0

	for i := 0; i < g.vertices-1; i++ {
		for currentVertex := 0; currentVertex < g.vertices; currentVertex++ {
			for _, neighbour := range g.adjacencyList[currentVertex] {
				if distance[currentVertex] != math.MaxInt && distance[currentVertex]+neighbour.weight < distance[neighbour.destination] {
					distance[neighbour.destination] = distance[currentVertex] + neighbour.weight
				}
			}
		}
	}

	for currentVertex := 0; currentVertex < g.vertices; currentVertex++ {
		for _, neighbour := range g.adjacencyList[currentVertex] {
			if distance[currentVertex] != math.MaxInt && distance[currentVertex]+neighbour.weight < distance[neighbour.destination] {
				return nil
			}

		}
	}

	return distance
}

func remove[T any](slice []T, s int) []T {
	return append(slice[:s], slice[s+1:]...)
}

type DijkstraResult struct {
	source      int
	dist_result []int
}

// // Returns null if negative
// // weight cycle is detected
func (g *Johnson) Johnsons() [][]int {
	// Add a new vertex q to the original graph,
	// connected by zero-weight edges to
	// all the other vertices of the graph

	g.vertices++
	g.adjacencyList = append(g.adjacencyList, make([]*Neighbour, 0))
	for i := 0; i < g.vertices-1; i++ {
		g.adjacencyList[g.vertices-1] = append(g.adjacencyList[g.vertices-1], NewNeighbour(i, 0))
	}

	// Use bellman ford with the new vertex q
	// as source, to find for each vertex v
	// the minimum weight h(v) of a path
	// from q to v.
	// If this step detects a negative cycle,
	// the algorithm is terminated.

	var h []int = g.bellmanford(g.vertices - 1)
	if h == nil {
		return nil
	}

	// Re-weight the edges of the original graph using
	// the values computed by the Bellman-Ford
	// algorithm. w'(u, v) = w(u, v) + h(u) - h(v).

	for u := 0; u < g.vertices; u++ {
		neighbours := g.adjacencyList[u]

		for _, neighbour := range neighbours {
			var v int = neighbour.destination
			var w int = neighbour.weight

			// new weight
			neighbour.weight = w + h[u] - h[v]
		}
	}

	// Step 4: Remove edge q and apply Dijkstra
	// from each node s to every other vertex
	// in the re-weighted graph

	g.adjacencyList = remove(g.adjacencyList, g.vertices-1)
	g.vertices--

	var distances [][]int = make([][]int, g.vertices)

	is_sequential := false

	if is_sequential {
		for s := 0; s < g.vertices; s++ {
			distances[s] = g.dijkstra(s)
		}
	} else {
		dijkstra_results := make(chan *DijkstraResult)
		for s := 0; s < g.vertices; s++ {
			go func(s int) {
				dist_result := g.dijkstra(s)
				dijkstra_results <- &DijkstraResult{
					source:      s,
					dist_result: dist_result,
				}
			}(s)
		}
		for s := 0; s < g.vertices; s++ {
			result := <-dijkstra_results
			distances[result.source] = result.dist_result
		}
	}

	// Compute the distance in the original graph
	// by adding h[v] - h[u] to the
	// distance returned by dijkstra

	for u := 0; u < g.vertices; u++ {
		for v := 0; v < g.vertices; v++ {

			// If no edge exist, continue
			if distances[u][v] == math.MaxInt {
				continue
			}

			distances[u][v] += (h[v] - h[u])
		}
	}

	return distances
}
