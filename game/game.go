package game

import (
	"craftgame/pkg/gl"
	"craftgame/pkg/glfw"
	"craftgame/pkg/glu"
	"errors"
	"log"
	"time"
)

const (
	fullscreenMode = false
)

type Game struct {
	window                            *glfw.Window
	width, height                     int
	lastMouseX, lastMouseY            float64
	prevLeftMbState, prevRightMbState int

	fogColor [4]float32

	timer *Timer

	level    *Level
	renderer *Renderer
	player   *Player

	bobs []*Bob

	viewportBuffer [16]int
	selectBuffer   [2000]uint32
	hitResult      *HitResult
}

func NewGame() (*Game, error) {
	game := new(Game)

	game.timer = NewTimer(60)

	game.viewportBuffer = [16]int{}
	game.selectBuffer = [2000]uint32{}

	const (
		color    = 920330
		fogRed   = 0.5
		fogGreen = 0.8
		fogBlue  = 1.0
	)

	game.fogColor = [4]float32{
		(color >> 16 & 255) / 255.0,
		(color >> 8 & 255) / 255.0,
		(color & 255) / 255.0,
		1.0,
	}

	game.width = 1024
	game.height = 768

	if !glfw.Init() {
		return nil, errors.New("failed to initialize GLFW")
	}
	glfw.WindowHint(glfw.ContextVersionMajor, 1)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	game.window = glfw.CreateWindow(
		game.width, game.height,
		"CraftGame",
		nil, nil,
	)
	glfw.SetInputMode(game.window, glfw.CursorMode, glfw.CursorDisabled)
	glfw.MakeContextCurrent(game.window)
	glfw.SwapInterval(0)

	gl.Enable(gl.Texture2D)
	gl.ShadeModel(gl.Smooth)
	gl.ClearColor(fogRed, fogGreen, fogBlue, 0.0)
	gl.ClearDepth(1.0)
	gl.Enable(gl.DepthTest)
	gl.DepthFunc(gl.Lequal)
	gl.MatrixMode(gl.Projection)
	gl.LoadIdentity()
	gl.MatrixMode(gl.ModelView)

	var err error

	game.level, err = NewLevel(256, 256, 64)
	if err != nil {
		return nil, err
	}
	game.renderer = NewRenderer(game.level)
	game.player = NewPlayer(game.level)

	game.bobs = make([]*Bob, 100)
	for i := range 100 {
		game.bobs[i] = NewBob(game.level, 128, 0, 128)
	}

	return game, nil
}

func (game *Game) tick() {
	for _, bob := range game.bobs {
		bob.Tick()
	}

	game.player.Tick(game.window)
}

func (game *Game) moveCameraToPlayer() {
	gl.Translatef(0.0, 0.0, -0.3)

	gl.Rotatef(game.player.Entity.XRotation, 1.0, 0.0, 0.0)
	gl.Rotatef(game.player.Entity.YRotation, 0.0, 1.0, 0.0)

	x := game.player.Entity.PrevX + (game.player.Entity.X-game.player.Entity.PrevX)*game.timer.PartialTicks
	y := game.player.Entity.PrevY + (game.player.Entity.Y-game.player.Entity.PrevY)*game.timer.PartialTicks
	z := game.player.Entity.PrevZ + (game.player.Entity.Z-game.player.Entity.PrevZ)*game.timer.PartialTicks

	gl.Translatef(-x, -y, -z)
}

func (game *Game) setupCamera() {
	gl.MatrixMode(gl.Projection)
	gl.LoadIdentity()

	glu.Perspective(70, float64(game.width)/float64(game.height), 0.05, 1000.0)

	gl.MatrixMode(gl.ModelView)
	gl.LoadIdentity()

	game.moveCameraToPlayer()
}

func (game *Game) setupPickCamera(x, y int) {
	gl.MatrixMode(gl.Projection)
	gl.LoadIdentity()

	gl.GetIntegerv(gl.Viewport, &game.viewportBuffer[0])

	glu.PickMatrix(float64(x), float64(y), 5, 5, &game.viewportBuffer[0])
	glu.Perspective(70, float64(game.width)/float64(game.height), 0.05, 1000.0)

	gl.MatrixMode(gl.ModelView)
	gl.LoadIdentity()

	game.moveCameraToPlayer()
}

func (game *Game) pick() {
	game.selectBuffer = [2000]uint32{}

	gl.SelectBuffer(2000, &game.selectBuffer[0])
	gl.RenderMode(gl.Select)

	game.setupPickCamera(game.width/2, game.height/2)

	game.renderer.Pick(game.player)

	selectBufferIndex := 0

	var closest uint32 = 0
	names := [10]int{}
	hitNameCount := 0

	hits := gl.RenderMode(gl.Render)
	for hitIndex := range hits {
		nameCount := game.selectBuffer[selectBufferIndex]
		selectBufferIndex++
		minZ := game.selectBuffer[selectBufferIndex]
		selectBufferIndex++
		selectBufferIndex++

		if minZ < closest || hitIndex == 0 {
			closest = minZ
			hitNameCount = int(nameCount)

			for nameIndex := range nameCount {
				names[nameIndex] = int(game.selectBuffer[selectBufferIndex])
				selectBufferIndex++
			}
		} else {
			for range nameCount {
				selectBufferIndex++
			}
		}
	}

	if hitNameCount > 0 {
		game.hitResult = NewHitResult(names[0], names[1], names[2], names[3], names[4])
	} else {
		game.hitResult = nil
	}
}

func (game *Game) render() {
	currentMouseX, currentMouseY := glfw.GetCursorPos(game.window)
	motionX := currentMouseX - game.lastMouseX
	motionY := currentMouseY - game.lastMouseY

	game.lastMouseX = currentMouseX
	game.lastMouseY = currentMouseY

	game.player.Entity.Turn(float32(motionX), float32(motionY))

	game.pick()

	leftMbState := glfw.GetMouseButton(game.window, glfw.MouseButton1)
	if leftMbState == glfw.Press && game.prevLeftMbState == glfw.Release && game.hitResult != nil {
		game.level.SetTile(game.hitResult.X, game.hitResult.Y, game.hitResult.Z, 0)
	}
	game.prevLeftMbState = leftMbState

	x := 0
	rightMbState := glfw.GetMouseButton(game.window, glfw.MouseButton2)
	if rightMbState == glfw.Press && game.prevRightMbState == glfw.Release && game.hitResult != nil {
		x = game.hitResult.X
		y := game.hitResult.Y
		z := game.hitResult.Z

		if game.hitResult.Face == 0 {
			y--
		}

		if game.hitResult.Face == 1 {
			y++
		}

		if game.hitResult.Face == 2 {
			z--
		}

		if game.hitResult.Face == 3 {
			z++
		}

		if game.hitResult.Face == 4 {
			x--
		}

		if game.hitResult.Face == 5 {
			x++
		}

		game.level.SetTile(x, y, z, 1)
	}
	game.prevRightMbState = rightMbState

	gl.Clear(gl.ColorBufferBit | gl.DepthBufferBit)

	game.setupCamera()

	gl.Enable(gl.CullFace)
	gl.Enable(gl.Fog)
	gl.Fogi(gl.FogMode, gl.Exp)
	gl.Fogf(gl.FogDensity, 0.2)
	gl.Fogfv(gl.FogColor, &game.fogColor[0])
	gl.Disable(gl.Fog)

	if err := game.renderer.Render(game.player, 0); err != nil {
		panic(err)
	}

	for _, bob := range game.bobs {
		if err := bob.Render(game.timer.PartialTicks); err != nil {
			panic(err)
		}
	}

	gl.Enable(gl.Fog)

	game.renderer.Render(game.player, 1)

	gl.Disable(gl.Texture2D)

	if game.hitResult != nil {
		game.renderer.RenderHit(game.hitResult)
	}

	gl.Disable(gl.Fog)

	// model.NewCube(0, 0)

	glfw.SwapBuffers(game.window)
	glfw.PollEvents()
}

func (game *Game) Run() {
	lastTime := time.Now().UnixMilli()
	frames := 0

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

		frames++

		for time.Now().UnixMilli() >= lastTime+1000 {
			log.Printf("FPS: %d | Chunk: %d\n", frames, ChunkUpdates)

			ChunkUpdates = 0

			lastTime += 1000
			frames = 0
		}
	}
}

func (game *Game) Close() {
	glfw.DestroyWindow(game.window)

	glfw.Terminate()

	game.level.Save()
}
