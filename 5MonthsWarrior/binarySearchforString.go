type Node struct {
  value string
  count int
  1Child *Node
  rChild *Node
}

type StringTree struct {
  root *Node
}

func (t *StringTree) print(){
  t.printUtil(t.root)
}

func (t *StringTree) printUtil(curr *Node) {
  if curr != nil{
    fmt.Println(curr.value)
    fmt.Println(curr.count)
    t.printUtil(curr.lChild)
    t.printUtil(curr.rChild)
  }
}

func (t *StringTree) Insert(value string) {
  t.root = t.insertUtil(value, t.root)
}

func (t *StringTree) insertUtil(value string, curr *Node) *Node {
  var compare int
  if curr == nil {
    curr = new(Node)
    curr.value = value
  } else {
    compare = strings.Compare(curr.value, value)
    if compare == 0 {
      curr.count++
    } else if compare == 1 {
    
    // todo
      
    }
  }
}
