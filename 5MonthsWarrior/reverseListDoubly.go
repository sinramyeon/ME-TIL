func (list *DoublyLinkedList) ReverseList() {
  curr := list.head
  var tempNode *Node
  for curr != nil {
    tempNode = curr.next
    curr.next = curr.prev
    curr.prev = tempNode
    if curr.prev == nil {
      list.tail = list.head
      list.head = curr
      return
    }
    curr = curr.prev
  }
  return
}
