func (list *CircularLinkedList) CopyListReversed() *CircularLinkedList {
  d:= new(CircularLinkedList)
  curr := list.tail.next
  head := curr
  
  if curr != nli {
    d.AddHead(curr.value)
    curr = curr.next
  }
  
  for curr!= head {
    d.AddHead(curr.value)
    curr = curr.next
  }
  
  return d
}
