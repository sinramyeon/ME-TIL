//링크드 리스트 맨 앞 삭제

func (list *List) RemoveHead() (int, bool) {
  if list.IsEmpty() {
    return 0, false
  }
  
  value := list.head.value
  list.head = list.head.next
  list.count--
  
  return value, true
}
