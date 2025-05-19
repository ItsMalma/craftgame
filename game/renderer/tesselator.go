package renderer

import (
	"unsafe"

	"github.com/go-gl/gl/v2.1/gl"
)

const maxVertices = 100000

type Tesselator struct {
	vertexBuffer   [3 * maxVertices]float32
	texCoordBuffer [2 * maxVertices]float32
	colorBuffer    [3 * maxVertices]float32

	vertices int32

	u, v, r, g, b        float32
	hasColor, hasTexture bool
}

func NewTesselator() *Tesselator {
	tesselator := new(Tesselator)
	tesselator.vertexBuffer = [3 * maxVertices]float32{}
	tesselator.texCoordBuffer = [2 * maxVertices]float32{}
	tesselator.colorBuffer = [3 * maxVertices]float32{}
	tesselator.vertices = 0
	tesselator.hasColor = false
	tesselator.hasTexture = false

	return tesselator
}

func (tesselator *Tesselator) Flush() {
	gl.VertexPointer(3, gl.FLOAT, 0, unsafe.Pointer(&tesselator.vertexBuffer[0]))
	if tesselator.hasTexture {
		gl.TexCoordPointer(2, gl.FLOAT, 0, unsafe.Pointer(&tesselator.texCoordBuffer[0]))
	}
	if tesselator.hasColor {
		gl.ColorPointer(3, gl.FLOAT, 0, unsafe.Pointer(&tesselator.colorBuffer[0]))
	}

	gl.EnableClientState(gl.VERTEX_ARRAY)
	if tesselator.hasTexture {
		gl.EnableClientState(gl.TEXTURE_COORD_ARRAY)
	}
	if tesselator.hasColor {
		gl.EnableClientState(gl.COLOR_ARRAY)
	}

	gl.DrawArrays(gl.QUADS, 0, tesselator.vertices)

	gl.DisableClientState(gl.VERTEX_ARRAY)
	if tesselator.hasTexture {
		gl.DisableClientState(gl.TEXTURE_COORD_ARRAY)
	}
	if tesselator.hasColor {
		gl.DisableClientState(gl.COLOR_ARRAY)
	}

	tesselator.clear()
}

func (tesselator *Tesselator) clear() {
	tesselator.vertices = 0
}

func (tesselator *Tesselator) Init() {
	tesselator.clear()
	tesselator.hasColor = false
	tesselator.hasTexture = false
}

func (tesselator *Tesselator) Tex(u, v float32) {
	tesselator.hasTexture = true
	tesselator.u = u
	tesselator.v = v
}

func (tesselator *Tesselator) Color(r, g, b float32) {
	tesselator.hasColor = true
	tesselator.r = r
	tesselator.g = g
	tesselator.b = b
}

func (tesselator *Tesselator) Vertex(x, y, z float32) {
	tesselator.vertexBuffer[tesselator.vertices*3+0] = x
	tesselator.vertexBuffer[tesselator.vertices*3+1] = y
	tesselator.vertexBuffer[tesselator.vertices*3+2] = z
	if tesselator.hasTexture {
		tesselator.texCoordBuffer[tesselator.vertices*2+0] = tesselator.u
		tesselator.texCoordBuffer[tesselator.vertices*2+1] = tesselator.v
	}
	if tesselator.hasColor {
		tesselator.colorBuffer[tesselator.vertices*3+0] = tesselator.r
		tesselator.colorBuffer[tesselator.vertices*3+1] = tesselator.g
		tesselator.colorBuffer[tesselator.vertices*3+2] = tesselator.b
	}

	tesselator.vertices++
	if tesselator.vertices == maxVertices {
		tesselator.Flush()
	}
}
