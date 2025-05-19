package windowing

type Key int

const (
	KeyA Key = iota
	KeyD
	KeyR
	KeyS
	KeyW

	KeyEscape
	KeySpace
	KeyEnter
)

type MouseButton int

const (
	MouseButtonLeft MouseButton = iota
	MouseButtonRight
)

type Input interface {
	IsQuit() bool

	IsKeyDown(key Key) bool

	IsMousePressed(button MouseButton) bool
	MouseDelta() (int, int)
}
