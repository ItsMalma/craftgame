package windowing

type Window interface {
	Create(title string, width, height int) error
	Destroy()

	Title() string
	Width() int
	Height() int

	GrabMouse()

	UpdateInput() Input
	UpdateScreen()
}
