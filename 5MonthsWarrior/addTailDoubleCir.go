func (list *DoublyCircularLinkedList) AddTail(value int) {
  newNode := new(Node)
  newNode.value = value
  
  if list.count == 0 {
    list.head = newNode
    list.tail = newNode
    newNode.next = newNode
    newNodw.prev = newNode
  } else {
    newNode.next = list.tail.next
    newNode.prev = list.tail
    list.tail.next = newNode
    newNode.next.prev = newNode
    list.tail = newNode
  }
  
  list.count++
}
