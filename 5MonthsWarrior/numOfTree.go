func (t *Tree) NumNodes() int {
  return numNodes(t.root)
}

func numNodes(curr *Node) int {
  if curr == nil {
    return 0
  }
  
  return (1+numNodes(curr.right)+numNodes(curr.left))
}

func (t *Tree) NumLeafNodes() int {
  return numLeafNodes(t.root)
}

func numLeafNodes(curr *Node) int {
  if curr == nil {
    return 0
  }
  
  if curr.left == nil && curr.right == nil {
    return 1
  }
  
  return (numLeafNodes(curr.right) + numLeafNodes(curr.left))
}
