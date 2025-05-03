package game

const (
	chunkSize = 8
)

type Renderer struct {
	chunkAmountX, chunkAmountY, chunkAmountZ int

	chunks []*Chunk

	texture int32
}

func NewRenderer(level *Level, texture int32) *Renderer {
	renderer := new(Renderer)

	renderer.chunkAmountX = level.Width / chunkSize
	renderer.chunkAmountY = level.Depth / chunkSize
	renderer.chunkAmountZ = level.Height / chunkSize

	renderer.chunks = make([]*Chunk, renderer.chunkAmountX*renderer.chunkAmountY*renderer.chunkAmountZ)

	for x := range renderer.chunkAmountX {
		for y := range renderer.chunkAmountY {
			for z := range renderer.chunkAmountZ {
				minChunkX := x * chunkSize
				minChunkY := y * chunkSize
				minChunkZ := z * chunkSize

				maxChunkX := (x + 1) * chunkSize
				maxChunkY := (y + 1) * chunkSize
				maxChunkZ := (z + 1) * chunkSize

				maxChunkX = min(maxChunkX, level.Width)
				maxChunkY = min(maxChunkY, level.Depth)
				maxChunkZ = min(maxChunkZ, level.Height)

				chunk := NewChunk(level, minChunkX, minChunkY, minChunkZ, maxChunkX, maxChunkY, maxChunkZ)
				renderer.chunks[(x+y*renderer.chunkAmountX)*renderer.chunkAmountZ+z] = chunk
			}
		}
	}

	renderer.texture = texture

	return renderer
}

func (renderer *Renderer) Render(layer int) {
	frustum := GetFrustum()

	ChunkRebuiltThisFrame = 0

	for _, chunk := range renderer.chunks {
		if frustum.CubeInFrustumAABB(chunk.BoundingBox) {
			chunk.Render(layer, renderer.texture)
		}
	}
}
