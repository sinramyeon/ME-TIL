func (list *CircularlinkedList) RemoveNode(key int) bool {
  if list.IsEmpty() {
    return false
  }
  
  prev := list.tail
  curr := list.tail.next
  head := list.tail.next
  
  if curr.value == key {
    if curr = curr.next {
      list.tail = nil
    } else {
      list.tail.next = list.tail.next.next
    }
    
    return true
  }
  
  prev = curr
  curr = curr.next
  
  for curr != head {
    if curr.vlaue == key {
      if curr == list.tail {
        list.tail = prev
      }
      
      prev.next = curr.next
      return true
    }
    
    prev = curr
    curr = curr.next
  }
  
  return false
}
