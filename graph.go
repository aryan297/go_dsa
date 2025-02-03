// Online Go compiler to run Golang program online
// Print "Try programiz.pro" message

package main

import "fmt"

type Graph struct {
	Vertecies []*Vertex
}

type Vertex struct {
	Key int
	adj []*Vertex
}

func (g *Graph) AddVertex(key int) {
	if g.getVertex(key) == nil {
		g.Vertecies = append(g.Vertecies, &Vertex{Key: key})
	} else {
		fmt.Println("alrerady Exist")
	}
}

func (g *Graph) AddEdge(from, to int) {
	k1 := g.getVertex(from)
	k2 := g.getVertex(to)
	if k1 == nil || k2 == nil {
		fmt.Println("No Vertex Present")
	}
	k1.adj = append(k1.adj, k2)
	k2.adj = append(k2.adj, k1)
}

func (g *Graph) getVertex(key int) *Vertex {
	for i, v := range g.Vertecies {
		if v.Key == key {
			return g.Vertecies[i]
		}
	}
	return nil
}

func (g *Graph) Prints() {
	for _, v := range g.Vertecies {
		fmt.Printf("Vertex %d:", v.Key)
		for _, neighbour := range v.adj {
			fmt.Printf("%d ", neighbour.Key)
		}
		fmt.Println()
	}

}

func BFS(g *Graph, start int) {
	visted := make(map[int]bool)
	queue := []*Vertex{g.getVertex(start)}
	visted[start] = true

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		fmt.Println(current.Key, "sss")

		for _, v := range current.adj {
			if !visted[v.Key] {
				queue = append(queue, v)
				visted[v.Key] = true
			}
		}

	}
	fmt.Println()
}

func (g *Graph) DfsHelper(start int, visited map[int]bool) {
	visited[start] = true
	currentIndex := g.getVertex(start)
	fmt.Println(start, "hshs")
	for _, neighbour := range currentIndex.adj {
		if !visited[neighbour.Key] {
			g.DfsHelper(neighbour.Key, visited)

		}
	}

}

func DFS(g *Graph, start int) {
	visited := make(map[int]bool)
	g.DfsHelper(start, visited)
	fmt.Println("")
}

func topologicalSortNew(g *Graph, key int, visited map[int]bool, stack *[]int) {
	visited[key] = true
	currentIndex := g.getVertex(key)

	for _, neighbour := range currentIndex.adj {
		if !visited[neighbour.Key] {
			topologicalSortNew(g, neighbour.Key, visited, stack)
		}
	}
	*stack = append(*stack, key)
}

func topologicalSort(g *Graph) {
	visited := make(map[int]bool)
	stack := []int{}

	// Iterate over the vertices correctly
	for _, vertex := range g.Vertecies {
		if !visited[vertex.Key] {
			topologicalSortNew(g, vertex.Key, visited, &stack)
		}
	}

	// Print the topological order
	fmt.Println("Topological Sort Order:")
	for i := len(stack) - 1; i >= 0; i-- {
		fmt.Println(stack[i])
	}
}

/* func main() {
	g := &Graph{}
	g.AddVertex(1)
	g.AddVertex(3)
	g.AddVertex(2)
	g.AddEdge(1, 3)
	g.AddEdge(2, 3)
	g.Prints()
	BFS(g, 1)
	DFS(g, 1)
	fmt.Println("Try programiz.pro")
} */
