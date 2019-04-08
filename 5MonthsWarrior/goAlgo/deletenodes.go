// 링크드 리스트 내 등장하는 같은수를 다 지우기?

func (list *List) DeleteNodes(delValue int) {
  currNoce := list.head
  for currNode != nil && currNode.value == delValue {
    list.head = currNode.next
    currNode = list.head
  }
  
  for currNode != nil {
    nextNode := currNode.next
    if nextNode != nil && nextNode.value == delValue {
      currNode.next = nextNode.next
    } else {
      currNode = nextNode
    }
  }
  
}
