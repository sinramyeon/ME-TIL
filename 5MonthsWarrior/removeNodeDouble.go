//더블 링크드 리스트 주어진 노드 삭제

func (list *DoublyLinkedList) RemoveNode(key int) bool {
  cur := list.head
  if curr == nil {
    return false
  }
  
  if curr.value == key {
    curr = curr.next
    list.count--
    if curr != nil {
      list.head = curr
      list.head.prev = nil
    } else {
      list.tail = nil // 하나밖에 값이없어서 curr 다음이 nil임
    }
    return true
  }
  
  for curr.next != nil {
    if curr.next.value == key {
      curr.next = curr.next.next
      if curr.next == nil {
        list.tail = curr
      } else {
        curr.next.prev = curr
      }
      list.count--
      return true
    }
    curr = curr.next
  }
  
  return false
}
