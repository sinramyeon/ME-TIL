func (g *Graph) TopologicalSort() {

  var count = g.count
  stk := new(stack.Stack)
  
  visited:= make([]int, count)
  for i:=0 ; i<count; i++ {
    visited[i] = 0
  }
  
  for i:=0; i<count; i++ {

  if visited[i] == 0 {
    visited[i] = 1
    g.RopologicalSortDFS(i, visited, stk)
  }
  }
  
  for stk.Len()!=0 {
    fmt.Println(stk.Pop().(int))
  }
}

func (g *Graph) TopologicalSortDFS(index int, visited []int, stk *stack.Stack) {
  head := g.VertextList[index].head
  for head != nil {
    if visted[head.destination] == 0 {
      visited[head.destination] = 1
      g.TopologicalSortDFS(head.destination, visited, stk)
    }
    head = head.next
  }
  stk.Push(index)  
}
