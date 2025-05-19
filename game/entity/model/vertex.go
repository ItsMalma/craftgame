package model

type Vertex struct {
	Pos  Vector3
	U, V float32
}

func NewVertex(x, y, z, u, v float32) Vertex {
	return Vertex{
		Pos: NewVector3(x, y, z),
		U:   u,
		V:   v,
	}
}

func NewVertexWithVector3(pos Vector3, u, v float32) Vertex {
	return Vertex{
		Pos: pos,
		U:   u,
		V:   v,
	}
}

func (vertex Vertex) Remap(u, v float32) Vertex {
	vertex.U = u
	vertex.V = v
	return vertex
}
