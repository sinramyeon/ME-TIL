//링크드 리스트 내 노드 삭제

func (list *List) DeleteNode(delValue int) bool {
  temp := list.head
  if list.IsEmpty() {
    return false
  }
  
  if delValue == list.head.value {
    list.head = list.head.next
    list.count--
    return true
  }
  
  for temp.next != nil {
    if temp.next.value == delValue {
      temp.next = temp.next.next
      list.count--
      return true
    }
    temp = temp.next
  }
  return false
}
