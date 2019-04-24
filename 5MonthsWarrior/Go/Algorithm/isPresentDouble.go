func (list *DoublyLinkedList) IsPresent(key int) bool {
  temp := list.head
  for temp != nil {
    if temp.value == key {
      return true
    }
    
    temp = temp.next
  }
  return false
}

// 더블 링크 리스트 값 존재하는지 찾기
