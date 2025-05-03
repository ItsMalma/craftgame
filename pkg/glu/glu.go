package glu

// #cgo linux LDFLAGS: -lGLU
// #cgo windows LDFLAGS: -lglu32
//
// #include <GL/glu.h>
import "C"
import "unsafe"

func Perspective(fovy, aspect, zNear, zFar float64) {
	C.gluPerspective(
		C.GLdouble(fovy),
		C.GLdouble(aspect),
		C.GLdouble(zNear),
		C.GLdouble(zFar),
	)
}

func Build2DMipmaps(target, internalFormat, width, height, format, typ int32, data *int32) int32 {
	return int32(C.gluBuild2DMipmaps(
		C.GLenum(target),
		C.GLint(internalFormat),
		C.GLsizei(width),
		C.GLsizei(height),
		C.GLenum(format),
		C.GLenum(typ),
		unsafe.Pointer(data),
	))
}
