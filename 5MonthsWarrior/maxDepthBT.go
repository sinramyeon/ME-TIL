func (t *Tree) TreeDepth() int {
      return treeDepth(t.root)
}
 
func treeDepth(root *Node) int {
      if root == nil {
            return 0
      }
      
      lDepth := treeDepth(root.left)
      rDepth := treeDepth(root.right)
 
      if lDepth > rDepth {
            return lDepth + 1
      }
      return rDepth + 1
}
