//copy the content of given linked list into another linked list

func (list *List) CopyList() *List {
  var headNode, tailNode, tempNode *Node
  curr := list.head
  
  if curr == nil {
    ll2 := new(List)
    ll2.head = nil
    return ll2
  }
  
  headNode = &Node{curr.value, nil}
  tailNode = headNode
  curr = curr.next
  
  for curr != nil {
    tempNode = &Node{curr.value, nil}
    tailNode.next = tempNode
    tailNode = tempNode
    curr = curr.next
  }
  
  ll2 := new(List)
  ll2.head = headNode
  return ll2
}
