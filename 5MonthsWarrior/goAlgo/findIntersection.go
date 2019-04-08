// 두 링크드 리스트가 만나는 지점 찾기

func (list *List) FindIntersection(head *Node, head2 *Node) *Node{

  l1 := 0
  l2 := 0
  
  tempHead := head
  tempHead2 := head2
  
  for tempHead != nil {
    l1++
    tempHead  = tempHead.next
  }
  
  for tempHead2 != nil {
    l2++
    tempHead2 = tempHead2.next
  }
  
  var diff int
  
  if l1 < l2 {
    temp := head
    head = head2
    head2 = temp
    diff = l2 -l1
  } else {
    diff= l1-l2
  }
  
  for; diff>0; diff-- {
    head = head.next
  }
  
  for head != head2 {
    head = head.next
    head2 = head2.next
  }
  
  return head

}
