package model

import "craftgame/pkg/gl"

type Cube struct {
	vertices []Vertex
	polygons []Polygon

	textureOffsetX, textureOffsetY int

	X, Y, Z                         float32
	XRotation, YRotation, ZRotation float32
}

func NewCube(textureOffsetX, textureOffsetY int) Cube {
	return Cube{
		vertices: []Vertex{},
		polygons: []Polygon{},

		textureOffsetX: textureOffsetX,
		textureOffsetY: textureOffsetY,
	}
}

func (cube Cube) SetTextureOffset(textureOffsetX, textureOffsetY int) Cube {
	cube.textureOffsetX = textureOffsetX
	cube.textureOffsetY = textureOffsetY

	return cube
}

func (cube Cube) AddBox(offsetX, offsetY, offsetZ float32, width, height, depth int) Cube {
	cube.vertices = make([]Vertex, 8)
	cube.polygons = make([]Polygon, 6)

	x := offsetX + float32(width)
	y := offsetY + float32(height)
	z := offsetZ + float32(depth)

	vertexBottom1 := NewVertex(offsetX, offsetY, offsetZ, 0.0, 0.0)
	vertexBottom2 := NewVertex(x, offsetY, offsetZ, 0.0, 8.0)
	vertexBottom3 := NewVertex(x, y, offsetZ, 8.0, 8.0)
	vertexBottom4 := NewVertex(offsetX, y, offsetZ, 8.0, 0.0)

	vertexTop1 := NewVertex(offsetX, offsetY, z, 0.0, 0.0)
	vertexTop2 := NewVertex(x, offsetY, z, 0.0, 8.0)
	vertexTop3 := NewVertex(x, y, z, 8.0, 8.0)
	vertexTop4 := NewVertex(offsetX, y, z, 8.0, 0.0)

	cube.vertices[0] = vertexBottom1
	cube.vertices[1] = vertexBottom2
	cube.vertices[2] = vertexBottom3
	cube.vertices[3] = vertexBottom4
	cube.vertices[4] = vertexTop1
	cube.vertices[5] = vertexTop2
	cube.vertices[6] = vertexTop3
	cube.vertices[7] = vertexTop4

	cube.polygons[0] = NewPolygonWithRemap(
		[]Vertex{vertexTop2, vertexBottom2, vertexBottom3, vertexTop3},
		cube.textureOffsetX+depth+width,
		cube.textureOffsetY+depth,
		cube.textureOffsetX+depth+width+depth,
		cube.textureOffsetY+depth+height,
	)
	cube.polygons[1] = NewPolygonWithRemap(
		[]Vertex{vertexBottom1, vertexTop1, vertexTop4, vertexBottom4},
		cube.textureOffsetX+0,
		cube.textureOffsetY+depth,
		cube.textureOffsetX+depth,
		cube.textureOffsetY+depth+height,
	)
	cube.polygons[2] = NewPolygonWithRemap(
		[]Vertex{vertexTop2, vertexTop1, vertexBottom1, vertexBottom2},
		cube.textureOffsetX+depth,
		cube.textureOffsetY+0,
		cube.textureOffsetX+depth+width,
		cube.textureOffsetY+depth,
	)
	cube.polygons[3] = NewPolygonWithRemap(
		[]Vertex{vertexBottom3, vertexBottom4, vertexTop4, vertexTop3},
		cube.textureOffsetX+depth+width,
		cube.textureOffsetY+0,
		cube.textureOffsetX+depth+width+width,
		cube.textureOffsetY+depth,
	)
	cube.polygons[4] = NewPolygonWithRemap(
		[]Vertex{vertexBottom2, vertexBottom1, vertexBottom4, vertexBottom3},
		cube.textureOffsetX+depth,
		cube.textureOffsetY+depth,
		cube.textureOffsetX+depth+width,
		cube.textureOffsetY+depth+height,
	)
	cube.polygons[5] = NewPolygonWithRemap(
		[]Vertex{vertexTop1, vertexTop2, vertexTop3, vertexTop4},
		cube.textureOffsetX+depth+width+depth,
		cube.textureOffsetY+depth,
		cube.textureOffsetX+depth+width+depth+width,
		cube.textureOffsetY+depth+height,
	)

	return cube
}

func (cube Cube) SetPos(x, y, z float32) Cube {
	cube.X = x
	cube.Y = y
	cube.Z = z

	return cube
}

func (cube Cube) Render() {
	var c float32 = 57.29578

	gl.PushMatrix()

	gl.Translatef(cube.X, cube.Y, cube.Z)
	gl.Rotatef(cube.ZRotation*c, 0.0, 0.0, 1.0)
	gl.Rotatef(cube.YRotation*c, 0.0, 1.0, 0.0)
	gl.Rotatef(cube.XRotation*c, 1.0, 0.0, 0.0)

	gl.Begin(gl.Quads)
	for _, polygon := range cube.polygons {
		polygon.Render()
	}
	gl.End()

	gl.PopMatrix()
}
