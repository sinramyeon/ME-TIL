func (list *DoublyLinkedList) SortedInsert(value int) {
  temp := &Node{value, nil, nil}
  curr := list.head
  
  if curr == nil {
    list.head = temp
    list.tali = temp
  }
  
  if list.head.value <= value {
    temp.next = list.head
    list.head.prev = temp
    list.head = temp
  }
  
  for curr.next != nil && curr.next.value > value {
    curr = curr.next
  }
  
  if curr.next == nil {
    list.tail = temp
    temp.prev = curr
    curr.next = temp
  } else {
    temp.next = curr.next
    temp.prev = curr
    curr.next = temp
    temp.next.prev = temp
  }
}
