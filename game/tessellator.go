package game

import "craftgame/pkg/gl"

const (
	tessellatorMaxVertices = 100000
)

type Tessellator struct {
	vertex            []float32
	textureCoordinate []float32
	color             []float32

	vertices int

	hasTexture bool
	textureU   float32
	textureV   float32

	hasColor bool
	colorR   float32
	colorG   float32
	colorB   float32
}

func NewTessellator() *Tessellator {
	tessellator := new(Tessellator)

	tessellator.vertex = make([]float32, tessellatorMaxVertices*3)
	tessellator.textureCoordinate = make([]float32, tessellatorMaxVertices*2)
	tessellator.color = make([]float32, tessellatorMaxVertices*3)
	tessellator.vertices = 0

	return tessellator
}

func (tessellator *Tessellator) Init() {
	tessellator.clear()
}

func (tessellator *Tessellator) Vertex(x, y, z float32) {
	tessellator.vertex[tessellator.vertices*3] = x
	tessellator.vertex[tessellator.vertices*3+1] = y
	tessellator.vertex[tessellator.vertices*3+2] = z

	if tessellator.hasTexture {
		tessellator.textureCoordinate[tessellator.vertices*2] = tessellator.textureU
		tessellator.textureCoordinate[tessellator.vertices*2+1] = tessellator.textureV
	}

	if tessellator.hasColor {
		tessellator.color[tessellator.vertices*3] = tessellator.colorR
		tessellator.color[tessellator.vertices*3+1] = tessellator.colorG
		tessellator.color[tessellator.vertices*3+2] = tessellator.colorB
	}

	tessellator.vertices++

	if tessellator.vertices == tessellatorMaxVertices {
		tessellator.Flush()
	}
}

func (tessellator *Tessellator) Texture(u, v float32) {
	tessellator.hasTexture = true
	tessellator.textureU = u
	tessellator.textureV = v
}

func (tessellator *Tessellator) Color(r, g, b float32) {
	tessellator.hasColor = true
	tessellator.colorR = r
	tessellator.colorG = g
	tessellator.colorB = b
}

func (tessellator *Tessellator) Flush() {
	gl.VertexPointer(3, gl.Float, 0, &tessellator.vertex[0])
	if tessellator.hasTexture {
		gl.TexCoordPointer(2, gl.Float, 0, &tessellator.textureCoordinate[0])
	}
	if tessellator.hasColor {
		gl.ColorPointer(3, gl.Float, 0, &tessellator.color[0])
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
	tessellator.hasTexture = false
	tessellator.hasColor = false
}
