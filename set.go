package main

type Coord struct {
  row int
  col int
}

var exists = struct{}{}

type set struct {
  m map[Coord]struct{}
}

func NewSet() *set {
  s := &set{}
  s.m = make(map[Coord]struct{})
  return s
}

func (s *set) Insert(value Coord) bool {
  if _, ok := s.m[value]; ok {
    return false
  }
  s.m[value] = exists
  return true
}

func (s *set) Remove(value Coord) {
  delete(s.m, value)
}

func (s *set) Contains(value Coord) bool {
  _, c := s.m[value]
  return c
}

func main() {
  visited := NewSet()
  visited.Insert(Coord{
    row: 1,
    col: 1,
  })
  visited.Insert(Coord{
    row: 1,
    col: 1,
  })
}

