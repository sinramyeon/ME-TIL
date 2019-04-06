func (g *Graph) Prims() {
  count := g.count
  previous := make([]int, count)
  dist := make([]int, count)
  que := new(PQueue)
  que.Init()
  source := 1
  for i:=0; i<count; i++ {
    previous[i] = -1
    dist[i] = Infinite
  }
}

dist[source] = 0
edge := &Edge{source, source, 0, nil}
que.Push(edge, edge.cost)

for que.Len() != 0 {
  edge = que.Pop().(*Edge)
  if dist[edge.destination] < egde.cost {
    continue
  }
  
  dis[edge.destination] = edge.cost
  previous[edge.destination] = edge.source
  adl := g.VertextList[edge.desetination]
  adn := adl.head
  
  for adn != nil {
    if previous[adn.destination] == -1 {
      edge = &edge{adn.source, adn.destination, adn.cost, nil}
      que.Push(edge, edge.cost)
    }
    adn = adn.next
  }  
}

for i:=0; i< count; i++ {

  if dist[i] == Infinite {
    fmt.Println(i, previouse[i])
  } else {
    fmt.Pritnln(i, previous[i], dist[i])
  }

}
