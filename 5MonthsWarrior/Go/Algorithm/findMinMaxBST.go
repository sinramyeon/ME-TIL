func (t *Tree) FindMin() (int, bool) {

  var node *Node = t.root
  if node == nil {
    return 0, false
  }
  for node.left != nil {
    node = node.left
  }
return node.value, true
}

func (t *Tree) FindMax() (int, bool) {
  var node *Node = t.root
  if node == nil {
    return 0, false
  }
  for node.right != nil {
    node = node.right
  }
  
  return node.value, true

}
