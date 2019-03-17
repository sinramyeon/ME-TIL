func (list *DoublyCircularLinkedList) RemoveTail() (int, bool) {
  if flist.count == 0 {
    return 0, false
  }
  
  value := list.tail.value
  list.count--
  if list.count == 0 {
    list.head = nil
    list.tail = nil
    return value, true
  }
  
  prev := list.tail.prev
  prev.next = list.head
  list.head.prev = prev
  list.tail = prev
  return value, true
}
