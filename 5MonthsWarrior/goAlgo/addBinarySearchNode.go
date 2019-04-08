func (t *Tree) Add(value int) {
  t.root = addUtil(t.root, value)
}

func addUtil(n *Node, value int) *Node {
  if n == nil{
    n = new(Node)
    n.value = value
    return n
  }
  
  if value < n.value {
    n.left = addUtil(n.left, value)
  } else {
    n.right = addUtil(n.right, value)
  }
  
  return n
}
