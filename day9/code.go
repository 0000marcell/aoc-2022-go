package main

import (
  "fmt"
  "math"
  "os"
  "bufio"
  "strconv"
  "strings"
  "example.com/set"
)

type Pos struct {
  row int 
  col int 
}

type Knot struct {
  index int
  prev *Knot
  next *Knot
  pos Pos 
  visited set.Set
}

func (k *Knot) Move(dir string) {
  switch dir {
	case "R":
    k.pos.col++
	case "L":
    k.pos.col--
	case "U":
    k.pos.row--
	case "D":
    k.pos.row++
	default:
	}
  k.visited.Insert(k.pos)
}

func (k *Knot) MovePrev() {
  for curr := k.prev; curr != nil; curr = curr.prev {
    curr.Follow()
 	}
}

func (k *Knot) GetTail() *Knot {
  var tail *Knot
  for curr := k.prev; curr != nil; curr = curr.prev {
    if curr.prev == nil {
      tail = curr
    }
 	}
  return tail
}

func (k *Knot) Follow() {
  head := k.next
  tail := k
  vdiff := head.pos.row - tail.pos.row
  hdiff := head.pos.col - tail.pos.col
  if math.Abs(float64(vdiff)) > 1 {
    // fix diag pos 
    if math.Abs(float64(hdiff)) > 0 {
      if hdiff > 0 {
        tail.pos.col++
      } else {
        tail.pos.col--
      }
      //tail.pos.col = head.pos.col
    }
    if vdiff > 0 {
      tail.pos.row++
    } else {
      tail.pos.row--
    }
  } else if math.Abs(float64(hdiff)) > 1 {
    // fix diag pos 
    if math.Abs(float64(vdiff)) > 0 {
      if vdiff > 0 {
        tail.pos.row++
      } else {
        tail.pos.row--
      }
      //tail.pos.row = head.pos.row
    }
    if hdiff > 0 {
      tail.pos.col++
    } else {
      tail.pos.col--
    }
  }
  k.visited.Insert(k.pos)
}

func main() {
  pos := Pos{
    row: 0,
    col: 0,
  }
  head := Knot{ index: 1, pos: pos }
  head.visited.Insert(pos)
  rhead := &head
  for i := 0; i < 9; i++ {
    tail := Knot{ index: i, pos: pos }
    tail.next = rhead
    tail.visited.Insert(pos)
    rhead.prev = &tail
    rhead = &tail
  }
  head.MovePrev()
  readFile, err := os.Open("input")
  if err != nil {
    fmt.Println(err)
  }
  fileScanner := bufio.NewScanner(readFile)
  for fileScanner.Scan() {
    str := fileScanner.Text()
    strsplit := strings.Split(str, " ")
    count, err := strconv.Atoi(strsplit[1])
    if err != nil {
      fmt.Println("Error during conversion")
      return
    }
    for i := 0; i < int(count); i++ {
      head.Move(strsplit[0])
      head.MovePrev()
    }
  }
  fmt.Println("tail :", head.GetTail().visited.Count())
}
