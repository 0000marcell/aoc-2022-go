package queue 

type Item struct {
  value int
}

type Queue struct {
  items []Item
}

func (q *Queue) Insert(item Item) {
  q.items = append(q.items, item) 
}

func (q *Queue) Len() int {
  return len(q.items)
}

func (q * Queue) Pop() Item {
  removeditem := q.items[len(q.items) - 1]
  q.items = q.items[0:len(q.items) - 1]
  return removeditem
}
