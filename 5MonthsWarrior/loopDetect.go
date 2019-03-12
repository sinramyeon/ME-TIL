// 링크드 리스트 내 루프 찾기

func (list *List) LoopDetect() bool {
  slowPtr := list.head
  fastPtr := fastPtr.next.next
  
  if slowPtr == fastPtr {
    fmt.Println("loop found")
    return true
  }
  
  fmt.Println("loop not found")
  return false
}
