package game

import (
	"errors"
	"fmt"
	"minecraft/pkg/gl"
	"minecraft/pkg/glfw"
	"minecraft/pkg/glu"
	"time"
)

type Game struct {
	window *glfw.Window

	timer    *Timer
	frames   int
	lastTime int64

	level    *Level
	renderer *Renderer
	player   *Player

	fogColor [4]float32

	lastMouseX float64
	lastMouseY float64
}

func NewGame() (*Game, error) {
	game := new(Game)

	game.timer = NewTimer(60)

	game.fogColor = [4]float32{
		14 / 255.0,
		11 / 255.0,
		10 / 255.0,
		255 / 255.0,
	}

	const width = 1024
	const height = 768

	if !glfw.Init() {
		return nil, errors.New("failed to initialize GLFW")
	}
	game.window = glfw.CreateWindow(
		width, height,
		"CraftGame",
		nil, nil,
	)
	glfw.MakeContextCurrent(game.window)
	glfw.SwapInterval(0)

	gl.Enable(gl.Texture2D)
	gl.ShadeModel(gl.Smooth)
	gl.ClearColor(0.5, 0.8, 1.0, 0.0)
	gl.ClearDepth(1.0)
	gl.Enable(gl.DepthTest)
	gl.Enable(gl.CullFace)
	gl.DepthFunc(gl.Lequal)

	gl.MatrixMode(gl.Projection)
	gl.LoadIdentity()
	glu.Perspective(70, float64(width)/float64(height), 0.05, 1000.0)
	gl.MatrixMode(gl.ModelView)

	var err error

	texture, err := LoadTexture("terrain.png", gl.Nearest)
	if err != nil {
		return nil, err
	}

	game.level, err = NewLevel(256, 256, 64)
	if err != nil {
		return nil, err
	}
	game.renderer = NewRenderer(game.level, texture)
	game.player = NewPlayer(game.level)

	glfw.SetInputMode(game.window, glfw.CursorMode, glfw.CursorDisabled)

	return game, nil
}

func (game *Game) tick() {
	game.player.Tick(game.window)
}

func (game *Game) render() {
	currentMouseX, currentMouseY := glfw.GetCursorPos(game.window)
	motionX := currentMouseX - game.lastMouseX
	motionY := currentMouseY - game.lastMouseY

	game.lastMouseX = currentMouseX
	game.lastMouseY = currentMouseY

	game.player.Turn(motionX, motionY)

	gl.Clear(gl.ColorBufferBit | gl.DepthBufferBit)
	gl.LoadIdentity()

	gl.Translatef(0.0, 0.0, -0.3)

	gl.Rotated(game.player.XRotation, 1.0, 0.0, 0.0)
	gl.Rotated(game.player.YRotation, 0.0, 1.0, 0.0)

	x := game.player.PrevX + (game.player.X-game.player.PrevX)*float64(game.timer.PartialTicks)
	y := game.player.PrevY + (game.player.Y-game.player.PrevY)*float64(game.timer.PartialTicks)
	z := game.player.PrevZ + (game.player.Z-game.player.PrevZ)*float64(game.timer.PartialTicks)

	gl.Translated(-x, -y, -z)

	gl.Enable(gl.Fog)
	gl.Fogi(gl.FogMode, gl.Linear)
	gl.Fogf(gl.FogStart, -10)
	gl.Fogf(gl.FogEnd, 20)
	gl.Fogfv(gl.FogColor, &game.fogColor[0])
	gl.Disable(gl.Fog)

	game.renderer.Render(0)

	gl.Enable(gl.Fog)

	game.renderer.Render(1)

	gl.Disable(gl.Fog)

	glfw.SwapBuffers(game.window)
	glfw.PollEvents()
}

func (game *Game) Run() {
	game.frames = 0
	game.lastTime = time.Now().UnixMilli()

	game.lastMouseX, game.lastMouseY = glfw.GetCursorPos(game.window)

	for !glfw.ShouldClose(game.window) {
		if glfw.GetKey(game.window, glfw.KeyEscape) == glfw.Press {
			glfw.SetShouldClose(game.window, true)
		}

		game.timer.AdvanceTime()

		for range game.timer.Ticks {
			game.tick()
		}

		game.render()

		game.frames++

		for time.Now().UnixMilli() >= game.lastTime+1000 {
			fmt.Printf("FPS: %d | Chunk: %d\n", game.frames, ChunkUpdates)

			ChunkUpdates = 0

			game.lastTime += 1000
			game.frames = 0
		}
	}
}

func (game *Game) Close() {
	glfw.DestroyWindow(game.window)

	glfw.Terminate()

	game.level.Save()
}
