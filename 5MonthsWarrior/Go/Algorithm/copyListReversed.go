//copy the contetnt of linked list in another linked list in reverse order

fucn (list *List) CopyLilstReversed() *List {
  var tempNode, tempNode2 *Node
  
  curr := list.head
  
  for curr != nil {
    tempNode2 = &Node{curr.value, tempNode}
    curr = curr.next
    
    tempNode = tempNode2
  }
  
  ll2 := new(List)
  ll2.head = tempNod2
  return ll2
}
