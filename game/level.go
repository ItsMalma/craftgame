package game

import (
	"compress/gzip"
	"io"
	"os"
)

type LevelListener interface {
	LightColumnChanged(x, z, minY, maxY int)
	TileChanged(x, y, z int)
	AllChanged()
}

type Level struct {
	Width, Height, Depth int

	blocks      []byte
	lightDepths []int

	listeners []LevelListener
}

func NewLevel(width, height, depth int) (*Level, error) {
	level := new(Level)

	level.Width = width
	level.Height = height
	level.Depth = depth

	level.blocks = make([]byte, width*height*depth)
	level.lightDepths = make([]int, width*height)

	level.listeners = []LevelListener{}

	for x := range width {
		for y := range depth {
			for z := range height {
				index := (y*height+z)*width + x

				var block byte = 0
				if y <= depth*2/3 {
					block = 1
				}

				level.blocks[index] = block
			}
		}
	}

	level.calculateLightDepths(0, 0, width, height)

	if err := level.Load(); err != nil {
		return nil, err
	}

	// for range 10000 {
	// 	caveSize := rand.Intn(7) + 1

	// 	caveX := rand.Intn(width)
	// 	caveY := rand.Intn(height)
	// 	caveZ := rand.Intn(depth)

	// 	for radius := range caveSize {
	// 		for range 1000 {
	// 			offsetX := int(rand.Float32()*float32(radius)*2.0 - float32(radius))
	// 			offsetY := int(rand.Float32()*float32(radius)*2.0 - float32(radius))
	// 			offsetZ := int(rand.Float32()*float32(radius)*2.0 - float32(radius))

	// 			distance := math32.Pow(float32(offsetX), 2) + math32.Pow(float32(offsetY), 2) + math32.Pow(float32(offsetZ), 2)
	// 			if distance > float32(radius)*float32(radius) {
	// 				continue
	// 			}

	// 			tileX := caveX + offsetX
	// 			tileY := caveY + offsetY
	// 			tileZ := caveZ + offsetZ

	// 			index := (tileY*height+tileZ)*width + tileX

	// 			if index >= 0 && index < len(level.blocks) {
	// 				if tileX > 0 && tileY > 0 && tileZ > 0 && tileX < width-1 && tileY < depth && tileZ < height-1 {
	// 					level.blocks[index] = 0
	// 				}
	// 			}
	// 		}
	// 	}
	// }

	return level, nil
}

func (level *Level) Load() error {
	levelFile, err := os.Open("level.dat")
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	defer levelFile.Close()

	gzipReader, err := gzip.NewReader(levelFile)
	if err != nil {
		return err
	}
	defer gzipReader.Close()

	_, err = io.ReadFull(gzipReader, level.blocks)
	if err != nil {
		return err
	}

	level.calculateLightDepths(0, 0, level.Width, level.Height)

	for _, listener := range level.listeners {
		listener.AllChanged()
	}

	return nil
}

func (level *Level) Save() error {
	levelFile, err := os.OpenFile("level.dat", os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		return err
	}
	defer levelFile.Close()

	gzipWriter := gzip.NewWriter(levelFile)
	defer gzipWriter.Close()

	_, err = gzipWriter.Write(level.blocks)
	if err != nil {
		return err
	}

	return nil
}

func (level *Level) calculateLightDepths(x0, y0, x1, y1 int) {
	for x := x0; x < x0+x1; x++ {
		for z := y0; z < y0+y1; z++ {
			oldDepth := level.lightDepths[x+z*level.Width]

			y := level.Depth - 1
			for y > 0 && !level.IsLightBlocker(x, y, z) {
				y--
			}

			level.lightDepths[x+z*level.Width] = y

			if oldDepth != y {
				yl0 := min(oldDepth, y)
				yl1 := max(oldDepth, y)

				for _, listener := range level.listeners {
					listener.LightColumnChanged(x, z, yl0, yl1)
				}
			}
		}
	}
}

func (level *Level) IsTile(x, y, z int) bool {
	if x >= 0 && y >= 0 && z >= 0 && x < level.Width && y < level.Depth && z < level.Height {
		return level.blocks[(y*level.Height+z)*level.Width+x] == 1
	}

	return false
}

func (level *Level) IsSolidTile(x, y, z int) bool {
	return level.IsTile(x, y, z)
}

func (level *Level) IsLightBlocker(x, y, z int) bool {
	return level.IsSolidTile(x, y, z)
}

func (level *Level) GetBrightness(x, y, z int) float32 {
	var dark float32 = 0.8
	var light float32 = 1.0

	if x >= 0 && y >= 0 && z >= 0 && x < level.Width && y < level.Depth && z < level.Height {
		if y < level.lightDepths[x+z*level.Width] {
			return dark
		} else {
			return light
		}
	} else {
		return light
	}
}

func (level *Level) GetCubes(boundingBox AABB) []AABB {
	boundingBoxes := []AABB{}

	x0 := int(boundingBox.MinX)
	x1 := int(boundingBox.MaxX + 1.0)
	y0 := int(boundingBox.MinY)
	y1 := int(boundingBox.MaxY + 1.0)
	z0 := int(boundingBox.MinZ)
	z1 := int(boundingBox.MaxZ + 1.0)

	if x0 < 0 {
		x0 = 0
	}
	if y0 < 0 {
		y0 = 0
	}
	if z0 < 0 {
		z0 = 0
	}

	if x1 > level.Width {
		x1 = level.Width
	}
	if y1 > level.Depth {
		y1 = level.Depth
	}
	if z1 > level.Height {
		z1 = level.Height
	}

	for x := x0; x < x1; x++ {
		for y := y0; y < y1; y++ {
			for z := z0; z < z1; z++ {
				if level.IsSolidTile(x, y, z) {
					boundingBoxes = append(
						boundingBoxes,
						NewAABB(
							float32(x), float32(y), float32(z),
							float32(x)+1, float32(y)+1, float32(z)+1,
						),
					)
				}
			}
		}
	}

	return boundingBoxes
}

func (level *Level) SetTile(x, y, z, id int) {
	if x >= 0 && y >= 0 && z >= 0 && x < level.Width && y < level.Depth && z < level.Height {
		level.blocks[(y*level.Height+z)*level.Width+x] = byte(id)
		level.calculateLightDepths(x, z, 1, 1)

		for _, listener := range level.listeners {
			listener.TileChanged(x, y, z)
		}
	}
}

func (level *Level) AddListener(listener LevelListener) {
	level.listeners = append(level.listeners, listener)
}
