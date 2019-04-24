func (t *Tree) NumFullNodesBT() int {
  return numFullNodesBT(t.root)
}

func numFullNodesBT(curr *Node) int {
  var count int
  if curr == nil {
    return 0
  }
  
  count = numFullNodesBT(curr.right) + numFullNodesBT(curr.left)
  if curr.right != nil && curr.left != nil {
    count++
  }
  
  return count
}
