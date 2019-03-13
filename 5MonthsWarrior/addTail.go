// 더블 링크드 리스트 꼬리에 더하기

func (list *DoublyLinkedList) AddTail(value int){

  newNode := &Node{value, nil, nil}
  if list.count == 0 {
    list.head = newNode
    list.tail = newNode
  } else {
    newNode.prev = list.tail
    list.tail.next = newNode
    list.tail = newNode
  }
  
  list.count++

}
