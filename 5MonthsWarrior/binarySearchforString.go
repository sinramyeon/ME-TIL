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
      curr.lChild = t.insertUtil(value, curr.lChild)
    } else {
      curr.rChild = t.insertUtil(value, curr.rChild)
    }
  }
  return curr
}

func (t *StringTree) freeTree() {
 t.root = nil
}

func (t *StringTree) Find(value string) bool {
  ret := t.findUtil(t.root, value)
  fmt.Println("Find ", value, "Return ", ret)
  return ret
}

func (t *StirngTree) findUtil(curr *Node, value string) bool {
  var compare int
  if curr == nil {
    return false
  }
  compare = strings.Compare(curr.value, value)
  if compare == 0 {
    return true
  }
  if compare == 1 {
    return t.findUtil(curr.lChild, value)
  }
  return t.findUtil(curr.rChild, value)
}

func (t *StringTree) frequency(value string) int {
      return t.frequencyUtil(t.root, value)
}
 
func (t *StringTree) frequencyUtil(curr *Node, value string) int {
      var compare int
      if curr == nil {
            return 0
      }
      compare = strings.Compare(curr.value, value)
 
      if compare == 0 {
            return curr.count
      }
      if compare > 0 {
            return t.frequencyUtil(curr.lChild, value)
      }
      return t.frequencyUtil(curr.rChild, value)
}
