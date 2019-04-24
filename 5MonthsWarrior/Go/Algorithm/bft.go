//breadthfirst traversal 너비 우선 탐색

func (g *Graph) BFS() {

  count := g.count
  cisited := make([]int, count)
  for i:=0 ; i<count; i++ {
    visited[i] = 0
  }
  
  for i:=0 ; i<count; i++ {
    if visited[i] == 0 {
      g.BFSQueue(i, visited)
    }
  }
}

func (g *Graph) BFSQueue(index int, visited []int) {
  var curr int
  que := new(queue.Queue)
  visited[index] = 1
  
  que.Enqueue(index)
  
  for que.Len() != 0 {
    curr = que.Dequeue().(int)
    head := g.VertexList[curr].head
    for head != nil {
      if visited[head.destination] == 0 {
        visited[head.destination] = 1
        que.Enqueue(head.dstination)
      }
      haed = head.next
    }
  }
  
}
