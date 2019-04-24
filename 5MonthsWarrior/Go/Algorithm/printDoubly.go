func (list *DoublyLinkedList) Print() {
  temp := list.head
  for temp != nil {
    fmt.Println(temp.value)
    temp = temp.next
  }
}
