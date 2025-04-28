package game

import (
	"compress/gzip"
	"io"
	"math"
	"math/rand"
	"os"

	"github.com/chewxy/math32"
)

type Level struct {
	Width, Height, Depth int

	blocks      []byte
	lightDepths []byte
}

func NewLevel(width, height, depth int) (*Level, error) {
	level := new(Level)

	level.Width = width
	level.Height = height
	level.Depth = depth

	level.blocks = make([]byte, width*height*depth)
	level.lightDepths = make([]byte, width*height*depth)

	_, err := os.Stat("level.dat")
	if err == nil || os.IsExist(err) {
		if err := level.Load(); err != nil {
			return nil, err
		}
	} else if !os.IsNotExist(err) {
		return nil, err
	}

	for x := range width {
		for y := range depth {
			for z := range height {
				index := (y*height+z)*width + x

				level.blocks[index] = 1
			}
		}
	}

	for range 10000 {
		caveSize := int(rand.Float32()*7) + 1

		caveX := int(rand.Float32() * float32(width))
		caveY := int(rand.Float32() * float32(height))
		caveZ := int(rand.Float32() * float32(depth))

		for radius := range caveSize {
			for range 1000 {
				offsetX := int(rand.Float32()*float32(radius)*2.0 - float32(radius))
				offsetY := int(rand.Float32()*float32(radius)*2.0 - float32(radius))
				offsetZ := int(rand.Float32()*float32(radius)*2.0 - float32(radius))

				distance := math32.Pow(float32(offsetX), 2) + math32.Pow(float32(offsetY), 2) + math32.Pow(float32(offsetZ), 2)
				if distance > float32(radius)*float32(radius) {
					continue
				}

				tileX := caveX + offsetX
				tileY := caveY + offsetY
				tileZ := caveZ + offsetZ

				index := (tileY*height+tileZ)*width + tileX

				if index >= 0 && index < len(level.blocks) {
					if tileX > 0 && tileY > 0 && tileZ > 0 && tileX < width-1 && tileY < depth && tileZ < height-1 {
						level.blocks[index] = 0
					}
				}
			}
		}
	}

	level.calculateLightDepths(0, 0, width, height)

	return level, nil
}

func (level *Level) Load() error {
	levelFile, err := os.Open("level.dat")
	if err != nil {
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

	return nil
}

func (level *Level) calculateLightDepths(minX, minZ, maxX, maxZ int) {
	for x := minX; x < minX+maxX; x++ {
		for z := minZ; z < minZ+maxZ; z++ {
			// prevDepth := level.lightDepths[x+z*level.Width]

			depth := level.Depth - 1
			for depth > 0 && !level.IsLightBlocker(x, depth, z) {
				depth--
			}

			level.lightDepths[x+z*level.Width] = byte(depth)
		}
	}
}

func (level *Level) IsTile(x, y, z int) bool {
	if x < 0 || y < 0 || z < 0 || x >= level.Width || y >= level.Depth || z >= level.Height {
		return false
	}

	index := (y*level.Height+z)*level.Width + x

	return level.blocks[index] != 0
}

func (level *Level) IsSolidTile(x, y, z int) bool {
	return level.IsTile(x, y, z)
}

func (level *Level) IsLightBlocker(x, y, z int) bool {
	return level.IsSolidTile(x, y, z)
}

func (level *Level) GetBrightness(x, y, z int) float32 {
	dark := 0.8
	light := 1.0

	if x < 0 || y < 0 || z < 0 || x >= level.Width || y >= level.Depth || z >= level.Height {
		return float32(light)
	}

	if y < int(level.lightDepths[x+z*level.Width]) {
		return float32(dark)
	}

	return float32(light)
}

func (level *Level) GetCubes(boundingBox *AABB) []*AABB {
	boundingBoxes := []*AABB{}

	minX := int(math.Floor(boundingBox.MinX) - 1)
	maxX := int(math.Ceil(boundingBox.MaxX) + 1)
	minY := int(math.Floor(boundingBox.MinY) - 1)
	maxY := int(math.Ceil(boundingBox.MaxY) + 1)
	minZ := int(math.Floor(boundingBox.MinZ) - 1)
	maxZ := int(math.Ceil(boundingBox.MaxZ) + 1)

	minX = max(0, minX)
	minY = max(0, minY)
	minZ = max(0, minZ)

	maxX = min(level.Width, maxX)
	maxY = min(level.Depth, maxY)
	maxZ = min(level.Height, maxZ)

	for x := minX; x < maxX; x++ {
		for y := minY; y < maxY; y++ {
			for z := minZ; z < maxZ; z++ {
				if level.IsSolidTile(x, y, z) {
					boundingBoxes = append(
						boundingBoxes,
						NewAABB(
							float64(x), float64(y), float64(z),
							float64(x)+1, float64(y)+1, float64(z)+1,
						),
					)
				}
			}
		}
	}

	return boundingBoxes
}
