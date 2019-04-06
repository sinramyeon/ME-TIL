func (g *Graph) PathExist(source int, destination int) bool {
  count := g.count
  visited := make([]int, count)
  for i:=0; i<count; i++ {
    visited[i] = 0
  }
  
  visited[source] = 1
  g.DFSRec(source, visited)
  return visited[desnitaion] != 0

}

