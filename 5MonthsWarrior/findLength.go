func (list *List) findLength() int {
  curr := list.head
  count := 0
  for curr != nil {
    count++
    curr = curr.next
 }
 
 return count
}
