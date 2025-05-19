package main

import (
	"fmt"
	"runtime"

	"github.com/ItsMalma/craftgame/game"
)

func init() {
	runtime.LockOSThread()
}

func main() {
	fmt.Println("Running!")

	g := game.NewGame()
	defer g.Destroy()
	g.Run()
}
