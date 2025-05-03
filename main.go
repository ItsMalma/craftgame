package main

import (
	"craftgame/game"
	"log"
	"runtime"
)

func init() {
	runtime.LockOSThread()
}

func main() {
	g, err := game.NewGame()
	if err != nil {
		log.Fatal(err)
	}
	defer g.Close()

	g.Run()
}
