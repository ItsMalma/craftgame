package renderer

import (
	"github.com/ItsMalma/craftgame/game/world"
	"github.com/ItsMalma/craftgame/phys"
	"github.com/go-gl/gl/v2.1/gl"
)

var chunkTesselator = NewTesselator()

var RebuiltChunkThisFrame = 0
var ChunkUpdates = 0

type Chunk struct {
	AABB phys.AABB

	w *world.World

	x0, y0, z0, x1, y1, z1 int

	dirty bool

	lists uint32
}

func NewChunk(w *world.World, x0, y0, z0, x1, y1, z1 int) *Chunk {
	chunk := new(Chunk)
	chunk.dirty = true
	chunk.w = w
	chunk.x0 = x0
	chunk.y0 = y0
	chunk.z0 = z0
	chunk.x1 = x1
	chunk.y1 = y1
	chunk.z1 = z1
	chunk.AABB = phys.NewAABB(float32(x0), float32(y0), float32(z0), float32(x1), float32(y1), float32(z1))
	chunk.lists = gl.GenLists(2)

	return chunk
}

func (chunk *Chunk) Rebuild(layer, texture uint32) {
	if RebuiltChunkThisFrame != 2 {
		chunk.dirty = false

		ChunkUpdates++
		RebuiltChunkThisFrame++

		gl.NewList(chunk.lists+layer, gl.COMPILE)
		gl.Enable(gl.TEXTURE_2D)
		gl.BindTexture(gl.TEXTURE_2D, texture)

		chunkTesselator.Init()

		tiles := 0
		for x := chunk.x0; x < chunk.x1; x++ {
			for y := chunk.y0; y < chunk.y1; y++ {
				for z := chunk.z0; z < chunk.z1; z++ {
					if chunk.w.IsTile(x, y, z) {
						tex := y != chunk.w.Depth()*2/3

						tiles++

						if !tex {
							Stone.Render(chunkTesselator, chunk.w, layer, x, y, z)
						} else {
							Grass.Render(chunkTesselator, chunk.w, layer, x, y, z)
						}
					}
				}
			}
		}

		chunkTesselator.Flush()
		gl.Disable(gl.TEXTURE_2D)
		gl.EndList()
	}
}

func (chunk *Chunk) Render(layer uint32, texture uint32) {
	if chunk.dirty {
		chunk.Rebuild(0, texture)
		chunk.Rebuild(1, texture)
	}

	gl.CallList(chunk.lists + layer)
}

func (chunk *Chunk) SetDirty() {
	chunk.dirty = true
}
