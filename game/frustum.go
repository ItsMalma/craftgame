package game

import (
	"craftgame/pkg/gl"

	"github.com/chewxy/math32"
)

type FrustumSide int

const (
	FrustumRight FrustumSide = iota
	FrustumLeft
	FrustumBottom
	FrustumTop
	FrustumBack
	FrustumFront
)

type FrustumLocation int

const (
	FrustumA FrustumLocation = iota
	FrustumB
	FrustumC
	FrustumD
)

type Frustum struct {
	matrix [6][4]float32

	projectionBuffer [16]float32
	modelViewBuffer  [16]float32
	clippingBuffer   [16]float32
}

var frustumInstance *Frustum = NewFrustum()

func NewFrustum() *Frustum {
	frustum := new(Frustum)

	frustum.matrix = [6][4]float32{}

	frustum.projectionBuffer = [16]float32{}
	frustum.modelViewBuffer = [16]float32{}
	frustum.clippingBuffer = [16]float32{}

	return frustum
}

func GetFrustum() *Frustum {
	frustumInstance.CalculateFrustum()
	return frustumInstance
}

func (frustum *Frustum) NormalizePlane(side FrustumSide) {
	magnitude := math32.Sqrt(
		frustum.matrix[side][FrustumA]*frustum.matrix[side][FrustumA] +
			frustum.matrix[side][FrustumB]*frustum.matrix[side][FrustumB] +
			frustum.matrix[side][FrustumC]*frustum.matrix[side][FrustumC],
	)

	frustum.matrix[side][FrustumA] /= magnitude
	frustum.matrix[side][FrustumB] /= magnitude
	frustum.matrix[side][FrustumC] /= magnitude
	frustum.matrix[side][FrustumD] /= magnitude
}

func (frustum *Frustum) CalculateFrustum() {
	gl.GetFloatv(gl.ProjectionMatrix, &frustum.projectionBuffer[0])
	gl.GetFloatv(gl.ModelViewMatrix, &frustum.modelViewBuffer[0])

	frustum.clippingBuffer[0] = frustum.modelViewBuffer[0]*frustum.projectionBuffer[0] + frustum.modelViewBuffer[1]*frustum.projectionBuffer[4] + frustum.modelViewBuffer[2]*frustum.projectionBuffer[8] + frustum.modelViewBuffer[3]*frustum.projectionBuffer[12]
	frustum.clippingBuffer[1] = frustum.modelViewBuffer[0]*frustum.projectionBuffer[1] + frustum.modelViewBuffer[1]*frustum.projectionBuffer[5] + frustum.modelViewBuffer[2]*frustum.projectionBuffer[9] + frustum.modelViewBuffer[3]*frustum.projectionBuffer[13]
	frustum.clippingBuffer[2] = frustum.modelViewBuffer[0]*frustum.projectionBuffer[2] + frustum.modelViewBuffer[1]*frustum.projectionBuffer[6] + frustum.modelViewBuffer[2]*frustum.projectionBuffer[10] + frustum.modelViewBuffer[3]*frustum.projectionBuffer[14]
	frustum.clippingBuffer[3] = frustum.modelViewBuffer[0]*frustum.projectionBuffer[3] + frustum.modelViewBuffer[1]*frustum.projectionBuffer[7] + frustum.modelViewBuffer[2]*frustum.projectionBuffer[11] + frustum.modelViewBuffer[3]*frustum.projectionBuffer[15]

	frustum.clippingBuffer[4] = frustum.modelViewBuffer[4]*frustum.projectionBuffer[0] + frustum.modelViewBuffer[5]*frustum.projectionBuffer[4] + frustum.modelViewBuffer[6]*frustum.projectionBuffer[8] + frustum.modelViewBuffer[7]*frustum.projectionBuffer[12]
	frustum.clippingBuffer[5] = frustum.modelViewBuffer[4]*frustum.projectionBuffer[1] + frustum.modelViewBuffer[5]*frustum.projectionBuffer[5] + frustum.modelViewBuffer[6]*frustum.projectionBuffer[9] + frustum.modelViewBuffer[7]*frustum.projectionBuffer[13]
	frustum.clippingBuffer[6] = frustum.modelViewBuffer[4]*frustum.projectionBuffer[2] + frustum.modelViewBuffer[5]*frustum.projectionBuffer[6] + frustum.modelViewBuffer[6]*frustum.projectionBuffer[10] + frustum.modelViewBuffer[7]*frustum.projectionBuffer[14]
	frustum.clippingBuffer[7] = frustum.modelViewBuffer[4]*frustum.projectionBuffer[3] + frustum.modelViewBuffer[5]*frustum.projectionBuffer[7] + frustum.modelViewBuffer[6]*frustum.projectionBuffer[11] + frustum.modelViewBuffer[7]*frustum.projectionBuffer[15]

	frustum.clippingBuffer[8] = frustum.modelViewBuffer[8]*frustum.projectionBuffer[0] + frustum.modelViewBuffer[9]*frustum.projectionBuffer[4] + frustum.modelViewBuffer[10]*frustum.projectionBuffer[8] + frustum.modelViewBuffer[11]*frustum.projectionBuffer[12]
	frustum.clippingBuffer[9] = frustum.modelViewBuffer[8]*frustum.projectionBuffer[1] + frustum.modelViewBuffer[9]*frustum.projectionBuffer[5] + frustum.modelViewBuffer[10]*frustum.projectionBuffer[9] + frustum.modelViewBuffer[11]*frustum.projectionBuffer[13]
	frustum.clippingBuffer[10] = frustum.modelViewBuffer[8]*frustum.projectionBuffer[2] + frustum.modelViewBuffer[9]*frustum.projectionBuffer[6] + frustum.modelViewBuffer[10]*frustum.projectionBuffer[10] + frustum.modelViewBuffer[11]*frustum.projectionBuffer[14]
	frustum.clippingBuffer[11] = frustum.modelViewBuffer[8]*frustum.projectionBuffer[3] + frustum.modelViewBuffer[9]*frustum.projectionBuffer[7] + frustum.modelViewBuffer[10]*frustum.projectionBuffer[11] + frustum.modelViewBuffer[11]*frustum.projectionBuffer[15]

	frustum.clippingBuffer[12] = frustum.modelViewBuffer[12]*frustum.projectionBuffer[0] + frustum.modelViewBuffer[13]*frustum.projectionBuffer[4] + frustum.modelViewBuffer[14]*frustum.projectionBuffer[8] + frustum.modelViewBuffer[15]*frustum.projectionBuffer[12]
	frustum.clippingBuffer[13] = frustum.modelViewBuffer[12]*frustum.projectionBuffer[1] + frustum.modelViewBuffer[13]*frustum.projectionBuffer[5] + frustum.modelViewBuffer[14]*frustum.projectionBuffer[9] + frustum.modelViewBuffer[15]*frustum.projectionBuffer[13]
	frustum.clippingBuffer[14] = frustum.modelViewBuffer[12]*frustum.projectionBuffer[2] + frustum.modelViewBuffer[13]*frustum.projectionBuffer[6] + frustum.modelViewBuffer[14]*frustum.projectionBuffer[10] + frustum.modelViewBuffer[15]*frustum.projectionBuffer[14]
	frustum.clippingBuffer[15] = frustum.modelViewBuffer[12]*frustum.projectionBuffer[3] + frustum.modelViewBuffer[13]*frustum.projectionBuffer[7] + frustum.modelViewBuffer[14]*frustum.projectionBuffer[11] + frustum.modelViewBuffer[15]*frustum.projectionBuffer[15]

	frustum.matrix[FrustumRight][FrustumA] = frustum.clippingBuffer[3] - frustum.clippingBuffer[0]
	frustum.matrix[FrustumRight][FrustumB] = frustum.clippingBuffer[7] - frustum.clippingBuffer[4]
	frustum.matrix[FrustumRight][FrustumC] = frustum.clippingBuffer[11] - frustum.clippingBuffer[8]
	frustum.matrix[FrustumRight][FrustumD] = frustum.clippingBuffer[15] - frustum.clippingBuffer[12]
	frustum.NormalizePlane(FrustumRight)

	frustum.matrix[FrustumLeft][FrustumA] = frustum.clippingBuffer[3] + frustum.clippingBuffer[0]
	frustum.matrix[FrustumLeft][FrustumB] = frustum.clippingBuffer[7] + frustum.clippingBuffer[4]
	frustum.matrix[FrustumLeft][FrustumC] = frustum.clippingBuffer[11] + frustum.clippingBuffer[8]
	frustum.matrix[FrustumLeft][FrustumD] = frustum.clippingBuffer[15] + frustum.clippingBuffer[12]
	frustum.NormalizePlane(FrustumLeft)

	frustum.matrix[FrustumBottom][FrustumA] = frustum.clippingBuffer[3] + frustum.clippingBuffer[1]
	frustum.matrix[FrustumBottom][FrustumB] = frustum.clippingBuffer[7] + frustum.clippingBuffer[5]
	frustum.matrix[FrustumBottom][FrustumC] = frustum.clippingBuffer[11] + frustum.clippingBuffer[9]
	frustum.matrix[FrustumBottom][FrustumD] = frustum.clippingBuffer[15] + frustum.clippingBuffer[13]
	frustum.NormalizePlane(FrustumBottom)

	frustum.matrix[FrustumTop][FrustumA] = frustum.clippingBuffer[3] - frustum.clippingBuffer[1]
	frustum.matrix[FrustumTop][FrustumB] = frustum.clippingBuffer[7] - frustum.clippingBuffer[5]
	frustum.matrix[FrustumTop][FrustumC] = frustum.clippingBuffer[11] - frustum.clippingBuffer[9]
	frustum.matrix[FrustumTop][FrustumD] = frustum.clippingBuffer[15] - frustum.clippingBuffer[13]
	frustum.NormalizePlane(FrustumTop)

	frustum.matrix[FrustumBack][FrustumA] = frustum.clippingBuffer[3] - frustum.clippingBuffer[2]
	frustum.matrix[FrustumBack][FrustumB] = frustum.clippingBuffer[7] - frustum.clippingBuffer[6]
	frustum.matrix[FrustumBack][FrustumC] = frustum.clippingBuffer[11] - frustum.clippingBuffer[10]
	frustum.matrix[FrustumBack][FrustumD] = frustum.clippingBuffer[15] - frustum.clippingBuffer[14]
	frustum.NormalizePlane(FrustumBack)

	frustum.matrix[FrustumFront][FrustumA] = frustum.clippingBuffer[3] + frustum.clippingBuffer[2]
	frustum.matrix[FrustumFront][FrustumB] = frustum.clippingBuffer[7] + frustum.clippingBuffer[6]
	frustum.matrix[FrustumFront][FrustumC] = frustum.clippingBuffer[11] + frustum.clippingBuffer[10]
	frustum.matrix[FrustumFront][FrustumD] = frustum.clippingBuffer[15] + frustum.clippingBuffer[14]
	frustum.NormalizePlane(FrustumFront)
}

func (frustum *Frustum) CubeFullyInFrustum(minX, minY, minZ, maxX, maxY, maxZ float32) bool {
	for i := range 6 {
		if frustum.matrix[i][FrustumA]*minX+frustum.matrix[i][FrustumB]*minY+frustum.matrix[i][FrustumC]*minZ+frustum.matrix[i][FrustumD] > 0 {
			continue
		}
		if frustum.matrix[i][FrustumA]*maxX+frustum.matrix[i][FrustumB]*minY+frustum.matrix[i][FrustumC]*minZ+frustum.matrix[i][FrustumD] > 0 {
			continue
		}
		if frustum.matrix[i][FrustumA]*minX+frustum.matrix[i][FrustumB]*maxY+frustum.matrix[i][FrustumC]*minZ+frustum.matrix[i][FrustumD] > 0 {
			continue
		}
		if frustum.matrix[i][FrustumA]*maxX+frustum.matrix[i][FrustumB]*maxY+frustum.matrix[i][FrustumC]*minZ+frustum.matrix[i][FrustumD] > 0 {
			continue
		}
		if frustum.matrix[i][FrustumA]*minX+frustum.matrix[i][FrustumB]*minY+frustum.matrix[i][FrustumC]*maxZ+frustum.matrix[i][FrustumD] > 0 {
			continue
		}
		if frustum.matrix[i][FrustumA]*maxX+frustum.matrix[i][FrustumB]*minY+frustum.matrix[i][FrustumC]*maxZ+frustum.matrix[i][FrustumD] > 0 {
			continue
		}
		if frustum.matrix[i][FrustumA]*minX+frustum.matrix[i][FrustumB]*maxY+frustum.matrix[i][FrustumC]*maxZ+frustum.matrix[i][FrustumD] > 0 {
			continue
		}
		if frustum.matrix[i][FrustumA]*maxX+frustum.matrix[i][FrustumB]*maxY+frustum.matrix[i][FrustumC]*maxZ+frustum.matrix[i][FrustumD] > 0 {
			continue
		}

		return false
	}

	return true
}

func (frustum *Frustum) CubeInFrustum(minX, minY, minZ, maxX, maxY, maxZ float32) bool {
	for i := range 6 {
		if !(frustum.matrix[i][0]*minX+frustum.matrix[i][1]*minY+frustum.matrix[i][2]*minZ+frustum.matrix[i][3] > 0.0) &&
			!(frustum.matrix[i][0]*maxX+frustum.matrix[i][1]*minY+frustum.matrix[i][2]*minZ+frustum.matrix[i][3] > 0.0) &&
			!(frustum.matrix[i][0]*minX+frustum.matrix[i][1]*maxY+frustum.matrix[i][2]*minZ+frustum.matrix[i][3] > 0.0) &&
			!(frustum.matrix[i][0]*maxX+frustum.matrix[i][1]*maxY+frustum.matrix[i][2]*minZ+frustum.matrix[i][3] > 0.0) &&
			!(frustum.matrix[i][0]*minX+frustum.matrix[i][1]*minY+frustum.matrix[i][2]*maxZ+frustum.matrix[i][3] > 0.0) &&
			!(frustum.matrix[i][0]*maxX+frustum.matrix[i][1]*minY+frustum.matrix[i][2]*maxZ+frustum.matrix[i][3] > 0.0) &&
			!(frustum.matrix[i][0]*minX+frustum.matrix[i][1]*maxY+frustum.matrix[i][2]*maxZ+frustum.matrix[i][3] > 0.0) &&
			!(frustum.matrix[i][0]*maxX+frustum.matrix[i][1]*maxY+frustum.matrix[i][2]*maxZ+frustum.matrix[i][3] > 0.0) {
			return false
		}
	}

	return true
}

func (frustum *Frustum) CubeInFrustumAABB(aabb AABB) bool {
	return frustum.CubeInFrustum(
		aabb.MinX,
		aabb.MinY,
		aabb.MinZ,
		aabb.MaxX,
		aabb.MaxY,
		aabb.MaxZ,
	)
}
