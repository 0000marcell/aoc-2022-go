// This example demonstrates a priority queue built using the heap interface.
package main

import (
	"container/heap"
	"fmt"
)

type Coord struct {
  y int 
  x int
}

type Node struct {
  coord Coord
  cost int
  index int
}

// A PriorityQueue implements heap.Interface and holds Nodes.
type PriorityQueue []*Node

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the lowest priority so we use smaller than here(min-heap).
	return pq[i].cost < pq[j].cost
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*Node)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Node in the queue.
func (pq *PriorityQueue) update(item *Node, coord Coord, cost int) {
	item.coord = coord 
	item.cost = cost 
	heap.Fix(pq, item.index)
}

// This example creates a PriorityQueue with some items, adds and manipulates an item,
// and then removes the items in priority order.
func main() {
	// Some items and their priorities.
  nodes := []Node{}
  i := 0
  for i < 3 {
    nodes = append(nodes, 
      Node{
        cost: i,
        coord: Coord{
          y: 1,
          x: 2,
        },
      },
    )
    i++
  }
  
	// Create a priority queue, put the nodes in it, and
	// establish the priority queue (heap) invariants.
	pq := make(PriorityQueue, 0)
  pq = append(pq, &Node{
    coord: Coord{
      x: 1,
      y: 2,
    },
    cost: 99,
    index: 0,
  })
  pq = append(pq, &Node{
    coord: Coord{
      x: 1,
      y: 2,
    },
    cost: 10,
    index: 0,
  })
	// i = 0
	heap.Init(&pq)

	// Insert a new item and then modify its cost.
	heap.Push(&pq, &Node {
    coord: Coord{
      x: 99,
      y: 99,
    },
		cost: 99,
  })
  //pq.update(node, Coord { x: 99, y: 99 }, 5)

	// Take the items out; they arrive in decreasing priority order.
	for pq.Len() > 0 {
		node := heap.Pop(&pq).(*Node)
    fmt.Println(node.cost)
    fmt.Println(node.coord)
    fmt.Println("============")
	}
}
