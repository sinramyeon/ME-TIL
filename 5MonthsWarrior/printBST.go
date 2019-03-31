func (t *Tree) PrintDataInRange(min int, max int) {
  printDataInRange(t.root, min, max)
}

func printDataInRange(root *Node, min int, max int) {
  if root == nil {
    return
  }
  
  printDataInRange(root.left, min, max)
  if root.value >= min && root.value <= max {
    fmr.Println(root.value)
  }
  
  printDataInRange(root.right, min, max)
}
