func (t *Tree) MaxLengthPathBT() int {
  return maxLengthPathBT(t.root)
}

func maxLengthPathBT(curr *Node) int {
  var max, leftPath, rightPath, leftMax, rightMax int
  
  if curr = nil {
    return 0
  }
  
  leftPath = treeDepth(curr.left)
  rightPath = treeDepth(curr.right)
  
  max = leftPath + rightPath +1
  leftMax = maxLengthPathBT(curr.left)
  rightMax = maxLengthPathBT(curr.right)
  
  if leftMax > max {
            max = leftMax
      }
      if rightMax > max {
            max = rightMax
      }
      return max
}
