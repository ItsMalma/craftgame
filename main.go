package main

import (
	"runtime"

	"github.com/ItsMalma/craftgame/game"
)

func init() {
	runtime.LockOSThread()
}

func main() {
	g := game.NewGame()
	defer g.Destroy()
	g.Run()
}
