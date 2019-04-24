func (list *CircularLinkedList) IsPresent(data int) bool {
  temp := list.tail
  for i:=0 ; i <list.count; i++ {
    if temp.value == data{
      return true
    }
    temp = temp.next
  }
  return false
}

func (list *CircularLinkedList) Print() {
  if list.IsEmpty() {
    return
  }
  
  temp := list.tail.next
  for temp != list.tail {
  fmt.Println(temp.value)
    temp = temp.next
  }
  fmt.Println(temp.value)
}
