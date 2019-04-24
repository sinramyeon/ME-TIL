type Node struct {
  value int
  next *Node
}

type Stack struct {
  head *Node
  size int
}

func (s *Stack) Size() int {
  return s.size
}

func (s *Stack) IsEmpty() bool {
  return s.size == 0
}

func (s *Stack) Peek() (int, bool) {
  if s.IsEmpty() {
    return 0, false
  }
  
  return s.head.value, true
}

func (s *Stack) Push(value int) {
  s.head = &{value, s.head}
  s.size++
}

func (s *Stack) Pop() (int, bool) {
    if s.IsEmpty() {
      reutrn 0, false
    }
    
    value := s.head.value
    s.head = s.head.next
    
    s.size--
    return value, true
}


func (s *Stack) Print() {
  temp := s.head
  
  for temp!= nil {
    fmt.Println(temp.value)
    temp = temp.next
  }
}
