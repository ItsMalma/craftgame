package game

var (
	TileGrass = NewTile(0)
	TileRock  = NewTile(1)
)

type Tile struct {
	textureId int
}

func NewTile(textureId int) *Tile {
	tile := new(Tile)

	tile.textureId = textureId

	return tile
}

func (tile *Tile) Render(tessellator *Tessellator, level *Level, layer, x, y, z int) {
	var (
		minU float32 = float32(tile.textureId) / 16.0
		maxU float32 = minU + 0.0625
		minV float32 = 0.0
		maxV float32 = minV + 0.0625

		shadeX float32 = 0.6
		shadeY float32 = 1.0
		shadeZ float32 = 0.8

		minX float32 = float32(x) + 0.0
		maxX float32 = float32(x) + 1.0
		minY float32 = float32(y) + 0.0
		maxY float32 = float32(y) + 1.0
		minZ float32 = float32(z) + 0.0
		maxZ float32 = float32(z) + 1.0
	)

	if !level.IsSolidTile(x, y-1, z) {
		brightness := level.GetBrightness(x, y-1, z) * float32(shadeY)

		if (layer == 1) != (brightness == float32(shadeY)) {
			tessellator.Color(brightness, brightness, brightness)
			tessellator.Texture(minU, maxV)
			tessellator.Vertex(minX, minY, maxZ)
			tessellator.Texture(minU, minV)
			tessellator.Vertex(minX, minY, minZ)
			tessellator.Texture(maxU, minV)
			tessellator.Vertex(maxX, minY, minZ)
			tessellator.Texture(maxU, maxV)
			tessellator.Vertex(maxX, minY, maxZ)
		}
	}

	if !level.IsSolidTile(x, y+1, z) {
		brightness := level.GetBrightness(x, y+1, z) * shadeY

		if (layer == 1) != (brightness == shadeY) {
			tessellator.Color(brightness, brightness, brightness)
			tessellator.Texture(maxU, maxV)
			tessellator.Vertex(maxX, maxY, maxZ)
			tessellator.Texture(maxU, minV)
			tessellator.Vertex(maxX, maxY, minZ)
			tessellator.Texture(minU, minV)
			tessellator.Vertex(minX, maxY, minZ)
			tessellator.Texture(minU, maxV)
			tessellator.Vertex(minX, maxY, maxZ)
		}
	}

	if !level.IsSolidTile(x, y, z-1) {
		brightness := level.GetBrightness(x, y, z-1) * shadeZ

		if (layer == 1) != (brightness == shadeZ) {
			tessellator.Color(brightness, brightness, brightness)
			tessellator.Texture(maxU, minV)
			tessellator.Vertex(minX, maxY, minZ)
			tessellator.Texture(minU, minV)
			tessellator.Vertex(maxX, maxY, minZ)
			tessellator.Texture(minU, maxV)
			tessellator.Vertex(maxX, minY, minZ)
			tessellator.Texture(maxU, maxV)
			tessellator.Vertex(minX, minY, minZ)
		}
	}

	if !level.IsSolidTile(x, y, z+1) {
		brightness := level.GetBrightness(x, y, z+1) * shadeZ

		if (layer == 1) != (brightness == shadeZ) {
			tessellator.Color(brightness, brightness, brightness)
			tessellator.Texture(minU, minV)
			tessellator.Vertex(minX, maxY, maxZ)
			tessellator.Texture(minU, maxV)
			tessellator.Vertex(minX, minY, maxZ)
			tessellator.Texture(maxU, maxV)
			tessellator.Vertex(maxX, minY, maxZ)
			tessellator.Texture(maxU, minV)
			tessellator.Vertex(maxX, maxY, maxZ)
		}
	}

	if !level.IsSolidTile(x-1, y, z) {
		brightness := level.GetBrightness(x-1, y, z) * shadeX

		if (layer == 1) != (brightness == shadeX) {
			tessellator.Color(brightness, brightness, brightness)
			tessellator.Texture(maxU, minV)
			tessellator.Vertex(minX, maxY, maxZ)
			tessellator.Texture(minU, minV)
			tessellator.Vertex(minX, maxY, minZ)
			tessellator.Texture(minU, maxV)
			tessellator.Vertex(minX, minY, minZ)
			tessellator.Texture(maxU, maxV)
			tessellator.Vertex(minX, minY, maxZ)
		}
	}

	if !level.IsSolidTile(x+1, y, z) {
		brightness := level.GetBrightness(x+1, y, z) * shadeX

		if (layer == 1) != (brightness == shadeX) {
			tessellator.Color(brightness, brightness, brightness)
			tessellator.Texture(minU, maxV)
			tessellator.Vertex(maxX, minY, maxZ)
			tessellator.Texture(maxU, maxV)
			tessellator.Vertex(maxX, minY, minZ)
			tessellator.Texture(maxU, minV)
			tessellator.Vertex(maxX, maxY, minZ)
			tessellator.Texture(minU, minV)
			tessellator.Vertex(maxX, maxY, maxZ)
		}
	}
}
