
func (g *Graph) DFS(){
count := g.count
visited := make([]int, count)
for i:=0; i<count; i++ {
  visited[i] = 0
}

for i:=0; i<count; i++ {
  if visited[i] == 0 {
    visited[i] = 1
    g.DFSRec(i, visited)
  }
}


func (g *Graph) DFSRec()index int, visited []int) {

  head := g.VertexList[index].head
  fmt.Println(index)
  for head != nil {
    if visited[head.destination] == 0 {
      visited[head.destination] = 1
      g.DFSRec(head.destination, visited)
      head = head.next
    }
  }

}
