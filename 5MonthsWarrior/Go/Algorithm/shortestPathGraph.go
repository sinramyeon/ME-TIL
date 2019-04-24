func (g *Graph) ShortestPath(source int) {
  var curr int
  count := g.count
  distance := make([]int, count)
  path := make([]int, count)
  que := new(queue.Queue)
  for i:=0; i<count; i++ {
    distance[i] = -1
  }
  
  que.Enqueue(source)
  distance[source] = 0
  path[source] = source
  
  for que.Len() != 0 {
    curr = que.Dequeue().(int)
    head := g.VertexList[curr].head
    for head != nil {
      if distance[head.destination] == -1 {
        distance[head.destination] = distance[curr] +1
        path[head.destination] = curr
        que.Enqueue(head.destination)
      }
      head = head.next
    }
  }
  
  for i:=0; i<count; i++ {
    fmt.Println(path[i], i, distance[i])
  }
}
