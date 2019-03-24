func (q *Queue) Remove() (int, bool) {
  if q.IsEmpty() {
    return 0, false
  }
  
  value := q.head.value
  q.head = q.head.next
  q.size--
  
  return value, true
}
