package renderer

import "github.com/ItsMalma/craftgame/game/world"

var Stone = newTile(0)
var Grass = newTile(1)

type Tile struct {
	texture int
}

func newTile(texture int) *Tile {
	tile := new(Tile)
	tile.texture = texture

	return tile
}

func (tile *Tile) Render(t *Tesselator, w *world.World, layer uint32, x, y, z int) {
	u0 := float32(tile.texture) / 16.0
	u1 := u0 + 0.0624375
	v0 := float32(0.0)
	v1 := v0 + 0.0624375
	c1 := float32(1.0)
	c2 := float32(0.8)
	c3 := float32(0.6)
	x0 := float32(x) + 0.0
	x1 := float32(x) + 1.0
	y0 := float32(y) + 0.0
	y1 := float32(y) + 1.0
	z0 := float32(z) + 0.0
	z1 := float32(z) + 1.0
	var br float32

	if !w.IsSolidTile(x, y-1, z) {
		br = w.GetBrightness(x, y-1, z) * c1
		if (br == c1) != (layer == 1) {
			t.Color(br, br, br)
			t.Tex(u0, v1)
			t.Vertex(x0, y0, z1)
			t.Tex(u0, v0)
			t.Vertex(x0, y0, z0)
			t.Tex(u1, v0)
			t.Vertex(x1, y0, z0)
			t.Tex(u1, v1)
			t.Vertex(x1, y0, z1)
		}
	}

	if !w.IsSolidTile(x, y+1, z) {
		br = w.GetBrightness(x, y, z) * c1
		if (br == c1) != (layer == 1) {
			t.Color(br, br, br)
			t.Tex(u1, v1)
			t.Vertex(x1, y1, z1)
			t.Tex(u1, v0)
			t.Vertex(x1, y1, z0)
			t.Tex(u0, v0)
			t.Vertex(x0, y1, z0)
			t.Tex(u0, v1)
			t.Vertex(x0, y1, z1)
		}
	}

	if !w.IsSolidTile(x, y, z-1) {
		br = w.GetBrightness(x, y, z-1) * c2
		if (br == c2) != (layer == 1) {
			t.Color(br, br, br)
			t.Tex(u1, v0)
			t.Vertex(x0, y1, z0)
			t.Tex(u0, v0)
			t.Vertex(x1, y1, z0)
			t.Tex(u0, v1)
			t.Vertex(x1, y0, z0)
			t.Tex(u1, v1)
			t.Vertex(x0, y0, z0)
		}
	}

	if !w.IsSolidTile(x, y, z+1) {
		br = w.GetBrightness(x, y, z+1) * c2
		if (br == c2) != (layer == 1) {
			t.Color(br, br, br)
			t.Tex(u0, v0)
			t.Vertex(x0, y1, z1)
			t.Tex(u0, v1)
			t.Vertex(x0, y0, z1)
			t.Tex(u1, v1)
			t.Vertex(x1, y0, z1)
			t.Tex(u1, v0)
			t.Vertex(x1, y1, z1)
		}
	}

	if !w.IsSolidTile(x-1, y, z) {
		br = w.GetBrightness(x-1, y, z) * c3
		if (br == c3) != (layer == 1) {
			t.Color(br, br, br)
			t.Tex(u1, v0)
			t.Vertex(x0, y1, z1)
			t.Tex(u0, v0)
			t.Vertex(x0, y1, z0)
			t.Tex(u0, v1)
			t.Vertex(x0, y0, z0)
			t.Tex(u1, v1)
			t.Vertex(x0, y0, z1)
		}
	}

	if !w.IsSolidTile(x+1, y, z) {
		br = w.GetBrightness(x+1, y, z) * c3
		if (br == c3) != (layer == 1) {
			t.Color(br, br, br)
			t.Tex(u0, v1)
			t.Vertex(x1, y0, z1)
			t.Tex(u1, v1)
			t.Vertex(x1, y0, z0)
			t.Tex(u1, v0)
			t.Vertex(x1, y1, z0)
			t.Tex(u0, v0)
			t.Vertex(x1, y1, z1)
		}
	}
}

func (tile *Tile) RenderFace(t *Tesselator, x, y, z, face int) {
	x0 := float32(x + 0.0)
	x1 := float32(x + 1.0)
	y0 := float32(y + 0.0)
	y1 := float32(y + 1.0)
	z0 := float32(z + 0.0)
	z1 := float32(z + 1.0)

	if face == 0 {
		t.Vertex(x0, y0, z1)
		t.Vertex(x0, y0, z0)
		t.Vertex(x1, y0, z0)
		t.Vertex(x1, y0, z1)
	}

	if face == 1 {
		t.Vertex(x1, y1, z1)
		t.Vertex(x1, y1, z0)
		t.Vertex(x0, y1, z0)
		t.Vertex(x0, y1, z1)
	}

	if face == 2 {
		t.Vertex(x0, y1, z0)
		t.Vertex(x1, y1, z0)
		t.Vertex(x1, y0, z0)
		t.Vertex(x0, y0, z0)
	}

	if face == 3 {
		t.Vertex(x0, y1, z1)
		t.Vertex(x0, y0, z1)
		t.Vertex(x1, y0, z1)
		t.Vertex(x1, y1, z1)
	}

	if face == 4 {
		t.Vertex(x0, y1, z1)
		t.Vertex(x0, y1, z0)
		t.Vertex(x0, y0, z0)
		t.Vertex(x0, y0, z1)
	}

	if face == 5 {
		t.Vertex(x1, y0, z1)
		t.Vertex(x1, y0, z0)
		t.Vertex(x1, y1, z0)
		t.Vertex(x1, y1, z1)
	}
}
