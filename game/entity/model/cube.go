package model

import "github.com/go-gl/gl/v2.1/gl"

type Cube struct {
	vertices []Vertex
	polygons []Polygon

	xTexOffs, yTexOffs int

	X, Y, Z          float32
	XRot, YRot, ZRot float32
}

func NewCube(xTexOffs, yTexOffs int) Cube {
	return Cube{
		vertices: []Vertex{},
		polygons: []Polygon{},

		xTexOffs: xTexOffs,
		yTexOffs: yTexOffs,
	}
}

func (cube Cube) SetTexOffs(xTexOffs, yTexOffs int) Cube {
	cube.xTexOffs = xTexOffs
	cube.yTexOffs = yTexOffs

	return cube
}

func (cube Cube) AddBox(x0, y0, z0 float32, w, h, d int) Cube {
	cube.vertices = make([]Vertex, 8)
	cube.polygons = make([]Polygon, 6)

	x1 := x0 + float32(w)
	y1 := y0 + float32(h)
	z1 := z0 + float32(d)

	u0 := NewVertex(x0, y0, z0, 0.0, 0.0)
	u1 := NewVertex(x1, y0, z0, 0.0, 8.0)
	u2 := NewVertex(x1, y1, z0, 8.0, 8.0)
	u3 := NewVertex(x0, y1, z0, 8.0, 0.0)
	l0 := NewVertex(x0, y0, z1, 0.0, 0.0)
	l1 := NewVertex(x1, y0, z1, 0.0, 8.0)
	l2 := NewVertex(x1, y1, z1, 8.0, 8.0)
	l3 := NewVertex(x0, y1, z1, 8.0, 0.0)

	cube.vertices[0] = u0
	cube.vertices[1] = u1
	cube.vertices[2] = u2
	cube.vertices[3] = u3
	cube.vertices[4] = l0
	cube.vertices[5] = l1
	cube.vertices[6] = l2
	cube.vertices[7] = l3

	cube.polygons[0] = NewPolygonWithRemap([]Vertex{l1, u1, u2, l2}, cube.xTexOffs+d+w, cube.yTexOffs+d, cube.xTexOffs+d+w+d, cube.yTexOffs+d+h)
	cube.polygons[1] = NewPolygonWithRemap([]Vertex{u0, l0, l3, u3}, cube.xTexOffs+0, cube.yTexOffs+d, cube.xTexOffs+d, cube.yTexOffs+d+h)
	cube.polygons[2] = NewPolygonWithRemap([]Vertex{l1, l0, u0, u1}, cube.xTexOffs+d, cube.yTexOffs+0, cube.xTexOffs+d+w, cube.yTexOffs+d)
	cube.polygons[3] = NewPolygonWithRemap([]Vertex{u2, u3, l3, l2}, cube.xTexOffs+d+w, cube.yTexOffs+0, cube.xTexOffs+d+w+w, cube.yTexOffs+d)
	cube.polygons[4] = NewPolygonWithRemap([]Vertex{u1, u0, u3, u2}, cube.xTexOffs+d, cube.yTexOffs+d, cube.xTexOffs+d+w, cube.yTexOffs+d+h)
	cube.polygons[5] = NewPolygonWithRemap([]Vertex{l0, l1, l2, l3}, cube.xTexOffs+d+w+d, cube.yTexOffs+d, cube.xTexOffs+d+w+d+w, cube.yTexOffs+d+h)

	return cube
}

func (cube Cube) SetPos(x, y, z float32) Cube {
	cube.X = x
	cube.Y = y
	cube.Z = z

	return cube
}

func (cube Cube) Render() {
	c := float32(57.29578)

	gl.PushMatrix()
	gl.Translatef(cube.X, cube.Y, cube.Z)
	gl.Rotatef(cube.ZRot*c, 0.0, 0.0, 0.1)
	gl.Rotatef(cube.YRot*c, 0.0, 1.0, 0.0)
	gl.Rotatef(cube.XRot*c, 1.0, 0.0, 0.0)
	gl.Begin(gl.QUADS)

	for _, polygon := range cube.polygons {
		polygon.Render()
	}

	gl.End()
	gl.PopMatrix()
}
