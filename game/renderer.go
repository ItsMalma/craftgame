package game

import (
	"craftgame/pkg/gl"
	"math"
	"time"
)

const (
	chunkSize = 16
)

type Renderer struct {
	level  *Level
	chunks []*Chunk

	xChunks, yChunks, zChunks int

	Tessellator *Tessellator
}

func NewRenderer(level *Level) *Renderer {
	renderer := new(Renderer)

	renderer.Tessellator = NewTessellator()

	renderer.level = level

	level.AddListener(renderer)

	renderer.xChunks = level.Width / chunkSize
	renderer.yChunks = level.Depth / chunkSize
	renderer.zChunks = level.Height / chunkSize

	renderer.chunks = make([]*Chunk, renderer.xChunks*renderer.yChunks*renderer.zChunks)

	for x := range renderer.xChunks {
		for y := range renderer.yChunks {
			for z := range renderer.zChunks {
				x0 := x * chunkSize
				y0 := y * chunkSize
				z0 := z * chunkSize

				x1 := (x + 1) * chunkSize
				y1 := (y + 1) * chunkSize
				z1 := (z + 1) * chunkSize

				if x1 > level.Width {
					x1 = level.Width
				}
				if y1 > level.Depth {
					y1 = level.Depth
				}
				if z1 > level.Height {
					z1 = level.Height
				}

				chunk := NewChunk(level, x0, y0, z0, x1, y1, z1)
				renderer.chunks[(x+y*renderer.xChunks)*renderer.zChunks+z] = chunk
			}
		}
	}

	return renderer
}

func (renderer *Renderer) Render(player *Player, layer int) error {
	ChunkRebuiltThisFrame = 0

	frustum := GetFrustum()

	for _, chunk := range renderer.chunks {
		if frustum.CubeInFrustumAABB(chunk.BoundingBox) {
			if err := chunk.Render(layer); err != nil {
				return err
			}
		}
	}

	return nil
}

func (renderer *Renderer) SetDirty(x0, y0, z0, x1, y1, z1 int) {
	x0 /= chunkSize
	y0 /= chunkSize
	z0 /= chunkSize
	x1 /= chunkSize
	y1 /= chunkSize
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

	if x1 >= renderer.xChunks {
		x1 = renderer.xChunks - 1
	}
	if y1 >= renderer.yChunks {
		y1 = renderer.yChunks - 1
	}
	if z1 >= renderer.zChunks {
		z1 = renderer.zChunks - 1
	}

	for x := x0; x <= x1; x++ {
		for y := y0; y <= y1; y++ {
			for z := z0; z <= z1; z++ {
				renderer.chunks[(x+y*renderer.xChunks)*renderer.zChunks+z].SetDirty()
			}
		}
	}
}

func (renderer *Renderer) Pick(player *Player) {
	radius := float32(3.0)
	boundingBox := player.Entity.BoundingBox.Grow(radius, radius, radius)

	x0 := int(boundingBox.MinX)
	x1 := int(boundingBox.MaxX + 1)
	y0 := int(boundingBox.MinY)
	y1 := int(boundingBox.MaxY + 1)
	z0 := int(boundingBox.MinZ)
	z1 := int(boundingBox.MaxZ + 1)

	gl.InitNames()
	for x := x0; x < x1; x++ {
		gl.PushName(x)

		for y := y0; y < y1; y++ {
			gl.PushName(y)

			for z := z0; z < z1; z++ {
				gl.PushName(z)

				if renderer.level.IsSolidTile(x, y, z) {
					gl.PushName(0)

					for face := range 6 {
						gl.PushName(face)

						renderer.Tessellator.Init()
						TileStone.RenderFace(renderer.Tessellator, x, y, z, face)
						renderer.Tessellator.Flush()

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

func (renderer *Renderer) RenderHit(hitResult *HitResult) {
	gl.Enable(gl.Blend)
	gl.BlendFunc(gl.SrcAlpha, gl.CurrentBit)
	gl.Color4f(1.0, 1.0, 1.0, float32(math.Sin(float64(time.Now().UnixMilli())/100.0)*0.2+0.4))

	renderer.Tessellator.Init()
	TileStone.RenderFace(renderer.Tessellator, hitResult.X, hitResult.Y, hitResult.Z, hitResult.Face)
	renderer.Tessellator.Flush()

	gl.Disable(gl.Blend)
}

func (renderer *Renderer) LightColumnChanged(x, z, y0, y1 int) {
	renderer.SetDirty(x-1, y0-1, z-1, x+1, y1+1, z+1)
}

func (renderer *Renderer) TileChanged(x, y, z int) {
	renderer.SetDirty(x-1, y-1, z-1, x+1, y+1, z+1)
}

func (renderer *Renderer) AllChanged() {
	renderer.SetDirty(0, 0, 0, renderer.level.Width, renderer.level.Depth, renderer.level.Height)
}
