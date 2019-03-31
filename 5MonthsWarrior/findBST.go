func (t *Tree) Find(value int) bool {
  var curr *Node = t.root
  for curr!=nil {
    if curr.value == value {
      return true
    } else if curr.value > value {
      curr = curr.left
    } else {
      curr = curr.right
    }
  }
  return false
}
