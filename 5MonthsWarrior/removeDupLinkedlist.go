// remove duplicates from the linked list

func (list *List) RemoveDuplicate() {
  curr := list.head
  for curr != nil {
    if curr.next != nil && curr.value == currnext.value {
      curr.next = currnext.next
    } else {
      curr = curr.next
    }
  }
}
