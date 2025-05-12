package gl

// #cgo LDFLAGS: -lGL
//
// #include <GL/gl.h>
import "C"
import "unsafe"

const (
	Texture2D        = C.GL_TEXTURE_2D
	Smooth           = C.GL_SMOOTH
	DepthTest        = C.GL_DEPTH_TEST
	CullFace         = C.GL_CULL_FACE
	Lequal           = C.GL_LEQUAL
	Projection       = C.GL_PROJECTION
	ModelView        = C.GL_MODELVIEW
	ColorBufferBit   = C.GL_COLOR_BUFFER_BIT
	DepthBufferBit   = C.GL_DEPTH_BUFFER_BIT
	Fog              = C.GL_FOG
	FogMode          = C.GL_FOG_MODE
	Linear           = C.GL_LINEAR
	Exp              = C.GL_EXP
	FogStart         = C.GL_FOG_START
	FogDensity       = C.GL_FOG_DENSITY
	FogEnd           = C.GL_FOG_END
	FogColor         = C.GL_FOG_COLOR
	Nearest          = C.GL_NEAREST
	Compile          = C.GL_COMPILE
	ProjectionMatrix = C.GL_PROJECTION_MATRIX
	ModelViewMatrix  = C.GL_MODELVIEW_MATRIX
	Float            = C.GL_FLOAT
	VertexArray      = C.GL_VERTEX_ARRAY
	TexCoordArray    = C.GL_TEXTURE_COORD_ARRAY
	ColorArray       = C.GL_COLOR_ARRAY
	Quads            = C.GL_QUADS
	TextureMinFilter = C.GL_TEXTURE_MIN_FILTER
	TextureMagFilter = C.GL_TEXTURE_MAG_FILTER
	RGBA             = C.GL_RGBA
	UnsignedByte     = C.GL_UNSIGNED_BYTE
	Blend            = C.GL_BLEND
	SrcAlpha         = C.GL_SRC_ALPHA
	CurrentBit       = C.GL_CURRENT_BIT
	Select           = C.GL_SELECT
	Viewport         = C.GL_VIEWPORT
	Render           = C.GL_RENDER
)

func Enable(cap uint32) {
	C.glEnable(C.GLenum(cap))
}

func Disable(cap uint32) {
	C.glDisable(C.GLenum(cap))
}

func ShadeModel(mode uint32) {
	C.glShadeModel(C.GLenum(mode))
}

func Clear(mask uint32) {
	C.glClear(C.GLenum(mask))
}

func ClearColor(red, green, blue, alpha float32) {
	C.glClearColor(C.GLfloat(red), C.GLfloat(green), C.GLfloat(blue), C.GLfloat(alpha))
}

func ClearDepth(depth float64) {
	C.glClearDepth(C.GLdouble(depth))
}

func DepthFunc(funcName uint32) {
	C.glDepthFunc(C.GLenum(funcName))
}

func MatrixMode(mode uint32) {
	C.glMatrixMode(C.GLenum(mode))
}

func LoadIdentity() {
	C.glLoadIdentity()
}

func Scalef(x, y, z float32) {
	C.glScalef(C.GLfloat(x), C.GLfloat(y), C.GLfloat(z))
}

func Scaled(x, y, z float64) {
	C.glScaled(C.GLdouble(x), C.GLdouble(y), C.GLdouble(z))
}

func Translatef(x, y, z float32) {
	C.glTranslatef(C.GLfloat(x), C.GLfloat(y), C.GLfloat(z))
}

func Translated(x, y, z float64) {
	C.glTranslated(C.GLdouble(x), C.GLdouble(y), C.GLdouble(z))
}

func Rotatef(angle, x, y, z float32) {
	C.glRotatef(C.GLfloat(angle), C.GLfloat(x), C.GLfloat(y), C.GLfloat(z))
}

func Rotated(angle, x, y, z float64) {
	C.glRotated(C.GLdouble(angle), C.GLdouble(x), C.GLdouble(y), C.GLdouble(z))
}

func Fogi(pname, param uint32) {
	C.glFogi(C.GLenum(pname), C.GLint(param))
}

func Fogf(pname uint32, param float32) {
	C.glFogf(C.GLenum(pname), C.GLfloat(param))
}

func Fogfv(pname uint32, params *float32) {
	C.glFogfv(C.GLenum(pname), (*C.GLfloat)(unsafe.Pointer(params)))
}

func GenLists(r int) int {
	return int(C.glGenLists(C.GLsizei(r)))
}

func NewList(list int, mode uint32) {
	C.glNewList(C.GLenum(list), C.GLenum(mode))
}

func EndList() {
	C.glEndList()
}

func CallList(list int) {
	C.glCallList(C.GLenum(list))
}

func BindTexture(target uint32, texture int32) {
	C.glBindTexture(C.GLenum(target), C.GLuint(texture))
}

func GetFloatv(pname uint32, params *float32) {
	C.glGetFloatv(C.GLenum(pname), (*C.GLfloat)(unsafe.Pointer(params)))
}

func GetIntegerv(pname uint32, params *int) {
	C.glGetIntegerv(C.GLenum(pname), (*C.GLint)(unsafe.Pointer(params)))
}

func Vertex3f(x, y, z float32) {
	C.glVertex3f(C.GLfloat(x), C.GLfloat(y), C.GLfloat(z))
}

func VertexPointer(size int32, typeName uint32, stride int32, data *float32) {
	C.glVertexPointer(C.GLint(size), C.GLenum(typeName), C.GLsizei(stride), unsafe.Pointer(data))
}

func TexCoord2f(s, t float32) {
	C.glTexCoord2f(C.GLfloat(s), C.GLfloat(t))
}

func TexCoordPointer(size int32, typeName uint32, stride int32, data *float32) {
	C.glTexCoordPointer(C.GLint(size), C.GLenum(typeName), C.GLsizei(stride), unsafe.Pointer(data))
}

func ColorPointer(size int32, typeName uint32, stride int32, data *float32) {
	C.glColorPointer(C.GLint(size), C.GLenum(typeName), C.GLsizei(stride), unsafe.Pointer(data))
}

func EnableClientState(array uint32) {
	C.glEnableClientState(C.GLenum(array))
}

func DisableClientState(array uint32) {
	C.glDisableClientState(C.GLenum(array))
}

func DrawArrays(mode uint32, first int32, count int) {
	C.glDrawArrays(C.GLenum(mode), C.GLint(first), C.GLsizei(count))
}

func GenTextures(n int32, textures *int32) {
	C.glGenTextures(C.GLsizei(n), (*C.GLuint)(unsafe.Pointer(textures)))
}

func TexParameteri(target, pname uint32, param int32) {
	C.glTexParameteri(C.GLenum(target), C.GLenum(pname), C.GLint(param))
}

func InitNames() {
	C.glInitNames()
}

func PushName(name int) {
	C.glPushName(C.GLuint(name))
}

func PopName() {
	C.glPopName()
}

func Color3f(red, green, blue float32) {
	C.glColor3f(C.GLfloat(red), C.GLfloat(green), C.GLfloat(blue))
}

func Color4f(red, green, blue, alpha float32) {
	C.glColor4f(C.GLfloat(red), C.GLfloat(green), C.GLfloat(blue), C.GLfloat(alpha))
}

func BlendFunc(sfactor, dfactor uint32) {
	C.glBlendFunc(C.GLenum(sfactor), C.GLenum(dfactor))
}

func SelectBuffer(size int32, buffer *uint32) {
	C.glSelectBuffer(C.GLsizei(size), (*C.GLuint)(unsafe.Pointer(buffer)))
}

func RenderMode(mode uint32) int {
	return int(C.glRenderMode(C.GLenum(mode)))
}

func PushMatrix() {
	C.glPushMatrix()
}

func PopMatrix() {
	C.glPopMatrix()
}

func Begin(mode int) {
	C.glBegin(C.GLenum(mode))
}

func End() {
	C.glEnd()
}
