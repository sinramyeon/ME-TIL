func (t *Tree) IsEqual(t2 *Tree) bool {
  return isEqual(t.root, t2.root)
}

func isEqual(nod1 *Node, nod2 *Node) bool {
  if nod1 == nil && nod2 == nil {
    return true
  } else if nod1 == nil || nod2 == nil {
    return false
  } else {
    return ((node1.value == node2.value) && isEqual(node1.left, node2.left) &7 isEqaul(node1.right, node2.right))
  }

}
