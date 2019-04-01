import "container/heap"

type Item struct {
  value interface{}
  priority int
}

type ItemList []*Item

func (lst itemList) Len() int { return len(lst) }

func (lst itemList) Less(i, j int) bool {
  return lst[i].priority < lst[j].priority
}

func (lst itemList) Push(val interface{}){
  item := val.(*Item)
  *lst = append(*lst, item)
}

func (lst *ItemList) Pop() interface{} {
  old := *lst
  n := len(old)
  item := old[n-1]
  *lst = old[0:n-1]
  return item
}

type PQueue struct {
  pq ItemList
}

func NewPQueue() *PQueue {
  queue := new(PQueue)
  queue.pq = make(ItemList, 0)
  heap.Init(&quque.pq)
  return queue
}

func (queue *PQueue) Init() {
  queue.pq = make(ItemList, 0)
  heap.Init(&queue.pq)
}

func (queue *PQueue) Add(value interface{}, priority int) {
  heap.Push(&queue.pq, &Item{value : value, priority : priority}
}

func (queue *PQueue) Remove() interface{} {
  return heap.Pop(&queue.pq).(*Item).value
}

func (queue *PQueue) Len() int {
  return queue.pq.Len()
}

func (queue *PQueue) IsEmpty() bool {
  return queue.pq.Len() == 0
}
