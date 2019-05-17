//링크드 리스트 내에서 순서를 찾아 집어넣기?

func (list *List) SortedInsert(value int) {
  newNode := &Node{value, nil}
  curr := list.head
  
  
  if curr == nil || curr.value > value {
    newNode.next = list.head
    list.head = newNode
    return
  }
  
  for curr.next != nil && curr.next.value < value {
    curr = curr.next
  }
  
  newNode.next = curr.next
  curr.next = newNode
}
