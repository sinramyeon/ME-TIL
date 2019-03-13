// 더블 링크드 리스트 헤드 삭제

func (list *DoublyLinkedList) RemoveHead() (int, bool) {
  if list.IsEmpty() {
    fmt.Println("empty")
    return 0, false
  }
  
  value := list.head.value
  list.head = list.head.next
  if list.head == nil {
    list.tail = nil
  }else {
    list.head.prev = nil
  }
  list.count --
  return value, true
}
