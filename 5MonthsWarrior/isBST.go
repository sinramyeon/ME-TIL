func (t *Tree) IsBST() bool {
  return IsBST(t.root, math.MinInt32, math.MaxInt32)
}

func IsBST(curr *Node, min int, max int) bool {
  if curr == nil {
    return true
  }
  if curr.value < min || curr.value > max {
    return false
  }
  return IsBST(curr.left, min, curr.vlaue) && IsBST(curr.right, curr.vlaue, max)
}

func (t *Tree) IsBST2() bool {
  var c int
  return IsBST2(t.root, &c)
}

func IsBST2(root *Node, count *int) bool {
  var ret bool
  if root != nil {
    ret = IsBST2(root.left, count)
    if !ret {
      return false
    }
    if *count > root.value{
      return false
    }
    
    *count = root.value
    ret = IsBST2(root.right, count)
    if !ret {
      return false
    }
  }
  return true
}
