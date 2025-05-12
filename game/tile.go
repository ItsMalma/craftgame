package game

var (
	TileStone = NewTile(0)
	TileGrass = NewTile(1)
)

type Tile struct {
	tex int
}

func NewTile(tex int) *Tile {
	tile := new(Tile)

	tile.tex = tex

	return tile
}

func (tile *Tile) Render(tessellator *Tessellator, level *Level, layer, x, y, z int) {
	var (
		u0 float32 = float32(tile.tex) / 16.0
		u1 float32 = u0 + 0.0624375
		v0 float32 = 0.0
		v1 float32 = v0 + 0.0624375

		c1 float32 = 1.0
		c2 float32 = 0.8
		c3 float32 = 0.6

		x0 float32 = float32(x) + 0.0
		x1 float32 = float32(x) + 1.0
		y0 float32 = float32(y) + 0.0
		y1 float32 = float32(y) + 1.0
		z0 float32 = float32(z) + 0.0
		z1 float32 = float32(z) + 1.0

		br float32
	)

	if !level.IsSolidTile(x, y-1, z) {
		br = level.GetBrightness(x, y-1, z) * c1
		if (br == c1) != (layer == 1) {
			tessellator.Color(br, br, br)
			tessellator.Tex(u0, v1)
			tessellator.Vertex(x0, y0, z1)
			tessellator.Tex(u0, v0)
			tessellator.Vertex(x0, y0, z0)
			tessellator.Tex(u1, v0)
			tessellator.Vertex(x1, y0, z0)
			tessellator.Tex(u1, v1)
			tessellator.Vertex(x1, y0, z1)
		}
	}

	if !level.IsSolidTile(x, y+1, z) {
		br = level.GetBrightness(x, y, z) * c1
		if (br == c1) != (layer == 1) {
			tessellator.Color(br, br, br)
			tessellator.Tex(u1, v1)
			tessellator.Vertex(x1, y1, z1)
			tessellator.Tex(u1, v0)
			tessellator.Vertex(x1, y1, z0)
			tessellator.Tex(u0, v0)
			tessellator.Vertex(x0, y1, z0)
			tessellator.Tex(u0, v1)
			tessellator.Vertex(x0, y1, z1)
		}
	}

	if !level.IsSolidTile(x, y, z-1) {
		br = level.GetBrightness(x, y, z-1) * c2
		if (br == c2) != (layer == 1) {
			tessellator.Color(br, br, br)
			tessellator.Tex(u1, v0)
			tessellator.Vertex(x0, y1, z0)
			tessellator.Tex(u0, v0)
			tessellator.Vertex(x1, y1, z0)
			tessellator.Tex(u0, v1)
			tessellator.Vertex(x1, y0, z0)
			tessellator.Tex(u1, v1)
			tessellator.Vertex(x0, y0, z0)
		}
	}

	if !level.IsSolidTile(x, y, z+1) {
		br = level.GetBrightness(x, y, z+1) * c2
		if (br == c2) != (layer == 1) {
			tessellator.Color(br, br, br)
			tessellator.Tex(u0, v0)
			tessellator.Vertex(x0, y1, z1)
			tessellator.Tex(u0, v1)
			tessellator.Vertex(x0, y0, z1)
			tessellator.Tex(u1, v1)
			tessellator.Vertex(x1, y0, z1)
			tessellator.Tex(u1, v0)
			tessellator.Vertex(x1, y1, z1)
		}
	}

	if !level.IsSolidTile(x-1, y, z) {
		br = level.GetBrightness(x-1, y, z) * c3
		if (br == c3) != (layer == 1) {
			tessellator.Color(br, br, br)
			tessellator.Tex(u1, v0)
			tessellator.Vertex(x0, y1, z1)
			tessellator.Tex(u0, v0)
			tessellator.Vertex(x0, y1, z0)
			tessellator.Tex(u0, v1)
			tessellator.Vertex(x0, y0, z0)
			tessellator.Tex(u1, v1)
			tessellator.Vertex(x0, y0, z1)
		}
	}

	if !level.IsSolidTile(x+1, y, z) {
		br = level.GetBrightness(x+1, y, z) * c3
		if (br == c3) != (layer == 1) {
			tessellator.Color(br, br, br)
			tessellator.Tex(u0, v1)
			tessellator.Vertex(x1, y0, z1)
			tessellator.Tex(u1, v1)
			tessellator.Vertex(x1, y0, z0)
			tessellator.Tex(u1, v0)
			tessellator.Vertex(x1, y1, z0)
			tessellator.Tex(u0, v0)
			tessellator.Vertex(x1, y1, z1)
		}
	}
}

func (tile *Tile) RenderFace(tessellator *Tessellator, x, y, z, face int) {
	var (
		x0 float32 = float32(x)
		x1 float32 = float32(x) + 1.0
		y0 float32 = float32(y)
		y1 float32 = float32(y) + 1.0
		z0 float32 = float32(z)
		z1 float32 = float32(z) + 1.0
	)

	if face == 0 {
		tessellator.Vertex(x0, y0, z1)
		tessellator.Vertex(x0, y0, z0)
		tessellator.Vertex(x1, y0, z0)
		tessellator.Vertex(x1, y0, z1)
	}
	if face == 1 {
		tessellator.Vertex(x1, y1, z1)
		tessellator.Vertex(x1, y1, z0)
		tessellator.Vertex(x0, y1, z0)
		tessellator.Vertex(x0, y1, z1)
	}
	if face == 2 {
		tessellator.Vertex(x0, y1, z0)
		tessellator.Vertex(x1, y1, z0)
		tessellator.Vertex(x1, y0, z0)
		tessellator.Vertex(x0, y0, z0)
	}
	if face == 3 {
		tessellator.Vertex(x0, y1, z1)
		tessellator.Vertex(x0, y0, z1)
		tessellator.Vertex(x1, y0, z1)
		tessellator.Vertex(x1, y1, z1)
	}
	if face == 4 {
		tessellator.Vertex(x0, y1, z1)
		tessellator.Vertex(x0, y1, z0)
		tessellator.Vertex(x0, y0, z0)
		tessellator.Vertex(x0, y0, z1)
	}
	if face == 5 {
		tessellator.Vertex(x1, y0, z1)
		tessellator.Vertex(x1, y0, z0)
		tessellator.Vertex(x1, y1, z0)
		tessellator.Vertex(x1, y1, z1)
	}
}
