package pkg

import (
	"fmt"
)

// Connected checks if the given nodes are connected by a path.
func (g *Graph) Connected(a, b int) bool {
	// mission 1: find out if there's a path from a to b

	if a == b {
		return true
	}

	var q Queue
	q.Enqueue(a)

	var visited = make([]bool, g.NumNodes())

	for q.Length() != 0 {
		element := q.Dequeue()
		if element == b {
			return true
		}

		visited[element] = true

		for _, n := range g.Neighbors(element) {
			if !visited[n] {
				q.Enqueue(n)
			}
		}

	}
	return false
}

// ShortestPath returns a path between the nodes with the minimum number of edges.
func (g *Graph) ShortestPath(a, b int) []int {

	if a == b {
		return []int{a}
	}

	var q Queue
	q.Enqueue(a)

	var visited = make([]bool, g.NumNodes())

	var parents = make([]int, g.NumNodes())

	for q.Length() != 0 {
		fmt.Printf("Parents: %v", parents)
		element := q.Dequeue()

		visited[element] = true

		for _, n := range g.Neighbors(element) {
			parents[n] = element
			if !visited[n] {
				q.Enqueue(n)
			}
		}

		if element == b {
			return getPath(a, b, parents)
		}

	}

	return []int{}
}

func getPath(a, b int, parents []int) []int {
	var path []int
	var i = b

	path = append(path, b)

	for parents[i] != a {
		i := parents[i]
		path = append(path, i)
	}
	reversePath(path)
	return path
}

// reversePath will reverse the path, e.g. [5, 7, 2] becomes [2, 7, 5].
func reversePath(path []int) {
	for i := 0; i < len(path)/2; i++ {
		j := len(path) - 1 - i
		path[i], path[j] = path[j], path[i]
	}
}
