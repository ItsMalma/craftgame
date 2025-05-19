package entity

import (
	"math"
	"math/rand"

	"github.com/ItsMalma/craftgame/game/world"
	"github.com/ItsMalma/craftgame/phys"
)

type Entity struct {
	world *world.World

	XO, YO, ZO, X, Y, Z, XD, YD, ZD, XRot, YRot float32
	BB                                          phys.AABB

	OnGround bool

	heightOffset float32
}

func New(w *world.World) *Entity {
	e := new(Entity)
	e.world = w
	e.OnGround = false
	e.resetPosition()

	return e
}

func (e *Entity) resetPosition() {
	x := rand.Float32() * float32(e.world.Width())
	y := float32(e.world.Depth() + 10)
	z := rand.Float32() * float32(e.world.Height())

	e.SetPosition(x, y, z)
}

func (e *Entity) SetPosition(x, y, z float32) {
	const width = 0.3
	const height = 0.9

	e.X = x
	e.Y = y
	e.Z = z
	e.BB = phys.NewAABB(x-width, y-height, z-width, x+width, y+height, z+width)
}

func (e *Entity) Turn(xo, yo float32) {
	e.YRot = float32(float64(e.YRot) + float64(xo)*0.15)
	e.XRot = float32(float64(e.XRot) + float64(yo)*0.15)

	if e.XRot < -90.0 {
		e.XRot = -90.0
	}
	if e.XRot > 90.0 {
		e.XRot = 90.0
	}
}

func (e *Entity) Tick() {
	e.XO = e.X
	e.YO = e.Y
	e.ZO = e.Z
}

func (e *Entity) Move(xa, ya, za float32) {
	xaOrg := xa
	yaOrg := ya
	zaOrg := za
	aabbs := e.world.GetCubes(e.BB.Expand(xa, ya, za))

	for _, aabb := range aabbs {
		ya = aabb.ClipYCollide(e.BB, ya)
	}
	e.BB = e.BB.Move(0, ya, 0)

	for _, aabb := range aabbs {
		xa = aabb.ClipXCollide(e.BB, xa)
	}
	e.BB = e.BB.Move(xa, 0, 0)

	for _, aabb := range aabbs {
		za = aabb.ClipZCollide(e.BB, za)
	}
	e.BB = e.BB.Move(0, 0, za)

	e.OnGround = yaOrg != ya && yaOrg < 0.0

	if xaOrg != xa {
		e.XD = 0.0
	}
	if yaOrg != ya {
		e.YD = 0.0
	}
	if zaOrg != za {
		e.ZD = 0.0
	}

	e.X = (e.BB.X0 + e.BB.X1) / 2.0
	e.Y = e.BB.Y0 + e.heightOffset
	e.Z = (e.BB.Z0 + e.BB.Z1) / 2.0
}

func (e *Entity) MoveRelative(xa, za, speed float32) {
	dist := xa*xa + za*za
	if dist >= 0.01 {
		dist = speed / float32(math.Sqrt(float64(dist)))
		xa *= dist
		za *= dist
		sin := float32(math.Sin(float64(e.YRot) * math.Pi / 180.0))
		cos := float32(math.Cos(float64(e.YRot) * math.Pi / 180.0))
		e.XD += xa*cos - za*sin
		e.ZD += za*cos + xa*sin
	}
}
