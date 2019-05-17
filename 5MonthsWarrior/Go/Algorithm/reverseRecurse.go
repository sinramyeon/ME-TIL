// reverse singly linked list using recursion

func (list *List) ReverseRecurse() {
  list.head = list.reverseRecurseUtil(list.head, nil)
}

func (list *List) reverseRecurseUtil(currentNode *Node, nextNode *Node) *Node {
  var ret *Node
  if currentNode == nil {
    return nil
  }
  
  if currentNOde.next == nil {
    currentNode.next = nextNode
    return currentNode
  }
  
  ret = list.reverseRecurseUtil(currentNode.next, currentNode)
  currentNode.next = nextNode
  return ret
}
