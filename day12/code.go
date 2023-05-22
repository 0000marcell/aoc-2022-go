package main

import (
	"fmt"
  "container/heap"
  "os"
  "bufio"
)

type Coord struct {
  y int 
  x int
}

func (c Coord) Neighbours(rows int, cols int) []Coord {
  result := []Coord{}
  // up
  if c.y > 0 {
    result = append(result, Coord {
      x: c.x,
      y: c.y - 1,
    });
  }
  // down
  if c.y < rows - 1 {
    result = append(result, Coord {
      x: c.x,
      y: c.y + 1,
    });
  }
  // left
  if c.x > 0 {
    result = append(result, Coord {
      x: c.x - 1,
      y: c.y,
    });
  }
  // right
  if c.x < cols - 1 {
    result = append(result, Coord {
      x: c.x + 1,
      y: c.y,
    });
  }
  return result
}

func (c Coord) CanMoveFrom(currHeight int, pmap [][]int) bool {
  height := pmap[c.y][c.x]
  return height <= currHeight || height == currHeight + 1
}

func (c Coord) IsEqualTo(p Coord) bool {
  return c.y == p.y && c.x == p.x
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

type ParseResult struct {
  start Coord 
  end Coord 
  pmap [][]int
  rows int
  cols int
  lowestPoints []Coord
}

var exists = struct{}{}



func Parse() ParseResult {
  readFile, err := os.Open("input-2")
  if err != nil {
    fmt.Println(err)
  }
  fileScanner := bufio.NewScanner(readFile)
  pmap := make([][]int, 50)
  var end Coord 
  var start Coord 
  row := 0
  cols := 0
  lowestPoints := []Coord{}
   
  for fileScanner.Scan() {
    str := fileScanner.Text()
    for col := 0; col < len(str); col++ {
      val := int(str[col])
      // set finish position
      if val == 69 {
        val = 122
        end.x = col 
        end.y = row 
      }
      // set start position
      if val == 83 {
        val = 97
        start.x = col 
        start.y = row
      }

      if val == 97 {
        lowestPoints = append(lowestPoints, Coord{
          x: col,
          y: row,
        })
      }
      pmap[row] = append(pmap[row], val)
      if cols <= col {
        cols = col
      }
    }
    row++
  }
  pmap = pmap[0:row]

  result := ParseResult{
    start: start,
    end: end,
    pmap: pmap,
    rows: row,
    cols: cols + 1,
    lowestPoints: lowestPoints,
  }

  return result 
}

func GetShortestPath(parseResult ParseResult) int {
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
  visited := NewSet()
  heap.Push(&pq, &Node {
		cost: 0,
    coord: parseResult.start,
  })
  visited.Insert(parseResult.start)
  shortestPath := 99999999999
	for pq.Len() > 0 {
		curr := heap.Pop(&pq).(*Node)
    if curr.coord.IsEqualTo(parseResult.end) {
      shortestPath = curr.cost
      break
    }
    currHeight := parseResult.pmap[curr.coord.y][curr.coord.x];
    neighbours := curr.coord.Neighbours(parseResult.rows, parseResult.cols);
    candidates := []Coord{}
    for _,n := range neighbours {
      if n.CanMoveFrom(currHeight, parseResult.pmap) {
        candidates = append(candidates, n) 
      }
    }
    for _, candidate := range candidates {
      if visited.Insert(candidate) {
        heap.Push(&pq, &Node {
          cost: curr.cost + 1,
          coord: candidate,
        })
      }
    }
  }
  if shortestPath == -1 {
    fmt.Println("error, wrong shortest path!!!")
  }
  return shortestPath
}

func Part1() {
  parseResult := Parse()
  result := GetShortestPath(parseResult)
  fmt.Println(result)
}

func Part2() {
  parseResult := Parse()
  var start Coord
  overallShortest := 9999999
  for i := -1; i < len(parseResult.lowestPoints); i++ {
    if i == -1 {
      start = parseResult.start
    } else {
      start = parseResult.lowestPoints[i]
    }
    result := GetShortestPath(ParseResult {
      start: start,
      end: parseResult.end, 
      pmap: parseResult.pmap, 
      rows: parseResult.rows, 
      cols: parseResult.cols, 
      lowestPoints: parseResult.lowestPoints,
    })
    if result < overallShortest {
      overallShortest = result
    }
  }
  fmt.Println(overallShortest)
}

func main() {
  //Part1() 
  Part2()
}
