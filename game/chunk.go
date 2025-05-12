package game

import (
	"craftgame/pkg/gl"
)

var (
	ChunkRebuiltThisFrame = 0
	ChunkUpdates          = 0

	chunkTessellator *Tessellator = NewTessellator()
)

type Chunk struct {
	AABB  AABB
	level *Level

	x0, y0, z0, x1, y1, z1 int

	dirty bool
	lists int
}

func NewChunk(level *Level, minX, minY, minZ, maxX, maxY, maxZ int) *Chunk {
	chunk := new(Chunk)

	chunk.dirty = true
	chunk.lists = -1

	chunk.level = level

	chunk.x0 = minX
	chunk.y0 = minY
	chunk.z0 = minZ
	chunk.x1 = maxX
	chunk.y1 = maxY
	chunk.z1 = maxZ

	chunk.AABB = NewAABB(float64(minX), float64(minY), float64(minZ), float64(maxX), float64(maxY), float64(maxZ))

	chunk.lists = gl.GenLists(2)

	return chunk
}

func (chunk *Chunk) Rebuild(layer int, texture int32) {
	if ChunkRebuiltThisFrame == 2 {
		return
	}

	chunk.dirty = false

	ChunkUpdates++
	ChunkRebuiltThisFrame++

	gl.NewList(chunk.lists+layer, gl.Compile)
	gl.Enable(gl.Texture2D)
	gl.BindTexture(gl.Texture2D, texture)

	chunkTessellator.Init()

	tiles := 0
	for x := chunk.x0; x < chunk.x1; x++ {
		for y := chunk.y0; y < chunk.y1; y++ {
			for z := chunk.z0; z < chunk.z1; z++ {
				if chunk.level.IsTile(x, y, z) {
					tex := y != chunk.level.Depth*2/3

					tiles++

					if !tex {
						TileStone.Render(chunkTessellator, chunk.level, int(layer), x, y, z)
					} else {
						TileGrass.Render(chunkTessellator, chunk.level, int(layer), x, y, z)
					}
				}
			}
		}
	}

	chunkTessellator.Flush()

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

func (chunk *Chunk) SetDirty() {
	chunk.dirty = true
}
