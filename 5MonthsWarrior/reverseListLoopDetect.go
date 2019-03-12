//reverse list 함수가 링크드 리스트 내 루프가 있을 경우 헤드를 반환

func (list *List) ReverseListLoopDetect() bool {
  tempHead := list.head
  list.Reverse()
  if tempHead == list.head {
    list.Reverse()
    fmt.Println("loop found")
    return true
  }
  
  list.Reverse()
  fmt.Println("loop not found")
  return false
}
