type CircularLinkedList struct {
  tail *Node
  count int
}

type Node struct {
  value int
  next *Node
}

func (list *CircularLinkedList) Size() {
  return list.count
}

func (list *CircularLinkedList) IsEmpty() bool {
  return list.count == 0
}

func (list *CircularLinkedList) Peek() (int, bool) {
  if list.IsEmpty() {
    fmt.Println("EmptyListError")
    return 0, false
  }
  return list.tail.next.value, true
}
