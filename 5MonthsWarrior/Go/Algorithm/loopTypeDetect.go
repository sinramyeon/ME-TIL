// 링크드 리스트 내 루프 찾기
// 없으면 0 있으면 1
// 원형 리스트면 2

func (list *List) LoopTypeDetect() int {
  slowPtr := list.head
  fastPtr := list.head
  for fastPtr.next != nil && fastPtr.next.next != nil {
    if list.head == fastPtr.next || list.head == fastPtr.next.next {
      fmt.Println("circular")
      return 2
    }
    
    slowPtr = slowPtr.next
    fastPtr = fastPtr.next.next
    if slowPtr == fastPtr{
      fmt.Println("loop found")
      return 1
    }
  }
  
  fmt.Println("loop not found")
  return 0
}
