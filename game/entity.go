package game

import (
	"math/rand"

	"github.com/chewxy/math32"
)

type Entity struct {
	level *Level

	X, Y, Z                   float32
	PrevX, PrevY, PrevZ       float32
	MotionX, MotionY, MotionZ float32
	XRotation, YRotation      float32

	BoundingBox AABB

	OnGround bool

	heightOffset float32
}

func NewEntity(level *Level) *Entity {
	entity := new(Entity)

	entity.OnGround = false
	entity.heightOffset = 0.0

	entity.level = level

	entity.resetPosition()

	return entity
}

func (entity *Entity) resetPosition() {
	x := rand.Float32() * float32(entity.level.Width)
	y := float32(entity.level.Depth) + 3.0
	z := rand.Float32() * float32(entity.level.Height)

	entity.setPosition(x, y, z)
}

func (entity *Entity) setPosition(x, y, z float32) {
	entity.X = x
	entity.Y = y
	entity.Z = z

	var (
		width  float32 = 0.3
		height float32 = 0.9
	)

	entity.BoundingBox = NewAABB(
		x-width,
		y-height,
		z-width,
		x+width,
		y+height,
		z+width,
	)
}

func (entity *Entity) Turn(x, y float32) {
	entity.YRotation += x * 0.15
	entity.XRotation += y * 0.15

	if entity.XRotation < -90.0 {
		entity.XRotation = -90.0
	}
	if entity.XRotation > 90.0 {
		entity.XRotation = 90.0
	}
}

func (entity *Entity) Tick() {
	entity.PrevX = entity.X
	entity.PrevY = entity.Y
	entity.PrevZ = entity.Z
}

func (entity *Entity) Move(x, y, z float32) {
	prevX := x
	prevY := y
	prevZ := z

	boundingBoxes := entity.level.GetCubes(entity.BoundingBox.Expand(x, y, z))

	for _, boundingBox := range boundingBoxes {
		y = boundingBox.ClipYCollide(entity.BoundingBox, y)
	}
	entity.BoundingBox = entity.BoundingBox.Move(0.0, y, 0.0)

	for _, boundingBox := range boundingBoxes {
		x = boundingBox.ClipXCollide(entity.BoundingBox, x)
	}
	entity.BoundingBox = entity.BoundingBox.Move(x, 0.0, 0.0)

	for _, boundingBox := range boundingBoxes {
		z = boundingBox.ClipZCollide(entity.BoundingBox, z)
	}
	entity.BoundingBox = entity.BoundingBox.Move(0.0, 0.0, z)

	entity.OnGround = prevY != y && prevY < 0.0

	if prevX != x {
		entity.MotionX = 0.0
	}
	if prevY != y {
		entity.MotionY = 0.0
	}
	if prevZ != z {
		entity.MotionZ = 0.0
	}

	entity.X = (entity.BoundingBox.MinX + entity.BoundingBox.MaxX) / 2.0
	entity.Y = entity.BoundingBox.MinY + entity.heightOffset
	entity.Z = (entity.BoundingBox.MinZ + entity.BoundingBox.MaxZ) / 2.0
}

func (entity *Entity) MoveRelative(x, z, speed float32) {
	distance := x*x + z*z

	if distance >= 0.01 {
		distance = speed / math32.Sqrt(distance)

		x *= distance
		z *= distance

		sin := math32.Sin(entity.YRotation * math32.Pi / 180.0)
		cos := math32.Cos(entity.YRotation * math32.Pi / 180.0)

		entity.MotionX += x*cos - z*sin
		entity.MotionZ += z*cos + x*sin
	}
}
