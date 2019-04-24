type List struct {
  head *Node
  count int
}

type Node struct {
  value int
  next *Node
}

fund (list *List) Size() int {
  return list.count
}

func (list *List) IsEmpty() bool {
  return (list.count == 0)
}

func (list *List) AddHead(value int) {
  list.head = &Node{value, list,head}
  list.count++
}

func (list *List) addTail(value int) {
  curr := list.head
  newNode := &Node{value, nil}
  
  if curr==nil {
    list.head = newNode
    return
  }
  
  for curr.next != nil {
    curr = curr.next
  }
  
  curr.next = newNode
}

func (list *List) Print() {
  temp := list.head
  for temp ! = nil {
    fmt.Println(temp.value, " ")
    temp = temp.next
  }
  
  fmt.Println(""
}
