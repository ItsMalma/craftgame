package game

import (
	"craftgame/game/model"
	"craftgame/pkg/gl"
	"math"
	"math/rand"
	"time"

	"github.com/chewxy/math32"
)

type Bob struct {
	Entity *Entity

	Head, Body, RightArm, LeftArm, RightLeg, LeftLeg model.Cube

	Rotation, RotationMotionFactor float32
	TimeOffset, Speed              float32
}

func NewBob(level *Level, x, y, z float32) *Bob {
	bob := new(Bob)

	bob.Entity = NewEntity(level)

	bob.RotationMotionFactor = (rand.Float32() + 1.0) * 0.01

	bob.Entity.X = x
	bob.Entity.Y = y
	bob.Entity.Z = z

	bob.TimeOffset = rand.Float32() * 1239813.0
	bob.Rotation = rand.Float32() * math32.Pi * 2.0
	bob.Speed = 1.0

	bob.Head = model.NewCube(0, 0).
		AddBox(-4.0, -8.0, -4.0, 8, 8, 8)
	bob.Body = model.NewCube(16, 16).
		AddBox(-4.0, 0.0, -2.0, 8, 12, 4)
	bob.RightArm = model.NewCube(40, 16).
		AddBox(-3.0, -2.0, -2.0, 4, 12, 4).
		SetPos(-5.0, 2.0, 0.0)
	bob.LeftArm = model.NewCube(40, 16).
		AddBox(-1.0, -2.0, -2.0, 4, 12, 4).
		SetPos(5.0, 2.0, 0.0)
	bob.RightLeg = model.NewCube(0, 16).
		AddBox(-2.0, 0.0, -2.0, 4, 12, 4).
		SetPos(-2.0, 12.0, 0.0)
	bob.LeftLeg = model.NewCube(0, 16).
		AddBox(-2.0, 0.0, -2.0, 4, 12, 4).
		SetPos(2.0, 12.0, 0.0)

	return bob
}

func (bob *Bob) Tick() {
	bob.Entity.Tick()

	x := float32(0.0)
	y := float32(0.0)

	bob.Rotation += bob.RotationMotionFactor
	bob.RotationMotionFactor *= 0.99
	bob.RotationMotionFactor += (rand.Float32() - rand.Float32()) * rand.Float32() * rand.Float32() * 0.009999999776482582

	x = math32.Sin(bob.Rotation)
	y = math32.Cos(bob.Rotation)

	if bob.Entity.OnGround && rand.Float32() < 0.01 {
		bob.Entity.MotionY = 0.12
	}

	speed := float32(0.005)
	if bob.Entity.OnGround {
		speed = 0.02
	}
	bob.Entity.MoveRelative(x, y, speed)

	bob.Entity.MotionY -= 0.005

	bob.Entity.Move(bob.Entity.MotionX, bob.Entity.MotionY, bob.Entity.MotionZ)

	bob.Entity.MotionX *= 0.91
	bob.Entity.MotionY *= 0.98
	bob.Entity.MotionZ *= 0.91

	if bob.Entity.Y > 100.0 {
		bob.Entity.resetPosition()
	}

	if bob.Entity.OnGround {
		bob.Entity.MotionX *= 0.8
		bob.Entity.MotionZ *= 0.8
	}
}

func (bob *Bob) Render(partialTicks float32) error {
	gl.Enable(gl.Texture2D)

	bobTexture, err := LoadTexture("bob.png", gl.Nearest)
	if err != nil {
		return err
	}
	gl.BindTexture(gl.Texture2D, bobTexture)

	gl.PushMatrix()

	time := float64(time.Now().UnixNano())/1.0e9*10.0*float64(bob.Speed) + float64(bob.TimeOffset)
	size := float32(0.058333334)
	offsetY := float32(-math.Abs(math.Sin(time*0.6662))*5.0 - 23.0)

	gl.Translatef(
		bob.Entity.PrevX+(bob.Entity.X-bob.Entity.PrevX)*partialTicks,
		bob.Entity.PrevY+(bob.Entity.Y-bob.Entity.PrevY)*partialTicks,
		bob.Entity.PrevZ+(bob.Entity.Z-bob.Entity.PrevZ)*partialTicks,
	)
	gl.Scalef(1.0, -1.0, 1.0)
	gl.Scalef(size, size, size)
	gl.Translatef(0.0, offsetY, 0.0)
	gl.Rotatef(
		bob.Rotation*57.29578+180.0,
		0.0, 1.0, 0.0,
	)

	bob.Head.YRotation = float32(math.Sin(time*0.83)) * 1.0
	bob.Head.XRotation = float32(math.Sin(time)) * 0.8
	bob.RightArm.XRotation = float32(math.Sin(time*0.6662+math.Pi)) * 2.0
	bob.RightArm.ZRotation = float32((math.Sin(time*0.2312) + 1.0)) * 1.0
	bob.LeftArm.XRotation = float32(math.Sin(time*0.6662)) * 2.0
	bob.LeftArm.ZRotation = float32((math.Sin(time*0.2812) - 1.0)) * 1.0
	bob.RightLeg.XRotation = float32(math.Sin(time*0.6662)) * 1.4
	bob.LeftLeg.XRotation = float32(math.Sin(time*0.6662+math.Pi)) * 1.4
	bob.Head.Render()
	bob.Body.Render()
	bob.RightArm.Render()
	bob.LeftArm.Render()
	bob.RightLeg.Render()
	bob.LeftLeg.Render()

	gl.PopMatrix()

	gl.Disable(gl.Texture2D)

	return nil
}
