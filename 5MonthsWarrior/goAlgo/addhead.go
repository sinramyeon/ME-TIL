// 더블 링크드 리스트 헤드에 더하기
func (list *DoublyLinkedList) AddHead(value int) {
  newNode := &Node{value, nil, nil}
  if list.count == 0 {
    list.tail = newNode
    list.head = newNode
  } else {
    list.head.prev = newNode
    newNode.next = list.head
    list.head = newNode
  }
  list.count++
}
