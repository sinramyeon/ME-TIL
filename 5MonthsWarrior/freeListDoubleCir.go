func (list *DoublyCircularLinkedList) FreeList() {
  list.head = nil
  list.tail = nil
  list.count = 0
}
