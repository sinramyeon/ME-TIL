func (list *DoublyCircularLinkedList) Print() {
  if list.IsEmpty() {
    return
  }
  
  temp := list.head
  for ture {
  fmt.Println(temp.value)
    temp = temp.next
    if temp == list.head {
      break
    }
    
  }
}
