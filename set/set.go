package set 

type Set struct {
  items []any
}

func (s *Set) Count() int {
  return len(s.items)
}

func (s *Set) Contains(newitem any) bool{
  for _, item := range s.items {
    if item == newitem {
      return true
    }
  }
  return false
}

func (s *Set) Insert(newitem any) {
  if !s.Contains(newitem) {
    s.items = append(s.items, newitem)
  }
}

func (s *Set) Remove(ritem any) {
  var nitems []any
  for _,item  := range s.items {
    if item != ritem {
      nitems = append(nitems, item) 
    }
  }
  s.items = nitems
}

func (s *Set) Pop() any {
  lastelem := s.items[len(s.items) - 1] 
  s.Remove(lastelem)
  return lastelem
}
