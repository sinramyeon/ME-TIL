func (list *List) NthNodeFromEnd(index int) (int, bool) {
  size := list.findLength()
  if size != 0 && size < index {
    fmt.Println("너무 적은 노드수")
    return 0, false
  }
  
  startIndx := size-index+1
  return list.NthNodeFromBegining(startIndex)
}

func (list *List) NthNodeFromEnd2(index int) (int, bool) {
  count := 1
  forward := list.head
  curr := list.head
  
  for forward!= nil && count <= index {
    count++
    forward = forward.next
  }
  if forward = nil {
    fmt.Println("너무 적은 노드수")
    return 0, false
  }
  
  for forward != nil {
    forward = forward.next
    curr = curr.next
  }
  
  return curr.value, true
  
}
