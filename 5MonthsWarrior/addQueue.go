func (q *Queue) Add(value int) {
  temp := &Node{value}
  if q.head == nil{
    q.head = temp
    q.tail = temp
  } else {
    q.tail.next = temp
    q.tail = temp
  }
  
  q.size++
}
