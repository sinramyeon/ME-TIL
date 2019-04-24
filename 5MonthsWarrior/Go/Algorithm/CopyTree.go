func (t *Tree) CopyTree() *Tree{

  Tree2 := new(Tree)
  Tree2.root = copyTree(t.root)
  return Tree2

}

func copyTree(curr *Node) *Node {
  var temp *Node
  if curr != nil {
    temp = new(Node)
    temp.value = curr.value
    temp.left = copyTree(curr.left)
    temp.right = copyTree(curr.right)
    return temp
  }
  
  return nil
}
