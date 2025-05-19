package model

import "github.com/go-gl/gl/v2.1/gl"

type Polygon struct {
	Vertices    []Vertex
	VertexCount int
}

func NewPolygon(vertices []Vertex) Polygon {
	return Polygon{
		Vertices:    vertices,
		VertexCount: len(vertices),
	}
}

func NewPolygonWithRemap(vertices []Vertex, u0, v0, u1, v1 int) Polygon {
	polygon := NewPolygon(vertices)
	polygon.Vertices[0] = polygon.Vertices[0].Remap(float32(u1), float32(v0))
	polygon.Vertices[1] = polygon.Vertices[1].Remap(float32(u0), float32(v0))
	polygon.Vertices[2] = polygon.Vertices[2].Remap(float32(u0), float32(v1))
	polygon.Vertices[3] = polygon.Vertices[3].Remap(float32(u1), float32(v1))

	return polygon
}

func (polygon Polygon) Render() {
	gl.Color3f(1.0, 1.0, 1.0)

	for i := 3; i >= 0; i-- {
		v := polygon.Vertices[i]
		gl.TexCoord2f(v.U/64.0, v.V/32.0)
		gl.Vertex3f(v.Pos.X, v.Pos.Y, v.Pos.Z)
	}
}
