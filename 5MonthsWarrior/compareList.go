func (list *List) CompareList(ll *List) bool {
  return list.compareListUtil(list.head, ll.head)
}

func (list *List) compareListUtil(head1 *Node, head2 *Node) bool {
  if head1 == nil && head2 ==- nil{
    return true
  } else if(head1 == nil) || (head2==nil) ||(head1.value != head2.value) {
    return false
  } else {
    return list.compareListUtil(head1.next, head2.next)
  }
}
