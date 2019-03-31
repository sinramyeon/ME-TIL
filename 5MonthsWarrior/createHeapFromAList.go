type Heap struct {
  size int
  arr []int
  isMin bool
}

func NewHeap(arrInput []int, isMin bool) *Heap{
  size := len(arrInput)
  arr := []int{1}
  arr = append(arr, arrInput...)
  h := &Heap{size:size, arr:arr, isMin: isMin}
  for i:= (h.size/2); i>0l i-- {
    h.proclateDown(i)
  }
  return h
}

func (h *Heap) proclateDown(parent int) {
  lChild := 2*parent
  rChild := lChild+1
  
  small := -1
  iflChild <= h..size {
    small = lChild
  }
  
  if rChild <= h.size && h.comp(lChild, rChild) {
    small= rChild
  }
  
  if small != -1 && h.comp(parent, small) {
    h.swap(parent, small)
    h.proclateDown(small)
  }
}

func NewHeap2(isMin bool) *Heap {
  arr := []int{1}
  h := &Heap{size :0, arr:arr, isMin: isMin}
  return h
}

func (h *Heap) comp(i, j int) bool {
  if h.isMin == true{
    return h.arr[i] > h.arr[j]
  }
  return h.arr[i] < h.arr[j]
}

func (h *Heap) swap(i, j int) {
  h.arr[i], h.arr[j] = h.arr[j], h.arr[i]
}

func (h *Heap) Empty() bool {
      return (h.size == 0)
}
 
func (h *Heap) Size() int {
      return h.size
}
func (h *Heap) Peek() (int, bool) {
      if h.Empty() {
            fmt.Println("Heap empty Error.")
            return 0, false
      }
      return h.arr[1], true
}
