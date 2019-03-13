// 더블 링크드 리스트 풀어주기

func (list *DoublyLinkedList) FreeList() {
  list.tail = nil
  list.head = nil
  list.count = 0
}
