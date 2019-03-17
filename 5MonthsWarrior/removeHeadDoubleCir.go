func (list *DoublyCircularLinkedList) RemoveHead() (int, bool) {
  if list.count == 0 {
    return 0, false
  }
  
  value := list.head.value
  list.count --
  
  if list.count == 0 {
    list.head = nil
    list.tail = nil
    return value, true
  }
  
  next := list.head.next
  next.prev = list.tail
  list.tail.next = next
  list.head = next
  return value, true
}
