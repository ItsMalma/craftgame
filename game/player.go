package game

import (
	"math"
	"math/rand"

	"github.com/go-gl/glfw/v3.3/glfw"
)

type Player struct {
	level *Level

	X, Y, Z                   float64
	PrevX, PrevY, PrevZ       float64
	MotionX, MotionY, MotionZ float64
	XRotation, YRotation      float64

	onGround bool

	BoundingBox *AABB
}

func NewPlayer(level *Level) *Player {
	player := new(Player)

	player.onGround = false

	player.level = level

	player.resetPosition()

	return player
}

func (player *Player) setPosition(x, y, z float32) {
	player.X = float64(x)
	player.Y = float64(y)
	player.Z = float64(z)

	width := 0.3
	height := 0.9

	player.BoundingBox = NewAABB(
		float64(x)-width,
		float64(y)-height,
		float64(z)-width,
		float64(x)+width,
		float64(y)+height,
		float64(z)+width,
	)
}

func (player *Player) resetPosition() {
	x := rand.Float32() * float32(player.level.Width)
	y := float32(player.level.Depth) + 3.0
	z := rand.Float32() * float32(player.level.Height)

	player.setPosition(x, y, z)
}

func (player *Player) Turn(x, y float64) {
	player.YRotation += x * 0.15
	player.XRotation += y * 0.15

	player.XRotation = max(-90.0, player.XRotation)
	player.XRotation = min(90.0, player.XRotation)
}

func (player *Player) Tick(window *glfw.Window) {
	player.PrevX = player.X
	player.PrevY = player.Y
	player.PrevZ = player.Z

	forward := 0.0
	vertical := 0.0

	if window.GetKey(glfw.KeyR) == glfw.Press {
		player.resetPosition()
	}

	if window.GetKey(glfw.KeyW) == glfw.Press {
		forward--
	}
	if window.GetKey(glfw.KeyS) == glfw.Press {
		forward++
	}
	if window.GetKey(glfw.KeyA) == glfw.Press {
		vertical--
	}
	if window.GetKey(glfw.KeyD) == glfw.Press {
		vertical++
	}
	if window.GetKey(glfw.KeySpace) == glfw.Press {
		if player.onGround {
			player.MotionY = 0.12
		}
	}

	speed := 0.005
	if player.onGround {
		speed = 0.02
	}
	player.moveRelative(vertical, forward, speed)

	player.MotionY -= 0.005

	player.Move(player.MotionX, player.MotionY, player.MotionZ)

	player.MotionX *= 0.91
	player.MotionY *= 0.98
	player.MotionZ *= 0.91

	if player.onGround {
		player.MotionX *= 0.8
		player.MotionZ *= 0.8
	}
}

func (player *Player) Move(x, y, z float64) {
	prevX := x
	prevY := y
	prevZ := z

	aabbs := player.level.GetCubes(player.BoundingBox.Expand(x, y, z))

	for _, aabb := range aabbs {
		y = aabb.ClipYCollide(player.BoundingBox, y)
	}
	player.BoundingBox.Move(0, y, 0)

	for _, aabb := range aabbs {
		x = aabb.ClipXCollide(player.BoundingBox, x)
	}
	player.BoundingBox.Move(x, 0, 0)

	for _, aabb := range aabbs {
		z = aabb.ClipZCollide(player.BoundingBox, z)
	}
	player.BoundingBox.Move(0, 0, z)

	player.onGround = prevY != y && prevY < 0.0

	if prevX != x {
		player.MotionX = 0.0
	}
	if prevY != y {
		player.MotionY = 0.0
	}
	if prevZ != z {
		player.MotionZ = 0.0
	}

	player.X = (player.BoundingBox.MinX + player.BoundingBox.MaxX) / 2.0
	player.Y = player.BoundingBox.MinY + 1.62
	player.Z = (player.BoundingBox.MinZ + player.BoundingBox.MaxZ) / 2.0
}

func (player *Player) moveRelative(x, z, speed float64) {
	distance := x*x + z*z

	if distance < 0.01 {
		return
	}

	distance = speed / math.Sqrt(distance)
	x *= distance
	z *= distance

	sin := math.Sin(player.YRotation * math.Pi / 180.0)
	cos := math.Cos(player.YRotation * math.Pi / 180.0)

	player.MotionX += x*cos - z*sin
	player.MotionZ += z*cos + x*sin
}
