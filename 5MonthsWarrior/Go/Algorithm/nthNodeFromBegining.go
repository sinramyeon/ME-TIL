//find nth node from beginning

func (list *List) NthNodeFromBegining(index int) (int, bool) {
 if index > list.Size() || index < 1 {
  fmt.Println("노드수불충분")
  return 0, false
 }
 
 count := 0
 curr := list.head
 for curr != nil && count <index-1 {
  count++
  curr = curr.next
  
 }
 
 return curr.value, true
}
