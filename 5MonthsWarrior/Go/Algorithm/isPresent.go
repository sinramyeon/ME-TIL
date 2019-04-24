//링크드 리스트 내 검색하기

func (list *List) IsPresent(data int) bool {
  temp := list.head
  for temp != nil {
    if temp.value == data{
      return true
    }
    temp = temp.next
  }
  return false
}
