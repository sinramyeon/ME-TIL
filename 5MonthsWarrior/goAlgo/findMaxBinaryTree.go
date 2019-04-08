func (t *Tree) FindMaxBT() int {
  return findMaxBT(t.root)
}

func findMaxBT(curr *Node) int {
  if curr == nil {
    return math.MinInt32
  }
  
  max := curr.value
  left := findMaxBT(curr.left)
  right := findMaxBT(curr.right)
  
  if left > max {
    max = left
  }
  if right > max {
    max = right
  }
  
  return max
}
