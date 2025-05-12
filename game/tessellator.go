package game

import "craftgame/pkg/gl"

const (
	tessellatorMaxVertices = 100000
)

type Tessellator struct {
	vertexBuffer   []float32
	texCoordBuffer []float32
	colorBuffer    []float32

	vertices int

	u, v    float32
	r, g, b float32

	hasColor   bool
	hasTexture bool
}

func NewTessellator() *Tessellator {
	tessellator := new(Tessellator)

	tessellator.vertexBuffer = make([]float32, tessellatorMaxVertices*3)
	tessellator.texCoordBuffer = make([]float32, tessellatorMaxVertices*2)
	tessellator.colorBuffer = make([]float32, tessellatorMaxVertices*3)
	tessellator.vertices = 0
	tessellator.hasColor = false
	tessellator.hasTexture = false

	return tessellator
}

func (tessellator *Tessellator) Init() {
	tessellator.clear()
	tessellator.hasColor = false
	tessellator.hasTexture = false
}

func (tessellator *Tessellator) Vertex(x, y, z float32) {
	tessellator.vertexBuffer[tessellator.vertices*3+0] = x
	tessellator.vertexBuffer[tessellator.vertices*3+1] = y
	tessellator.vertexBuffer[tessellator.vertices*3+2] = z

	if tessellator.hasTexture {
		tessellator.texCoordBuffer[tessellator.vertices*2+0] = tessellator.u
		tessellator.texCoordBuffer[tessellator.vertices*2+1] = tessellator.v
	}

	if tessellator.hasColor {
		tessellator.colorBuffer[tessellator.vertices*3+0] = tessellator.r
		tessellator.colorBuffer[tessellator.vertices*3+1] = tessellator.g
		tessellator.colorBuffer[tessellator.vertices*3+2] = tessellator.b
	}

	tessellator.vertices++

	if tessellator.vertices == tessellatorMaxVertices {
		tessellator.Flush()
	}
}

func (tessellator *Tessellator) Tex(u, v float32) {
	tessellator.hasTexture = true
	tessellator.u = u
	tessellator.v = v
}

func (tessellator *Tessellator) Color(r, g, b float32) {
	tessellator.hasColor = true
	tessellator.r = r
	tessellator.g = g
	tessellator.b = b
}

func (tessellator *Tessellator) Flush() {
	gl.VertexPointer(3, gl.Float, 0, &tessellator.vertexBuffer[0])
	if tessellator.hasTexture {
		gl.TexCoordPointer(2, gl.Float, 0, &tessellator.texCoordBuffer[0])
	}
	if tessellator.hasColor {
		gl.ColorPointer(3, gl.Float, 0, &tessellator.colorBuffer[0])
	}

	gl.EnableClientState(gl.VertexArray)
	if tessellator.hasTexture {
		gl.EnableClientState(gl.TexCoordArray)
	}
	if tessellator.hasColor {
		gl.EnableClientState(gl.ColorArray)
	}

	gl.DrawArrays(gl.Quads, 0, tessellator.vertices)

	gl.DisableClientState(gl.VertexArray)
	if tessellator.hasTexture {
		gl.DisableClientState(gl.TexCoordArray)
	}
	if tessellator.hasColor {
		gl.DisableClientState(gl.ColorArray)
	}

	tessellator.clear()
}

func (tessellator *Tessellator) clear() {
	tessellator.vertices = 0
}
