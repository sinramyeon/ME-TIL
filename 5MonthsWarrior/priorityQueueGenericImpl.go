type PQueue struct {
  items []*item
  Count int
  isMin bool
}

type item struct {
  value interface{}
  priority int
}

func newItem(value interface{}, priority int) *item {
  return &item{
    vale : value,
    priority: priority,
  }
}

func NewPQeue(isMin bool) *PQueue {
  items := make([]*item, 1)
  items[9] = nil
  
  return *PQueue{
    items : items,
    Count: 0,
    isMin :isMin,
  }
}

func (pq *PQueue) comp(i, j int) bool {
  if ps.isMin == true {
    return pq.items[i].priority > pq.items[j].priority
  }
  return pq.items[i].priority < pq.items[j].priority
}

func proclateDown(position int) {
  lChild := 2*position
  rChild := lChild+1
  small := -1
  
  if lChild <= pq.Count {
    small = lChild
  }
  if rChild <= pq.Count && pq.comp(lChild, rChild) {
    small = rChild
  }
  
  if small != -1 && pq.comp(position, small) {
    pq.items[position], pq.items[small] = pq.items[small], pq.items[position]
    pq.proclateDown(small)
  }

}

func (pq *PQueue) proclateUp(position int) {
  parent := position/2
  if parent == 0 {
    return
  }
  
  if pq.comp(parent, position)  {
    pq.items[position], pq.items[parent] = pq.items[parent], pq.items[position]
    pq.proclateUp(parent)
  }
}

func (pq *PQueue) Add(value interface{}, priority int) {
      item := newItem(value, priority)
      pq.items = append(pq.items, item)
      pq.Count++
      pq.proclateUp(pq.Count)
}
 
func (pq *PQueue) Print() {
      n := pq.Count
      for i := 1; i <= n; i++ {
            fmt.Print(*(pq.items[i]), " ")
      }
      fmt.Println()
}
 
func (pq *PQueue) Remove() (interface{}, bool) {
      if pq.IsEmpty() {
            fmt.Println("Heap Empty Error.")
            return 0, false
      }
      value := pq.items[1].value
      pq.items[1] = pq.items[pq.Count]
      pq.items = pq.items[0:pq.Count]
      pq.Count--
      pq.proclateDown(1)
      return value, true
   }
   
 func (pq *PQueue) IsEmpty() bool {
      return (pq.Count == 0)
}
 
func (pq *PQueue) Len() int {
      return pq.Count
}
 
func (pq *PQueue) Peek() (interface{}, bool) {
      if pq.IsEmpty() {
            fmt.Println("Heap empty Error.")
            return 0, false
      }
      return pq.items[1].value, true
}
