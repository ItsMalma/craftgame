package game

import (
	"github.com/go-gl/gl/v2.1/gl"
)

var (
	ChunkRebuiltThisFrame = 0
	ChunkUpdates          = 0

	chunkTexture       uint32
	chunkTextureLoaded = false

	ChunkTessellator *Tessellator = NewTessellator()
)

func ChunkTexture() uint32 {
	if !chunkTextureLoaded {
		chunkTexture = LoadTexture("terrain.png", gl.NEAREST)
		chunkTextureLoaded = true
	}

	return chunkTexture
}

type Chunk struct {
	level *Level

	BoundingBox                        *AABB
	minX, minY, minZ, maxX, maxY, maxZ int

	lists uint32
	dirty bool
}

func NewChunk(level *Level, minX, minY, minZ, maxX, maxY, maxZ int) *Chunk {
	chunk := new(Chunk)

	chunk.dirty = true

	chunk.level = level

	chunk.minX = minX
	chunk.minY = minY
	chunk.minZ = minZ
	chunk.maxX = maxX
	chunk.maxY = maxY
	chunk.maxZ = maxZ

	chunk.lists = gl.GenLists(2)

	chunk.BoundingBox = NewAABB(float64(minX), float64(minY), float64(minZ), float64(maxX), float64(maxY), float64(maxZ))

	return chunk
}

func (chunk *Chunk) Rebuild(layer uint32) {
	if ChunkRebuiltThisFrame == 2 {
		return
	}

	ChunkUpdates++
	ChunkRebuiltThisFrame++

	chunk.dirty = false

	gl.NewList(chunk.lists+layer, gl.COMPILE)
	gl.Enable(gl.TEXTURE_2D)
	gl.BindTexture(gl.TEXTURE_2D, ChunkTexture())
	ChunkTessellator.Init()

	for x := chunk.minX; x < chunk.maxX; x++ {
		for y := chunk.minY; y < chunk.maxY; y++ {
			for z := chunk.minZ; z < chunk.maxZ; z++ {
				if chunk.level.IsTile(x, y, z) {
					if y > chunk.level.Depth-7 && chunk.level.GetBrightness(x, y, z) == 1.0 {
						TileDirt.Render(ChunkTessellator, chunk.level, int(layer), x, y, z)
					} else {
						TileStone.Render(ChunkTessellator, chunk.level, int(layer), x, y, z)
					}
				}
			}
		}
	}

	ChunkTessellator.Flush()
	gl.Disable(gl.TEXTURE_2D)
	gl.EndList()
}

func (chunk *Chunk) Render(layer uint32) {
	if chunk.dirty {
		chunk.Rebuild(0)
		chunk.Rebuild(1)
	}

	gl.CallList(chunk.lists + layer)
}
