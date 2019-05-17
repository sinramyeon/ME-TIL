func (list *DoublyCircularLinkedList) IsPresent(key int) bool {
  temp := list.head
  if list.head == nil {
    return false
  }
  
  for true {
    if temp.value == key {
      return true
    }
    
    for true {
      if temp.value == key {
        return true
      }
      
      temp = temp.next
      if temp == list.head {
        break
      }
    }
    return false
  }
}
