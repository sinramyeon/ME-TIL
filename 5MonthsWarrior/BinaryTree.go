func LevelOrderBinaryTree(arr []int) *Tree {
  tree := new(Tree)
  tree.root = levelOrderBinaryTree(arr, 0, len(arr))
  return tree
}

func levelOrderBinaryTree(arr []int, start int, size int) *Node {
  curr := &Node{arr[start], nil, nil}
  
  left := 2*start+1
  right := 2*start +2
  
  if left < size {
    curr.left = levelOrderBinaryTree(arr, left, size)
  }
  if right < size {
    curr.right = levelOrderBinaryTree(arr, right, size)
  }
  
  return curr
}

func (t *Tree) PrintPreOrder() {
  printPreOrder(t.root)
  fmt.Println()
}

func printPreOrder(n *Node) {
  if n == nil {
    return
  }
  
  fmt.Print(n.value, " ")
  printPreOrder(n.left)
  printPreorder(n.right)
}

func (t *Tree) PrintPostOrder() {
  printPostOrder(t.root)
  fmt.Println()
}

func printPostOrder(n *Node) {
  if n == nil {
    return
  }
  
  printPostOrder(n.left)
  printPostOrder(n.right)
  fmt.Print(n.value)
}


func (t *Tree) PrintInOrder() {
  printInOrder(t.root)
  fmt.Println()
}

func printInOrder(n *Node) {
  if n == nil {
    return
  }
  
  printInOrder(n.left)
  fmt.Print(n.value)
  printInOrder(n.right)
}

func (t *Tree) PrintBredthFirst() {
  que := new(queue.Queue)
  var temp *Node
  
  if t.root != nil {
    que.Put(t.root)
  }
  
  for que.Empty() == false {
    temp2, _ := que.Get(1)
    temp = temp2[0].(*Node)
    fmt.Print(temp.value, " ")
    if temp.left != nil {
      que.Put(temp.left)
    }
    if temp.right != nil {
      que.Put(temp.right)
    }
    
    fmt.Println()
  }
  
}


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
      return lDepth +1
    }
    
    return rDepth +1
  }



func (t *Tree) NthPreOrder (index int) {
  var counter int
  nthPreOrder(t.root, index, &counter)
}

func nthPreOrder(node *Node, index int, counter *int) {
  if node != nil {
    (*counter)++
    if *counter == index {
      fmt.Println(node.value)
    }
    
    nthPreOrder(node.left, index, count)
    nthPreOrder(node.right, index, count)
  }
}

func (t *Tree) NthPostOrder(index int) {
  var counter int
  nthPostOrder(t.root, index, &counter)
}

func nthPostOrder(node *Node, index int, counter *int) {
  if node!=nil{
    nthPostOrder(node.left, index, counter)
    nthPostOrder(node.right, index, counter)
    (*counter)++
    if *counter == index {
      fmt.Println(node.value)
    }
  }
}


func (t *Tree) NthInOrder(index int) {
  var counter int
  nthInOrder(t.root, index, &counter)
}


func nthInOrder(node *Node, index int, counter *int) {
  if node!= nil {
    nthInOrder(node.left, index, counter)
    *counter++
    if *counter == index {
      fmt.Println(node.value)
    }
    nthInOrder(node.right, index, counter)
  }

}
