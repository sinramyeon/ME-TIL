type DoublyLinkedList struct {
  head *Node
  tail *Node
  count int

}

type Node struct {
  value int
  next *Node
  prev *Node
}

func (list *DoublyLinkedList) Size() int {
  return list.count
}

func (list *DoublyLinkedList) IsEmpty() bool {
  return list.count == 0
}

func (list *DoublyLinkedList) Peek() (int, bool) {
  if list.IsEmpty(){
    fmt.Println("error")
    return 0, false
  }
  
  return list.head.value, true
}

