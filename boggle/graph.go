package boggle

//Graph holds all the vertices.
type Graph struct {
	Vertices []*Vertex
	Count    int
}

//AddVertex addes vertex to graph based on label
func (g *Graph) AddVertex(label string) {
	g.Vertices = append(g.Vertices, NewVertex(label, g.Count))
	g.Count++
}

//FindVertexByLabel is a helper method to find vertex object by label
func (g *Graph) FindVertexByLabel(label string) *Vertex {
	for _, v := range g.Vertices {
		if v.Label == label {
			return v
		}
	}

	return nil
}

//AddEdge connects to vertices
func (g *Graph) AddEdge(from string, to string) {
	f := g.FindVertexByLabel(from)
	t := g.FindVertexByLabel(to)

	f.AddEdge(t)
}

func genrateAllPossibleWords(g *Graph, index int, state []bool, str string, cb func(string)) {
	state[index] = true
	str = str + g.Vertices[index].Label
	cb(str)
	adjacencies := g.Vertices[index].Adjacencies
	for _, adj := range adjacencies {
		if state[adj.Index] == false {
			genrateAllPossibleWords(g, adj.Index, state, str, cb)
		}
	}
	state[index] = false
}

//GenrateAllPossibleWords generates all possible words based on modfied version of DFS
func GenrateAllPossibleWords(g *Graph, cb func(string)) {
	for i := 0; i < g.Count; i++ {
		state := make([]bool, g.Count)
		genrateAllPossibleWords(g, i, state, "", cb)
	}
}
