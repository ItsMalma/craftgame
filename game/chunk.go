package game

import (
	"craftgame/pkg/gl"
)

var (
	ChunkRebuiltThisFrame = 0
	ChunkUpdates          = 0

	ChunkTessellator *Tessellator = NewTessellator()
)

type Chunk struct {
	level *Level

	BoundingBox                        AABB
	minX, minY, minZ, maxX, maxY, maxZ int

	lists int
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

func (chunk *Chunk) Rebuild(layer int, texture int32) {
	if ChunkRebuiltThisFrame == 2 {
		return
	}

	ChunkUpdates++
	ChunkRebuiltThisFrame++

	chunk.dirty = false

	gl.NewList(chunk.lists+layer, gl.Compile)
	gl.Enable(gl.Texture2D)
	gl.BindTexture(gl.Texture2D, texture)
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
	gl.Disable(gl.Texture2D)
	gl.EndList()
}

func (chunk *Chunk) Render(layer int, texture int32) {
	if chunk.dirty {
		chunk.Rebuild(0, texture)
		chunk.Rebuild(1, texture)
	}

	gl.CallList(chunk.lists + layer)
}
