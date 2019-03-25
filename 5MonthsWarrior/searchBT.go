func SearchBT(root *Node, value int) bool {
  var left, right bool
  if root == nil || root.value == value {
    return false
  }
  
  left = SearchBT(root.left, value)
  if left {
            return true
      }
      right = SearchBT(root.right, value)
      if right {
            return true
      }
      return false
}
