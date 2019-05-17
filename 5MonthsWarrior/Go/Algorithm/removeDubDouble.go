func (list *DoublyLinkedList) RemoveDuplicate() {
  curr := list.head
  var deleteMe *Node
  
  for curr != nil {
    if (curr.next != nil) && curr.value == curr.next.value {
      deleteMe = curr.next
      curr.next = deleteMe.next
      curr.nex.prev = curr
      if deleteMe == list.tail {
        list.tail = curr
      }
    } else {
      curr = curr.next
    }
  }
}
