package game

import (
	"craftgame/pkg/glfw"
)

type Player struct {
	Entity *Entity
}

func NewPlayer(level *Level) *Player {
	player := new(Player)

	player.Entity = NewEntity(level)

	player.Entity.heightOffset = 1.62

	return player
}

func (player *Player) Tick(window *glfw.Window) {
	player.Entity.Tick()

	forward := float32(0.0)
	vertical := float32(0.0)

	if glfw.GetKey(window, glfw.KeyR) == glfw.Press {
		player.Entity.resetPosition()
	}

	if glfw.GetKey(window, glfw.KeyW) == glfw.Press {
		forward--
	}
	if glfw.GetKey(window, glfw.KeyS) == glfw.Press {
		forward++
	}
	if glfw.GetKey(window, glfw.KeyA) == glfw.Press {
		vertical--
	}
	if glfw.GetKey(window, glfw.KeyD) == glfw.Press {
		vertical++
	}
	if glfw.GetKey(window, glfw.KeySpace) == glfw.Press && player.Entity.OnGround {
		player.Entity.MotionY = 0.12
	}

	speed := float32(0.005)
	if player.Entity.OnGround {
		speed = 0.02
	}
	player.Entity.MoveRelative(vertical, forward, speed)

	player.Entity.MotionY -= 0.005

	player.Entity.Move(player.Entity.MotionX, player.Entity.MotionY, player.Entity.MotionZ)

	player.Entity.MotionX *= 0.91
	player.Entity.MotionY *= 0.98
	player.Entity.MotionZ *= 0.91

	if player.Entity.OnGround {
		player.Entity.MotionX *= 0.8
		player.Entity.MotionZ *= 0.8
	}
}
