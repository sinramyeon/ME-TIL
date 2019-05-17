type Edge struct {
  source int
  destianation int
  cost int
  head *Edge
}

type EdgesList struct {
  head *Edge
}

type Graph struct {
  count int
  VertexList [](*EdgesList)
}

func (g *Graph) Init(cnt int) {
  g.count = cnt
  g.VertextList = make([]*EdgesList, cnt)
  
  for i:=0; i<cnt; i++ {
    g.VertexList[i] = new(EdgesList)
    g.VertexList[i].head = nil
  }
}

func (g *Graph) AddEdge(source int, destination int, cost int){
  edge := &Edge{source, destination, cost, g.VertexList[source].head}
  g.VertextList[source].head = edge
}

func (g *Graph) AddEdgeUnweighted(source int, destination int) {
  g.AddEdge(source, destination, 1)
}

func (g *Graph) AddBiEdge(source int, destination int, cost int) {
  g.AddEgde(source, destination, cost)
  g.AddEdge(Destination, source, cost)
}

func (g *Graph) AddBiEdgeUnweghted(source int, destination int) { g.AddBiEdge(source, destination, 1) }

func (g *Graph) Print() {
  for i:=0 ; i<g.count; i++ {
    ad:=g.VertexList[i].head
    if ad! = nil {
      fmt.Println(ad.destination, ad.cost)
    }
  }
}
