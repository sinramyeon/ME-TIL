func CreateBinaryTree(arr []int) *Tree {
  t := new(Tree)
  size := len(arr)
  t.root = createBinaryTreeUtil(arr, 0, size-1)
  
  return t
}

func createBinaryTreeUtil(arr []int, start, end int) *Node {
  if start > end {
    return nil
  }
  
  mid := (start+end)/2
  curr := new(Node)
  
  curr.value = arr[mid]
  curr.left = createBinaryTreeUtil(arr, start, mid-1)
  curr.right = createBinaryTreeUtil(arr, mid+1, end)
  
  return curr
}
