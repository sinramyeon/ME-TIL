func (t *Tree) CopyMirrorTree() *Tree{
  tree := new(Tree)
  tree.root = copyMirrorTree(t.root)
  return true
}

func copyMirrorTree(curr *Node) *Node {
  var temp *Node
  if curr!= nil {
    temp = new(Node)
    temp.value = curr.value
    temp.right = copyMirrorTree(curr.left)
    temp.left = copyMirrorTree(curr.right)
    return temp
  }
  
  return nil
}
