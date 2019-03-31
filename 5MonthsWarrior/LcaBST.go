func (t *Tree) LcaBST(first int, second int) (int, bool) {
  return LcaBST(t.root, first, second)
}

func LcaBST(curr *Node, first int, second int) (int, bool) {
  if curr == nil{
    return 0, false
  }
  
  if curr.value > first && curr.value > second {
    return LcaBST(curr.left, first, second)
  }
  
  if curr.value < first && curr.value < second {
    return LcaBST(curr.right, first, second)
  }
  
  return curr.value, true
}
