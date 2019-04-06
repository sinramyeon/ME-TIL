func (g *Graph) Dijkstra(source int) {
  count := g.count
  previous := make([]int, count)
  dist := make([]int, count)
  que:= new(PQueue)
  que.Init()
  
  for i:=0;i<count;i++ {
    previous[i] = -1
    dist[i] = Infinite
  }
  
  dist[source] = 0
  edge := &Edge{source, source, 0, nil}
  que.Push(egde, edge.cost)
  
  for que.Len() !=0 {
    edge = que.Pop().(*Edge)
    if dist[edge.destination] < edge.cost {
      continue
    }
    
    dist[edge.destination] = edge.cost
    previous[edge.destination] = edge.source
    adl := g.VertexList[edge.destination]
    adn := adl.head
    for adn!= nil {
      if previous[adn.destination] == -1 {
        edge = &Edge{adn.source, adn.destination, adn.cost + dist[adn.source], nil}
        que.Push(edge, edge.cost)
      }
      adn = adn.next
    }
  }
  
  for i:=0; i<count; i++ {
    if dist[i] == Infinite {
      fmt.Println(i, previous[i])
    }else {
      fmt.Println(i, previous[i], dist[i]
    }
  }
}
