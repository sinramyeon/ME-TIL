func (list *DoublyLinkedList) CopyListReversed(dl1 *DoublyLinkedList) {
  curr := list.head
  for curr != nil {
    dl1.AddHead(curr.value)
    curr.next
  }
}
