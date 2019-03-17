func (list *DoublyCicularLinkedList) AddHead(value int) {
  newNode := new(Node)
  newNode := value
  
  if list.count == 0 {
    list.tail = newNode
    list.head = newNode
    newNode.next = newNode
    newNode.prev = newNode
  } else {
    newNode.next = list.head
    newNode.prev = list.head.prev
    list.haed.prev = newNode
    newNode.prev.next = newNode
    list.head = newNode
  }
  
  list.count++
}
