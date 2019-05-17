type Node struct {
  value int
  next *Node
}

type Queue struct {
  head *Node
  tail *Node
  size int
}

func (q *Queue) Size() int {
  return q.size
}

func (q *Queue) IsEmpty() bool {
  return q.size == 0
}

func (q *Queue) Peek() (int, bool) {
  if q.IsEmpty() {
    return 0, false
  }
  return q.head.value, true
}

func (q *Queue) Print() {
  temp := q.head
  for temp != nil {
    temp = temp.next
  }
}
