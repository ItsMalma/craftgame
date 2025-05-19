package windowing

import (
	"github.com/veandco/go-sdl2/sdl"
)

type SDLInput struct {
	quit bool

	keys    map[sdl.Keycode]uint8
	buttons map[uint8]uint8

	mouseDeltaX, mouseDeltaY int32
}

func (i SDLInput) IsQuit() bool {
	return i.quit
}

func (i SDLInput) IsKeyDown(key Key) bool {
	switch key {
	case KeyA:
		return i.keys[sdl.K_a] == 1
	case KeyD:
		return i.keys[sdl.K_d] == 1
	case KeyR:
		return i.keys[sdl.K_r] == 1
	case KeyS:
		return i.keys[sdl.K_s] == 1
	case KeyW:
		return i.keys[sdl.K_w] == 1
	case KeyEscape:
		return i.keys[sdl.K_ESCAPE] == 1
	case KeySpace:
		return i.keys[sdl.K_SPACE] == 1
	case KeyEnter:
		return i.keys[sdl.K_RETURN] == 1
	default:
		return false
	}
}

func (i SDLInput) IsMousePressed(button MouseButton) bool {
	switch button {
	case MouseButtonLeft:
		return i.buttons[sdl.BUTTON_LEFT] == 1
	case MouseButtonRight:
		return i.buttons[sdl.BUTTON_RIGHT] == 1
	default:
		return false
	}
}

func (i SDLInput) MouseDelta() (int, int) {
	return int(i.mouseDeltaX), int(i.mouseDeltaY)
}

type SDLWindow struct {
	title  string
	width  int
	height int

	window *sdl.Window
	input  SDLInput
}

func NewSDLWindow() Window {
	window := new(SDLWindow)
	window.input = SDLInput{
		keys:    make(map[sdl.Keycode]uint8),
		buttons: make(map[uint8]uint8),
	}

	return window
}

func (w *SDLWindow) Create(title string, width, height int) (err error) {
	w.title = title
	w.width = width
	w.height = height

	if err = sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		return err
	}

	w.window, err = sdl.CreateWindow(
		title,
		sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		int32(width), int32(height),
		sdl.WINDOW_OPENGL|sdl.WINDOW_SHOWN|sdl.WINDOW_MOUSE_FOCUS|sdl.WINDOW_INPUT_GRABBED,
	)
	if err != nil {
		return err
	}

	if _, err = w.window.GLCreateContext(); err != nil {
		return err
	}

	if err := sdl.GLSetSwapInterval(0); err != nil {
		return err
	}

	return nil
}

func (w *SDLWindow) Destroy() {
	if w.window != nil {
		w.window.Destroy()
	}

	sdl.Quit()
}

func (w *SDLWindow) Title() string {
	return w.title
}

func (w *SDLWindow) Width() int {
	return w.width
}

func (w *SDLWindow) Height() int {
	return w.height
}

func (w *SDLWindow) GrabMouse() {
	sdl.SetRelativeMouseMode(true)
}

func (w *SDLWindow) UpdateInput() Input {
	w.input.mouseDeltaX = 0
	w.input.mouseDeltaY = 0

	for button, value := range w.input.buttons {
		if value == 1 {
			w.input.buttons[button] = 2
		}
	}

	for eventInterface := sdl.PollEvent(); eventInterface != nil; eventInterface = sdl.PollEvent() {
		switch event := eventInterface.(type) {
		case *sdl.QuitEvent:
			w.input.quit = true
		case *sdl.KeyboardEvent:
			if event.Type == sdl.KEYDOWN {
				w.input.keys[event.Keysym.Sym] = 1
			}
			if event.Type == sdl.KEYUP {
				w.input.keys[event.Keysym.Sym] = 0
			}
		case *sdl.MouseButtonEvent:
			if event.Type == sdl.MOUSEBUTTONDOWN {
				w.input.buttons[event.Button] = 1
			}
			if event.Type == sdl.MOUSEBUTTONUP {
				w.input.buttons[event.Button] = 0
			}
		case *sdl.MouseMotionEvent:
			w.input.mouseDeltaX = event.XRel
			w.input.mouseDeltaY = event.YRel
		}
	}

	return w.input
}

func (w *SDLWindow) UpdateScreen() {
	w.window.GLSwap()
}
