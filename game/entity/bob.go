package entity

import (
	"math"
	"math/rand"
	"time"

	"github.com/ItsMalma/craftgame/game/entity/model"
	"github.com/ItsMalma/craftgame/game/world"
	"github.com/go-gl/gl/v2.1/gl"
)

type Bob struct {
	*Entity

	texture uint32

	Head, Body, RightArm, LeftArm, RightLeg, LeftLeg model.Cube

	Rot, TimeOffs, Speed, RotA float32
}

func NewBob(w *world.World, x, y, z float32, texture uint32) *Bob {
	bob := new(Bob)
	bob.Entity = New(w)
	bob.X = x
	bob.Y = y
	bob.Z = z
	bob.texture = texture
	bob.TimeOffs = rand.Float32() * 1239813.0
	bob.Rot = rand.Float32() * math.Pi * 2.0
	bob.Speed = 1.0
	bob.RotA = (rand.Float32() + 1.0) * 0.01

	bob.Head = model.NewCube(0, 0).AddBox(-4.0, -8.0, -4.0, 8, 8, 8)
	bob.Body = model.NewCube(16, 16).AddBox(-4.0, 0.0, -2.0, 8, 12, 4)
	bob.RightArm = model.NewCube(40, 16).AddBox(-3.0, -2.0, -2.0, 4, 12, 4).SetPos(-5.0, 2.0, 0.0)
	bob.LeftArm = model.NewCube(40, 16).AddBox(-1.0, -2.0, -2.0, 4, 12, 4).SetPos(5.0, 2.0, 0.0)
	bob.RightLeg = model.NewCube(0, 16).AddBox(-2.0, 0.0, -2.0, 4, 12, 4).SetPos(-2.0, 12.0, 0.0)
	bob.LeftLeg = model.NewCube(0, 16).AddBox(-2.0, 0.0, -2.0, 4, 12, 4).SetPos(2.0, 12.0, 0.0)

	return bob
}

func (bob *Bob) Tick() {
	bob.Entity.Tick()

	var xa, ya float32

	bob.Rot += bob.RotA
	bob.RotA = bob.RotA * 0.99
	bob.RotA = bob.RotA + (rand.Float32()-rand.Float32())*rand.Float32()*rand.Float32()*0.009999999776482582

	xa = float32(math.Sin(float64(bob.Rot)))
	ya = float32(math.Cos(float64(bob.Rot)))

	if bob.OnGround && rand.Float32() < 0.01 {
		bob.YD = 0.12
	}

	speed := float32(0.005)
	if bob.OnGround {
		speed = 0.02
	}
	bob.MoveRelative(xa, ya, speed)

	bob.YD = float32(float64(bob.YD) - 0.005)

	bob.Move(bob.XD, bob.YD, bob.ZD)

	bob.XD *= 0.91
	bob.YD *= 0.98
	bob.ZD *= 0.91

	if bob.Y > 100.0 {
		bob.resetPosition()
	}

	if bob.OnGround {
		bob.XD *= 0.8
		bob.ZD *= 0.8
	}
}

func (bob *Bob) Render(a float32) {
	gl.Enable(gl.TEXTURE_2D)
	gl.BindTexture(gl.TEXTURE_2D, bob.texture)
	gl.PushMatrix()

	time := float64(time.Now().UnixNano())/1.0e9*10.0*float64(bob.Speed) + float64(bob.TimeOffs)
	size := float32(0.058333334)
	yy := float32(-math.Abs(math.Sin(time*0.6662))*5.0 - 23.0)

	gl.Translatef(bob.XO+(bob.X-bob.XO)*a, bob.YO+(bob.Y-bob.YO)*a, bob.ZO+(bob.Z-bob.ZO)*a)
	gl.Scalef(1.0, -1.0, 1.0)
	gl.Scalef(size, size, size)
	gl.Translatef(0.0, yy, 0.0)
	c := float32(57.29578)
	gl.Rotatef(bob.Rot*c+180.0, 0.0, 1.0, 0.0)

	bob.Head.YRot = float32(math.Sin(time*0.83)) * 1.0
	bob.Head.XRot = float32(math.Sin(time)) * 0.8
	bob.RightArm.XRot = float32(math.Sin(time*0.6662+math.Pi)) * 2.0
	bob.RightArm.ZRot = float32(math.Sin(time*0.2312)+1.0) * 1.0
	bob.LeftArm.XRot = float32(math.Sin(time*0.6662)) * 2.0
	bob.LeftArm.ZRot = float32(math.Sin(time*0.2812)-1.0) * 1.0
	bob.RightLeg.XRot = float32(math.Sin(time*0.6662)) * 1.4
	bob.LeftLeg.XRot = float32(math.Sin(time*0.6662+math.Pi)) * 1.4

	bob.Head.Render()
	bob.Body.Render()
	bob.RightArm.Render()
	bob.LeftArm.Render()
	bob.RightLeg.Render()
	bob.LeftLeg.Render()

	gl.PopMatrix()
	gl.Disable(gl.TEXTURE_2D)
}
