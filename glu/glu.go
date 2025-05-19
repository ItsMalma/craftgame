package glu

// #cgo darwin LDFLAGS: -framework Carbon -framework OpenGL -framework GLUT
// #cgo linux LDFLAGS: -lGLU
// #cgo windows LDFLAGS: -lglu32
//
// #ifdef __APPLE__
//   #include <OpenGL/glu.h>
// #else
//   #include <GL/glu.h>
// #endif
import "C"
import "unsafe"

func Build2DMipmaps(target uint32, internalFormat int32, width, height int, format uint32, type_ uint32, data *int32) int32 {
	return int32(C.gluBuild2DMipmaps(
		C.GLenum(target),
		C.GLint(internalFormat),
		C.GLsizei(width),
		C.GLsizei(height),
		C.GLenum(format),
		C.GLenum(type_),
		unsafe.Pointer(data),
	))
}

func Perspective(fovY, aspect, zNear, zFar float32) {
	C.gluPerspective(
		C.GLdouble(fovY),
		C.GLdouble(aspect),
		C.GLdouble(zNear),
		C.GLdouble(zFar),
	)
}

func PickMatrix(x, y, delX, delY float32, viewport *int32) {
	C.gluPickMatrix(
		C.GLdouble(x),
		C.GLdouble(y),
		C.GLdouble(delX),
		C.GLdouble(delY),
		(*C.GLint)(unsafe.Pointer(viewport)),
	)
}
