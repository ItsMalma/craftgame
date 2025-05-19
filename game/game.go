package game

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/ItsMalma/craftgame/game/entity"
	"github.com/ItsMalma/craftgame/game/renderer"
	"github.com/ItsMalma/craftgame/game/windowing"
	"github.com/ItsMalma/craftgame/game/world"
	"github.com/ItsMalma/craftgame/glu"
	"github.com/go-gl/gl/v2.1/gl"
)

type Game struct {
	window  windowing.Window
	onError func(error)

	fogColor [4]float32

	timer Timer

	world         *world.World
	worldRenderer *renderer.WorldRenderer
	player        *entity.Player
	bobs          []*entity.Bob

	viewportBuffer [16]int32
	selectBuffer   [2000]uint32

	hasHitResult bool
	hitResult    renderer.HitResult

	textures map[string]uint32
}

func NewGame() *Game {
	g := new(Game)

	g.onError = func(originErr error) {
		errorFile, err := os.Create("error.log")
		if err != nil {
			log.Fatal(err)
		}
		defer errorFile.Close()

		if _, err := errorFile.WriteString(originErr.Error()); err != nil {
			log.Fatal(err)
		}
	}

	g.viewportBuffer = [16]int32{}
	g.selectBuffer = [2000]uint32{}

	g.textures = make(map[string]uint32)

	return g
}

func (g *Game) Init() {
	g.timer = NewTimer(60)

	col := 920330
	fr := float32(0.5)
	fg := float32(0.8)
	fb := float32(1.0)
	g.fogColor = [4]float32{
		float32(col>>16&255) / 255.0,
		float32(col>>8&255) / 255.0,
		float32(col&255) / 255.0,
		1.0,
	}

	var err error

	g.window = windowing.NewSDLWindow()
	if err = g.window.Create("CraftGame", 1024, 768); err != nil {
		g.onError(err)
	}

	if err := gl.Init(); err != nil {
		g.onError(err)
	}

	gl.Enable(gl.TEXTURE_2D)
	gl.ShadeModel(gl.SMOOTH)
	gl.ClearColor(fr, fg, fb, 1.0)
	gl.ClearDepth(1.0)
	gl.Enable(gl.DEPTH_TEST)
	gl.DepthFunc(gl.LEQUAL)
	gl.MatrixMode(gl.PROJECTION)
	gl.LoadIdentity()
	gl.MatrixMode(gl.MODELVIEW)
	g.world, err = world.New(256, 256, 64)
	if err != nil {
		g.onError(err)
	}
	g.worldRenderer = renderer.NewWorldRenderer(g.world, g.LoadTexture("terrain.png", gl.NEAREST))
	g.player = entity.NewPlayer(g.world)

	g.window.GrabMouse()

	for range 100 {
		g.bobs = append(g.bobs, entity.NewBob(g.world, 128.0, 0, 128.0, g.LoadTexture("bob.png", gl.NEAREST)))
	}
}

func (g *Game) Destroy() {
	g.world.Save()
	g.window.Destroy()
}

func (g *Game) Run() {
	g.Init()

	lastTime := time.Now().UnixMilli()
	frames := 0

	running := true

	for running {
		input := g.window.UpdateInput()
		if input.IsQuit() || input.IsKeyDown(windowing.KeyEscape) {
			running = false
		}

		g.timer = g.timer.AdvanceTime()

		for range g.timer.Ticks {
			g.Tick(input)
		}

		g.Render(input)
		frames++

		for time.Now().UnixMilli() >= lastTime+1000 {
			fmt.Printf("FPS: %d | Chunk: %d\n", frames, renderer.ChunkUpdates)
			renderer.ChunkUpdates = 0
			lastTime += 1000
			frames = 0
		}
	}
}

func (g *Game) Tick(input windowing.Input) {
	for _, bob := range g.bobs {
		bob.Tick()
	}

	g.player.Tick(input)
}

func (g *Game) moveCameraToPlayer() {
	gl.Translatef(0.0, 0.0, -0.3)
	gl.Rotatef(g.player.XRot, 1.0, 0.0, 0.0)
	gl.Rotatef(g.player.YRot, 0.0, 1.0, 0.0)
	gl.Translatef(
		-(g.player.XO + (g.player.X-g.player.XO)*g.timer.A),
		-(g.player.YO + (g.player.Y-g.player.YO)*g.timer.A),
		-(g.player.ZO + (g.player.Z-g.player.ZO)*g.timer.A),
	)
}

func (g *Game) setupCamera() {
	gl.MatrixMode(gl.PROJECTION)
	gl.LoadIdentity()
	glu.Perspective(70.0, float32(g.window.Width())/float32(g.window.Height()), 0.05, 1000.0)
	gl.MatrixMode(gl.MODELVIEW)
	gl.LoadIdentity()
	g.moveCameraToPlayer()
}

func (g *Game) setupPickCamera(x, y int) {
	gl.MatrixMode(gl.PROJECTION)
	gl.LoadIdentity()
	gl.GetIntegerv(gl.VIEWPORT, &g.viewportBuffer[0])
	glu.PickMatrix(float32(x), float32(y), 5.0, 5.0, &g.viewportBuffer[0])
	glu.Perspective(70.0, float32(g.window.Width())/float32(g.window.Height()), 0.05, 1000.0)
	gl.MatrixMode(gl.MODELVIEW)
	gl.LoadIdentity()
	g.moveCameraToPlayer()
}

func (g *Game) pick() {
	gl.SelectBuffer(2000, &g.selectBuffer[0])
	gl.RenderMode(gl.SELECT)
	g.setupPickCamera(g.window.Width()/2, g.window.Height()/2)
	g.worldRenderer.Pick(g.player)
	hits := gl.RenderMode(gl.RENDER)
	closest := uint32(0)
	names := [10]int{}
	hitNameCount := 0
	selectBufferIndex := 0

	for i := range hits {
		nameCount := g.selectBuffer[selectBufferIndex]
		selectBufferIndex++
		minZ := g.selectBuffer[selectBufferIndex]
		selectBufferIndex += 2

		if minZ >= closest && i != 0 {
			for range nameCount {
				selectBufferIndex++
			}
		} else {
			closest = minZ
			hitNameCount = int(nameCount)

			for j := uint32(0); j < nameCount; j++ {
				names[j] = int(g.selectBuffer[selectBufferIndex])
				selectBufferIndex++
			}
		}
	}

	if hitNameCount > 0 {
		g.hasHitResult = true
		g.hitResult = renderer.HitResult{
			X: names[0],
			Y: names[1],
			Z: names[2],
			O: names[3],
			F: names[4],
		}
	} else {
		g.hasHitResult = false
	}
}

func (g *Game) Render(input windowing.Input) {
	dx, dy := input.MouseDelta()
	g.player.Turn(float32(dx), float32(dy))

	g.pick()

	if input.IsMousePressed(windowing.MouseButtonLeft) && g.hasHitResult {
		g.world.SetTile(g.hitResult.X, g.hitResult.Y, g.hitResult.Z, 0)
	}
	if input.IsMousePressed(windowing.MouseButtonRight) && g.hasHitResult {
		x := g.hitResult.X
		y := g.hitResult.Y
		z := g.hitResult.Z
		if g.hitResult.F == 0 {
			y--
		}
		if g.hitResult.F == 1 {
			y++
		}
		if g.hitResult.F == 2 {
			z--
		}
		if g.hitResult.F == 3 {
			z++
		}
		if g.hitResult.F == 4 {
			x--
		}
		if g.hitResult.F == 5 {
			x++
		}
		g.world.SetTile(x, y, z, 1)
	}
	if input.IsKeyDown(windowing.KeyEnter) {
		g.world.Save()
	}

	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

	g.setupCamera()

	gl.Enable(gl.CULL_FACE)
	gl.Enable(gl.FOG)
	gl.Fogi(gl.FOG_MODE, gl.EXP)
	gl.Fogf(gl.FOG_DENSITY, 0.2)
	gl.Fogfv(gl.FOG_COLOR, &g.fogColor[0])
	gl.Disable(gl.FOG)

	g.worldRenderer.Render(g.player, 0)

	for _, bob := range g.bobs {
		bob.Render(g.timer.A)
	}

	gl.Enable(gl.FOG)

	g.worldRenderer.Render(g.player, 1)

	gl.Disable(gl.TEXTURE_2D)

	if g.hasHitResult {
		g.worldRenderer.RenderHit(g.hitResult)
	}

	gl.Disable(gl.FOG)

	g.window.UpdateScreen()
}
