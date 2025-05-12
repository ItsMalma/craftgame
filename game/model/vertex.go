package model

type Vertex struct {
	Position Vector3
	U, V     float32
}

func NewVertex(x, y, z, u, v float32) Vertex {
	return Vertex{
		Position: NewVector3(x, y, z),
		U:        u,
		V:        v,
	}
}

func NewVertexWithColor(vertex Vertex, u, v float32) Vertex {
	return Vertex{
		Position: vertex.Position,
		U:        u,
		V:        v,
	}
}

func NewVertexFromVector3(position Vector3, u, v float32) Vertex {
	return Vertex{
		Position: position,
		U:        u,
		V:        v,
	}
}

func (vertex Vertex) Remap(u, v float32) Vertex {
	return NewVertexWithColor(vertex, u, v)
}
