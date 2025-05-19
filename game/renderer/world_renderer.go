package renderer

import (
	"math"
	"time"

	"github.com/ItsMalma/craftgame/game/entity"
	"github.com/ItsMalma/craftgame/game/world"
	"github.com/go-gl/gl/v2.1/gl"
)

const chunkSize = 16

type WorldRenderer struct {
	w                         *world.World
	chunks                    []*Chunk
	xChunks, yChunks, zChunks int
	t                         *Tesselator

	terrainTexture uint32
}

func NewWorldRenderer(w *world.World, terrainTexture uint32) *WorldRenderer {
	worldRenderer := new(WorldRenderer)
	worldRenderer.t = NewTesselator()
	worldRenderer.w = w
	worldRenderer.terrainTexture = terrainTexture
	w.AddListener(worldRenderer)
	worldRenderer.xChunks = w.Width() / chunkSize
	worldRenderer.yChunks = w.Depth() / chunkSize
	worldRenderer.zChunks = w.Height() / chunkSize
	worldRenderer.chunks = make([]*Chunk, worldRenderer.xChunks*worldRenderer.yChunks*worldRenderer.zChunks)

	for x := range worldRenderer.xChunks {
		for y := range worldRenderer.yChunks {
			for z := range worldRenderer.zChunks {
				x0 := x * chunkSize
				y0 := y * chunkSize
				z0 := z * chunkSize

				x1 := (x + 1) * chunkSize
				y1 := (y + 1) * chunkSize
				z1 := (z + 1) * chunkSize

				if wWidth := w.Width(); x1 > wWidth {
					x1 = wWidth
				}
				if wDepth := w.Depth(); y1 > wDepth {
					y1 = wDepth
				}
				if wHeight := w.Height(); z1 > wHeight {
					z1 = wHeight
				}

				worldRenderer.chunks[(x+y*worldRenderer.xChunks)*worldRenderer.zChunks+z] = NewChunk(w, x0, y0, z0, x1, y1, z1)
			}
		}
	}

	return worldRenderer
}

func (worldRenderer *WorldRenderer) Render(player *entity.Player, layer uint32) {
	RebuiltChunkThisFrame = 0
	frustum := GetFrustum()

	for _, chunk := range worldRenderer.chunks {
		if frustum.AABBInFrustum(chunk.AABB) {
			chunk.Render(layer, worldRenderer.terrainTexture)
		}
	}
}

func (worldRenderer *WorldRenderer) Pick(player *entity.Player) {
	r := float32(3.0)
	box := player.BB.Grow(r, r, r)

	x0 := int(box.X0)
	x1 := int(box.X1 + 1.0)
	y0 := int(box.Y0)
	y1 := int(box.Y1 + 1.0)
	z0 := int(box.Z0)
	z1 := int(box.Z1 + 1.0)

	gl.InitNames()
	for x := x0; x < x1; x++ {
		gl.PushName(uint32(x))
		for y := y0; y < y1; y++ {
			gl.PushName(uint32(y))
			for z := z0; z < z1; z++ {
				gl.PushName(uint32(z))
				if worldRenderer.w.IsSolidTile(x, y, z) {
					gl.PushName(0)
					for i := range 6 {
						gl.PushName(uint32(i))
						worldRenderer.t.Init()
						Stone.RenderFace(worldRenderer.t, x, y, z, i)
						worldRenderer.t.Flush()
						gl.PopName()
					}
					gl.PopName()
				}
				gl.PopName()
			}
			gl.PopName()
		}
		gl.PopName()
	}
}

func (worldRenderer *WorldRenderer) RenderHit(h HitResult) {
	gl.Enable(gl.BLEND)
	gl.BlendFunc(gl.SRC_ALPHA, gl.CURRENT_BIT)
	gl.Color4f(1.0, 1.0, 1.0, float32(math.Sin(float64(time.Now().UnixMilli())/100.0)*0.2+0.4))

	worldRenderer.t.Init()
	Stone.RenderFace(worldRenderer.t, h.X, h.Y, h.Z, h.F)
	worldRenderer.t.Flush()

	gl.Disable(gl.BLEND)
}

func (worldRenderer *WorldRenderer) SetDirty(x0, y0, z0, x1, y1, z1 int) {
	x0 /= chunkSize
	x1 /= chunkSize
	y0 /= chunkSize
	y1 /= chunkSize
	z0 /= chunkSize
	z1 /= chunkSize

	if x0 < 0 {
		x0 = 0
	}
	if y0 < 0 {
		y0 = 0
	}
	if z0 < 0 {
		z0 = 0
	}

	if x1 >= worldRenderer.xChunks {
		x1 = worldRenderer.xChunks - 1
	}
	if y1 >= worldRenderer.yChunks {
		y1 = worldRenderer.yChunks - 1
	}
	if z1 >= worldRenderer.zChunks {
		z1 = worldRenderer.zChunks - 1
	}

	for x := x0; x <= x1; x++ {
		for y := y0; y <= y1; y++ {
			for z := z0; z <= z1; z++ {
				worldRenderer.chunks[(x+y*worldRenderer.xChunks)*worldRenderer.zChunks+z].SetDirty()
			}
		}
	}
}

func (worldRenderer *WorldRenderer) TileChanged(x, y, z int) {
	worldRenderer.SetDirty(x-1, y-1, z-1, x+1, y+1, z+1)
}

func (worldRenderer *WorldRenderer) LightColumnChanged(x, z, y0, y1 int) {
	worldRenderer.SetDirty(x-1, y0-1, z-1, x+1, y1+1, z+1)
}

func (worldRenderer *WorldRenderer) AllChanged() {
	worldRenderer.SetDirty(0, 0, 0, worldRenderer.w.Width(), worldRenderer.w.Depth(), worldRenderer.w.Height())
}
