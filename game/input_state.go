package game

import (
	"github.com/veandco/go-sdl2/sdl"
)

type InputState struct {
}

func (inputState *InputState) IsKeyDown(key sdl.Keycode) bool {
	return true
}
