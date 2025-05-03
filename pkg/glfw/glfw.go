package glfw

// #cgo LDFLAGS: -lglfw
//
// #include <GLFW/glfw3.h>
import "C"

const (
	CursorMode     = C.GLFW_CURSOR
	CursorDisabled = C.GLFW_CURSOR_DISABLED

	KeyEscape = C.GLFW_KEY_ESCAPE
	KeyR      = C.GLFW_KEY_R
	KeyW      = C.GLFW_KEY_W
	KeyS      = C.GLFW_KEY_S
	KeyA      = C.GLFW_KEY_A
	KeyD      = C.GLFW_KEY_D
	KeySpace  = C.GLFW_KEY_SPACE
	Press     = C.GLFW_PRESS
)

type Window = C.GLFWwindow

func Init() bool {
	return C.glfwInit() == 1
}

func Terminate() {
	C.glfwTerminate()
}

func CreateWindow(
	width, height int,
	title string,
	monitor *C.GLFWmonitor,
	share *C.GLFWwindow,
) *Window {
	return C.glfwCreateWindow(
		C.int(width),
		C.int(height),
		C.CString(title),
		monitor,
		share,
	)
}

func MakeContextCurrent(window *Window) {
	C.glfwMakeContextCurrent(window)
}

func SwapInterval(interval int) {
	C.glfwSwapInterval(C.int(interval))
}

func SetInputMode(window *Window, mode int, value int) {
	C.glfwSetInputMode(window, C.int(mode), C.int(value))
}

func GetCursorPos(window *Window) (float64, float64) {
	var xpos, ypos C.double
	C.glfwGetCursorPos(window, &xpos, &ypos)
	return float64(xpos), float64(ypos)
}

func ShouldClose(window *Window) bool {
	return C.glfwWindowShouldClose(window) == 1
}

func SetShouldClose(window *Window, value bool) {
	if value {
		C.glfwSetWindowShouldClose(window, 1)
	} else {
		C.glfwSetWindowShouldClose(window, 0)
	}
}

func GetKey(window *Window, key int) int {
	return int(C.glfwGetKey(window, C.int(key)))
}

func SwapBuffers(window *Window) {
	C.glfwSwapBuffers(window)
}

func PollEvents() {
	C.glfwPollEvents()
}

func DestroyWindow(window *Window) {
	C.glfwDestroyWindow(window)
}
