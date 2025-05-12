package model

import (
	"craftgame/pkg/gl"
)

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

func NewPolygonWithRemap(vertices []Vertex, minU, minV, maxU, maxV int) Polygon {
	polygon := NewPolygon(vertices)
	polygon.Vertices[0] = polygon.Vertices[0].Remap(float32(maxU), float32(minV))
	polygon.Vertices[1] = polygon.Vertices[1].Remap(float32(minU), float32(minV))
	polygon.Vertices[2] = polygon.Vertices[2].Remap(float32(minU), float32(maxV))
	polygon.Vertices[3] = polygon.Vertices[3].Remap(float32(maxU), float32(maxV))

	return polygon
}

func (polygon Polygon) Render() {
	gl.Color3f(1.0, 1.0, 1.0)

	for i := 3; i >= 0; i-- {
		vertex := polygon.Vertices[i]
		gl.TexCoord2f(vertex.U/64.0, vertex.V/32.0)
		gl.Vertex3f(vertex.Position.X, vertex.Position.Y, vertex.Position.Z)
	}
}
