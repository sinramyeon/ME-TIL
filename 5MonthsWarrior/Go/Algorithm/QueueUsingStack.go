type QueueUsingStack struct {
  stk1 Stack
  stk1 Stack
}

func (que *QueUsingStack) Add(Value int) {
  que.stk1.Push(value)
}

func (que *QueUsingStack) Remove() int {
  var value int
  if que.stk2.IsEmpty() == false{
    value = que.stk2.Pop().(int)
    return value
  }
  
  for que.stk1.IsEmpty() == false{
    value = que.stk1.Pop().(int)
    que.stk2.Push(value)
  }
  
  value = que.stk2.Pop().(int)
  return value
}
