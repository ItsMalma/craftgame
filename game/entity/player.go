package entity

import (
	"math"
	"math/rand"

	"github.com/ItsMalma/craftgame/game/windowing"
	"github.com/ItsMalma/craftgame/game/world"
	"github.com/ItsMalma/craftgame/phys"
)

type Player struct {
	w *world.World

	XO, YO, ZO, X, Y, Z, XD, YD, ZD, XRot, YRot float32
	BB                                          phys.AABB

	OnGround bool
}

func NewPlayer(w *world.World) *Player {
	player := new(Player)
	player.w = w
	player.OnGround = false

	player.resetPosition()

	return player
}

func (player *Player) resetPosition() {
	x := float32(rand.Float64()) * float32(player.w.Width())
	y := float32(player.w.Depth() + 10)
	z := float32(rand.Float64()) * float32(player.w.Height())

	player.SetPosition(x, y, z)
}

func (player *Player) SetPosition(x, y, z float32) {
	const width = 0.3
	const height = 0.9

	player.X = x
	player.Y = y
	player.Z = z
	player.BB = phys.NewAABB(x-width, y-height, z-width, x+width, y+height, z+width)
}

func (player *Player) Turn(xo, yo float32) {
	player.YRot = float32(float64(player.YRot) + float64(xo)*0.15)
	player.XRot = float32(float64(player.XRot) + float64(yo)*0.15)

	if player.XRot < -90.0 {
		player.XRot = -90.0
	}
	if player.XRot > 90.0 {
		player.XRot = 90.0
	}
}

func (player *Player) Tick(input windowing.Input) {
	player.XO = player.X
	player.YO = player.Y
	player.ZO = player.Z

	var xa, ya float32

	if input.IsKeyDown(windowing.KeyR) {
		player.resetPosition()
	}

	if input.IsKeyDown(windowing.KeyW) {
		ya--
	}
	if input.IsKeyDown(windowing.KeyS) {
		ya++
	}
	if input.IsKeyDown(windowing.KeyA) {
		xa--
	}
	if input.IsKeyDown(windowing.KeyD) {
		xa++
	}
	if input.IsKeyDown(windowing.KeySpace) && player.OnGround {
		player.YD = 0.12
	}

	speed := float32(0.005)
	if player.OnGround {
		speed = 0.02
	}
	player.MoveRelative(xa, ya, speed)

	player.YD = float32(float64(player.YD) - 0.005)

	player.Move(player.XD, player.YD, player.ZD)

	player.XD *= 0.91
	player.YD *= 0.98
	player.ZD *= 0.91

	if player.OnGround {
		player.XD *= 0.8
		player.ZD *= 0.8
	}
}

func (player *Player) Move(xa, ya, za float32) {
	xaOrg := xa
	yaOrg := ya
	zaOrg := za

	aabbs := player.w.GetCubes(player.BB.Expand(xa, ya, za))

	for _, aabb := range aabbs {
		ya = aabb.ClipYCollide(player.BB, ya)
	}
	player.BB = player.BB.Move(0, ya, 0)

	for _, aabb := range aabbs {
		xa = aabb.ClipXCollide(player.BB, xa)
	}
	player.BB = player.BB.Move(xa, 0, 0)

	for _, aabb := range aabbs {
		za = aabb.ClipZCollide(player.BB, za)
	}
	player.BB = player.BB.Move(0, 0, za)

	player.OnGround = yaOrg != ya && yaOrg < 0.0

	if xaOrg != xa {
		player.XD = 0.0
	}
	if yaOrg != ya {
		player.YD = 0.0
	}
	if zaOrg != za {
		player.ZD = 0.0
	}

	player.X = (player.BB.X0 + player.BB.X1) / 2.0
	player.Y = player.BB.Y0 + 1.62
	player.Z = (player.BB.Z0 + player.BB.Z1) / 2.0
}

func (player *Player) MoveRelative(xa, za, speed float32) {
	dist := xa*xa + za*za
	if dist >= 0.01 {
		dist = speed / float32(math.Sqrt(float64(dist)))
		xa *= dist
		za *= dist
		sin := float32(math.Sin(float64(player.YRot) * math.Pi / 180.0))
		cos := float32(math.Cos(float64(player.YRot) * math.Pi / 180.0))
		player.XD += xa*cos - za*sin
		player.ZD += za*cos + xa*sin
	}
}
