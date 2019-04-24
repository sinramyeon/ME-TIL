type DoublyCircularLinkedList struct {
  head *Node
  tail *Node
  count int
}

type Node struct {
  value int
  next *Node
  prev *Node
}

func (list *DoublyCircularLinkedList) Size() int {
  return list.count
}

func (list *DoublyCircularLinekdList) IsEmpty() bool {
  return list.count == 0
}

func (list *DoublyCircularLinkedList) peekHead() (int, bool) {
  if list.IsEmpty() {
    return 0, false
  }
  return list.head.value, true
}
