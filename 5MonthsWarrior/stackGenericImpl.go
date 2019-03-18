type Stack struct {
  s []interface{}
}

func (s *Stack) Push(value interface{}) {
  s.s = append(s.s, value)
}

func (s *Stack) Pop() interface{} {

  if s.IsEmpty(){
    return nil
  }
  
  length := len(s.s)
  res := s.s[length-1]
  s.s = s.s[:length-1]
  
  return res
}

func (s *Stack) Top() interface{} {
  length := len(s.s)
  res := s.s[length-1]
  
  return res
}
