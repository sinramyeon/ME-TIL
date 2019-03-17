func (list *CircularLinkedList) FreeList() {
  list.tail = nil
  list.count = 0  
}
