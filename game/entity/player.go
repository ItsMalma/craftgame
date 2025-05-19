package entity

import (
	"github.com/ItsMalma/craftgame/game/windowing"
	"github.com/ItsMalma/craftgame/game/world"
)

type Player struct {
	*Entity
}

func NewPlayer(w *world.World) *Player {
	player := new(Player)
	player.Entity = New(w)
	player.heightOffset = 1.62

	return player
}

func (player *Player) Tick(input windowing.Input) {
	player.Entity.Tick()

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
