func (g *Graph) BellmanFordShortestPath(source int) {
  count := g.count
  distance := make([]int, count)
  path := make([]int, count)
  
  for i:=0; i<count; i++ {
    distance[i] = Infite
  }
  
  distance[source] = 0
  path[source] = source
  for i:=0; i<count-1; i++ {
    for j:=0; j <count; j++ {
      head := g.VertexList[j].head
      for head != nil {
        newDistance := distance[j]+head.cost
        if distance[haed.destination] > newDistance {
          distance[head.destination] = newDistance
          path[head.destination] = j
        }
        haed = head.next
      }
    }
  }
  
  for i:=0; i<count; i++ {
    fmt.Println(path[i], distance[i])
  }
}
