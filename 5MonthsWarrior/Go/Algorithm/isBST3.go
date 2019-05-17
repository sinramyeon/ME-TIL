func IsBST3(root *Node) bool {
  if root == nil {
    return true
  }
if root.left != nil && FindMax(root.left).value > root.value {
    return false
}
if root.right != nil && FindMin(root.right).value <= root.value {
    return false
}
   return (IsBST3(root.left) && IsBST3(root.right))
}
