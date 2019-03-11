//reverse linked list

func (list *List) Reverse() {

  curr := list.head
  var prev, next *Node
  
  for curr != nil {
    next = curr.next
    curr.next = prev
    prev = curr
    curr= next
  }
  
  list.head = prev
  

}
