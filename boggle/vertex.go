package boggle

//Vertex is a simple struct to holds vertex information
type Vertex struct {
	Label       string
	Index       int
	Adjacencies []*Vertex
}

//AddEdge appends a new vertex object to a slice
func (v *Vertex) AddEdge(vertex *Vertex) {
	v.Adjacencies = append(v.Adjacencies, vertex)
}

//NewVertex create a new Vertex object based on label and index. index controls by Graph
func NewVertex(label string, index int) *Vertex {
	return &Vertex{Label: label, Index: index}
}
