func (list *CircularLinkedList) RemoveHead() (int, boll) {
  if list.IsEmpty() {
    return 0, false
  }
  
  value := list.tail.next.value
  
  if list.tail == list.tail.next {
    list.tail = nil
  } else {
    list.tail.next = list.tail.next.next
  }
  
  list.count --
  return value, true
}
