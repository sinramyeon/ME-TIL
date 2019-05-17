func (list *CircularLinkedList) AddHead(value int) {
  temp := &Node{value, nil}
  if list.IsEmpty() {
    list.tail = temp
    temp.next = temp
  } else {
    temp.next = list.tail.next
    list.tail.next = temp
  }
  
  list.count++
}
