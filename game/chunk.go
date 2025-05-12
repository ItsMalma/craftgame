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
	BoundingBox AABB
	Level       *Level

	MinX, MinY, MinZ, MaxX, MaxY, MaxZ int

	dirty bool
	lists int
}

func NewChunk(level *Level, minX, minY, minZ, maxX, maxY, maxZ int) *Chunk {
	chunk := new(Chunk)

	chunk.dirty = true
	chunk.lists = -1

	chunk.Level = level

	chunk.MinX = minX
	chunk.MinY = minY
	chunk.MinZ = minZ
	chunk.MaxX = maxX
	chunk.MaxY = maxY
	chunk.MaxZ = maxZ

	chunk.BoundingBox = NewAABB(float32(minX), float32(minY), float32(minZ), float32(maxX), float32(maxY), float32(maxZ))

	chunk.lists = gl.GenLists(2)

	return chunk
}

func (chunk *Chunk) Rebuild(layer int) error {
	if ChunkRebuiltThisFrame == 2 {
		return nil
	}

	chunk.dirty = false

	ChunkUpdates++
	ChunkRebuiltThisFrame++

	terrain, err := LoadTexture("terrain.png", gl.Nearest)
	if err != nil {
		return err
	}

	gl.NewList(chunk.lists+layer, gl.Compile)
	gl.Enable(gl.Texture2D)
	gl.BindTexture(gl.Texture2D, terrain)

	chunkTessellator.Init()

	tiles := 0
	for x := chunk.MinX; x < chunk.MaxX; x++ {
		for y := chunk.MinY; y < chunk.MaxY; y++ {
			for z := chunk.MinZ; z < chunk.MaxZ; z++ {
				if chunk.Level.IsTile(x, y, z) {
					tex := y != chunk.Level.Depth*2/3

					tiles++

					if !tex {
						TileStone.Render(chunkTessellator, chunk.Level, int(layer), x, y, z)
					} else {
						TileGrass.Render(chunkTessellator, chunk.Level, int(layer), x, y, z)
					}
				}
			}
		}
	}

	chunkTessellator.Flush()

	gl.Disable(gl.Texture2D)
	gl.EndList()

	return nil
}

func (chunk *Chunk) Render(layer int) error {
	if chunk.dirty {
		if err := chunk.Rebuild(0); err != nil {
			return err
		}
		if err := chunk.Rebuild(1); err != nil {
			return err
		}
	}

	gl.CallList(chunk.lists + layer)

	return nil
}

func (chunk *Chunk) SetDirty() {
	chunk.dirty = true
}
