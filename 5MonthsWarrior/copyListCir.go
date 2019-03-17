func (list *CircualrLinkedList) CopyList() *CircularLinkedList {
  d := new(CircularLinkedList)
  curr := list.tail.next
  head := curr
  if curr != nil {
    d.addTail(curr.value)
    curr = curr.next
  }
  
  for curr!= head {
    d.addTail(curr.value)
    curr = curr.next
  }
  
  return d
}
