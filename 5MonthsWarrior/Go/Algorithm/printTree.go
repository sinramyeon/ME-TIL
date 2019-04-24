func (t *Tree) PrintAllPath() {
  stk := new(Stack)
  printAllPath(t.,root, stk)
}

func printAllPath(curr *Node, stk *Stack) {

  if curr == nil {
    return
  }
  
  stk.Push(curr.value)
  if curr.left == nil && curr.right == nil {
    stk.Print()
    stk.Pop()
    return
  }

printAllPath(curr.right, stk)
printAllPath(curr.left, stk)
stk.Pop()
}
