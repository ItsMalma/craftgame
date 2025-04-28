package main

import (
	"log"
	"minecraft/game"
	"minecraft/pkg/glu"
	"runtime"
	"time"

	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

var (
	lastMouseX float64
	lastMouseY float64
)

func init() {
	runtime.LockOSThread()
}

func main() {
	timer := game.NewTimer(60)

	fogColor := [4]float32{
		14.0 / 255.0,
		11.0 / 255.0,
		10.0 / 255.0,
		1,
	}

	width := 1024
	height := 768

	if err := glfw.Init(); err != nil {
		log.Fatal(err)
	}

	window, err := glfw.CreateWindow(
		width, height,
		"Minecraft",
		nil, nil,
	)
	if err != nil {
		log.Fatal(err)
	}
	defer window.Destroy()

	window.MakeContextCurrent()
	glfw.SwapInterval(0)

	if err := gl.Init(); err != nil {
		log.Fatal(err)
	}

	gl.Enable(gl.TEXTURE_2D)
	gl.ShadeModel(gl.SMOOTH)
	gl.ClearColor(0.5, 0.8, 1.0, 0.0)
	gl.ClearDepth(1.0)
	gl.Enable(gl.DEPTH_TEST)
	gl.Enable(gl.CULL_FACE)
	gl.DepthFunc(gl.LEQUAL)

	gl.MatrixMode(gl.PROJECTION)
	gl.LoadIdentity()
	glu.Perspective(70, float64(width)/float64(height), 0.05, 1000.0)
	gl.MatrixMode(gl.MODELVIEW)

	level, err := game.NewLevel(256, 256, 64)
	if err != nil {
		log.Fatal(err)
	}
	renderer := game.NewRenderer(level)
	player := game.NewPlayer(level)

	window.SetInputMode(glfw.CursorMode, glfw.CursorDisabled)

	frames := 0
	lastTime := time.Now().UnixMilli()

	lastMouseX, lastMouseY = window.GetCursorPos()

	for !window.ShouldClose() {
		if window.GetKey(glfw.KeyEscape) == glfw.Press {
			window.SetShouldClose(true)
		}

		timer.AdvanceTime()

		for range timer.Ticks {
			player.Tick(window)
		}

		currentMouseX, currentMouseY := window.GetCursorPos()

		partialTicks := timer.PartialTicks

		motionX := currentMouseX - lastMouseX
		motionY := currentMouseY - lastMouseY
		lastMouseX = currentMouseX
		lastMouseY = currentMouseY

		player.Turn(motionX, motionY)

		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
		gl.LoadIdentity()

		gl.Translatef(0.0, 0.0, -0.3)

		gl.Rotated(player.XRotation, 1.0, 0.0, 0.0)
		gl.Rotated(player.YRotation, 0.0, 1.0, 0.0)

		x := player.PrevX + (player.X-player.PrevX)*float64(partialTicks)
		y := player.PrevY + (player.Y-player.PrevY)*float64(partialTicks)
		z := player.PrevZ + (player.Z-player.PrevZ)*float64(partialTicks)

		gl.Translated(-x, -y, -z)

		gl.Enable(gl.FOG)
		gl.Fogi(gl.FOG_MODE, gl.LINEAR)
		gl.Fogf(gl.FOG_START, -10)
		gl.Fogf(gl.FOG_END, 20)
		gl.Fogfv(gl.FOG_COLOR, &fogColor[0])
		gl.Disable(gl.FOG)

		renderer.Render(0)

		gl.Enable(gl.FOG)

		renderer.Render(1)

		gl.Disable(gl.TEXTURE_2D)

		window.SwapBuffers()
		glfw.PollEvents()

		frames++

		for time.Now().UnixMilli() >= lastTime+1000 {
			log.Printf("FPS: %d | Chunk: %d", frames, game.ChunkUpdates)

			game.ChunkUpdates = 0

			lastTime += 1000
			frames = 0
		}
	}
}
