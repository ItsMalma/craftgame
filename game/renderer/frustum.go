package renderer

import (
	"github.com/ItsMalma/craftgame/phys"
	"github.com/go-gl/gl/v2.1/gl"
)

const (
	FrustumRight  = 0
	FrustumLeft   = 1
	FrustumBottom = 2
	FrustumTop    = 3
	FrustumBack   = 4
	FrustumFront  = 5

	FrustumA = 0
	FrustumB = 1
	FrustumC = 2
	FrustumD = 3
)

var frustumInstance = newFrustum()

func GetFrustum() *Frustum {
	frustumInstance.CalculateFrustum()
	return frustumInstance
}

type Frustum struct {
	Matrix [6][4]float32

	projectionBuffer [16]float32
	modelviewBuffer  [16]float32
	clippingBuffer   [16]float32
}

func newFrustum() *Frustum {
	frustum := new(Frustum)
	frustum.Matrix = [6][4]float32{}
	frustum.projectionBuffer = [16]float32{}
	frustum.modelviewBuffer = [16]float32{}
	frustum.clippingBuffer = [16]float32{}

	return frustum
}

func (frustum *Frustum) NormalizePlane(side int) {
	magnitude := float32(
		frustum.Matrix[side][0]*frustum.Matrix[side][0] +
			frustum.Matrix[side][1]*frustum.Matrix[side][1] +
			frustum.Matrix[side][2]*frustum.Matrix[side][2],
	)
	frustum.Matrix[side][0] /= magnitude
	frustum.Matrix[side][1] /= magnitude
	frustum.Matrix[side][2] /= magnitude
	frustum.Matrix[side][3] /= magnitude
}

func (frustum *Frustum) CalculateFrustum() {
	gl.GetFloatv(gl.PROJECTION_MATRIX, &frustum.projectionBuffer[0])
	gl.GetFloatv(gl.MODELVIEW_MATRIX, &frustum.modelviewBuffer[0])

	frustum.clippingBuffer[0] = frustum.modelviewBuffer[0]*frustum.projectionBuffer[0] + frustum.modelviewBuffer[1]*frustum.projectionBuffer[4] + frustum.modelviewBuffer[2]*frustum.projectionBuffer[8] + frustum.modelviewBuffer[3]*frustum.projectionBuffer[12]
	frustum.clippingBuffer[1] = frustum.modelviewBuffer[0]*frustum.projectionBuffer[1] + frustum.modelviewBuffer[1]*frustum.projectionBuffer[5] + frustum.modelviewBuffer[2]*frustum.projectionBuffer[9] + frustum.modelviewBuffer[3]*frustum.projectionBuffer[13]
	frustum.clippingBuffer[2] = frustum.modelviewBuffer[0]*frustum.projectionBuffer[2] + frustum.modelviewBuffer[1]*frustum.projectionBuffer[6] + frustum.modelviewBuffer[2]*frustum.projectionBuffer[10] + frustum.modelviewBuffer[3]*frustum.projectionBuffer[14]
	frustum.clippingBuffer[3] = frustum.modelviewBuffer[0]*frustum.projectionBuffer[3] + frustum.modelviewBuffer[1]*frustum.projectionBuffer[7] + frustum.modelviewBuffer[2]*frustum.projectionBuffer[11] + frustum.modelviewBuffer[3]*frustum.projectionBuffer[15]
	frustum.clippingBuffer[4] = frustum.modelviewBuffer[4]*frustum.projectionBuffer[0] + frustum.modelviewBuffer[5]*frustum.projectionBuffer[4] + frustum.modelviewBuffer[6]*frustum.projectionBuffer[8] + frustum.modelviewBuffer[7]*frustum.projectionBuffer[12]
	frustum.clippingBuffer[5] = frustum.modelviewBuffer[4]*frustum.projectionBuffer[1] + frustum.modelviewBuffer[5]*frustum.projectionBuffer[5] + frustum.modelviewBuffer[6]*frustum.projectionBuffer[9] + frustum.modelviewBuffer[7]*frustum.projectionBuffer[13]
	frustum.clippingBuffer[6] = frustum.modelviewBuffer[4]*frustum.projectionBuffer[2] + frustum.modelviewBuffer[5]*frustum.projectionBuffer[6] + frustum.modelviewBuffer[6]*frustum.projectionBuffer[10] + frustum.modelviewBuffer[7]*frustum.projectionBuffer[14]
	frustum.clippingBuffer[7] = frustum.modelviewBuffer[4]*frustum.projectionBuffer[3] + frustum.modelviewBuffer[5]*frustum.projectionBuffer[7] + frustum.modelviewBuffer[6]*frustum.projectionBuffer[11] + frustum.modelviewBuffer[7]*frustum.projectionBuffer[15]
	frustum.clippingBuffer[8] = frustum.modelviewBuffer[8]*frustum.projectionBuffer[0] + frustum.modelviewBuffer[9]*frustum.projectionBuffer[4] + frustum.modelviewBuffer[10]*frustum.projectionBuffer[8] + frustum.modelviewBuffer[11]*frustum.projectionBuffer[12]
	frustum.clippingBuffer[9] = frustum.modelviewBuffer[8]*frustum.projectionBuffer[1] + frustum.modelviewBuffer[9]*frustum.projectionBuffer[5] + frustum.modelviewBuffer[10]*frustum.projectionBuffer[9] + frustum.modelviewBuffer[11]*frustum.projectionBuffer[13]
	frustum.clippingBuffer[10] = frustum.modelviewBuffer[8]*frustum.projectionBuffer[2] + frustum.modelviewBuffer[9]*frustum.projectionBuffer[6] + frustum.modelviewBuffer[10]*frustum.projectionBuffer[10] + frustum.modelviewBuffer[11]*frustum.projectionBuffer[14]
	frustum.clippingBuffer[11] = frustum.modelviewBuffer[8]*frustum.projectionBuffer[3] + frustum.modelviewBuffer[9]*frustum.projectionBuffer[7] + frustum.modelviewBuffer[10]*frustum.projectionBuffer[11] + frustum.modelviewBuffer[11]*frustum.projectionBuffer[15]
	frustum.clippingBuffer[12] = frustum.modelviewBuffer[12]*frustum.projectionBuffer[0] + frustum.modelviewBuffer[13]*frustum.projectionBuffer[4] + frustum.modelviewBuffer[14]*frustum.projectionBuffer[8] + frustum.modelviewBuffer[15]*frustum.projectionBuffer[12]
	frustum.clippingBuffer[13] = frustum.modelviewBuffer[12]*frustum.projectionBuffer[1] + frustum.modelviewBuffer[13]*frustum.projectionBuffer[5] + frustum.modelviewBuffer[14]*frustum.projectionBuffer[9] + frustum.modelviewBuffer[15]*frustum.projectionBuffer[13]
	frustum.clippingBuffer[14] = frustum.modelviewBuffer[12]*frustum.projectionBuffer[2] + frustum.modelviewBuffer[13]*frustum.projectionBuffer[6] + frustum.modelviewBuffer[14]*frustum.projectionBuffer[10] + frustum.modelviewBuffer[15]*frustum.projectionBuffer[14]
	frustum.clippingBuffer[15] = frustum.modelviewBuffer[12]*frustum.projectionBuffer[3] + frustum.modelviewBuffer[13]*frustum.projectionBuffer[7] + frustum.modelviewBuffer[14]*frustum.projectionBuffer[11] + frustum.modelviewBuffer[15]*frustum.projectionBuffer[15]
	frustum.Matrix[0][0] = frustum.clippingBuffer[3] - frustum.clippingBuffer[0]
	frustum.Matrix[0][1] = frustum.clippingBuffer[7] - frustum.clippingBuffer[4]
	frustum.Matrix[0][2] = frustum.clippingBuffer[11] - frustum.clippingBuffer[8]
	frustum.Matrix[0][3] = frustum.clippingBuffer[15] - frustum.clippingBuffer[12]
	frustum.NormalizePlane(0)
	frustum.Matrix[1][0] = frustum.clippingBuffer[3] + frustum.clippingBuffer[0]
	frustum.Matrix[1][1] = frustum.clippingBuffer[7] + frustum.clippingBuffer[4]
	frustum.Matrix[1][2] = frustum.clippingBuffer[11] + frustum.clippingBuffer[8]
	frustum.Matrix[1][3] = frustum.clippingBuffer[15] + frustum.clippingBuffer[12]
	frustum.NormalizePlane(1)
	frustum.Matrix[2][0] = frustum.clippingBuffer[3] + frustum.clippingBuffer[1]
	frustum.Matrix[2][1] = frustum.clippingBuffer[7] + frustum.clippingBuffer[5]
	frustum.Matrix[2][2] = frustum.clippingBuffer[11] + frustum.clippingBuffer[9]
	frustum.Matrix[2][3] = frustum.clippingBuffer[15] + frustum.clippingBuffer[13]
	frustum.NormalizePlane(2)
	frustum.Matrix[3][0] = frustum.clippingBuffer[3] - frustum.clippingBuffer[1]
	frustum.Matrix[3][1] = frustum.clippingBuffer[7] - frustum.clippingBuffer[5]
	frustum.Matrix[3][2] = frustum.clippingBuffer[11] - frustum.clippingBuffer[9]
	frustum.Matrix[3][3] = frustum.clippingBuffer[15] - frustum.clippingBuffer[13]
	frustum.NormalizePlane(3)
	frustum.Matrix[4][0] = frustum.clippingBuffer[3] - frustum.clippingBuffer[2]
	frustum.Matrix[4][1] = frustum.clippingBuffer[7] - frustum.clippingBuffer[6]
	frustum.Matrix[4][2] = frustum.clippingBuffer[11] - frustum.clippingBuffer[10]
	frustum.Matrix[4][3] = frustum.clippingBuffer[15] - frustum.clippingBuffer[14]
	frustum.NormalizePlane(4)
	frustum.Matrix[5][0] = frustum.clippingBuffer[3] + frustum.clippingBuffer[2]
	frustum.Matrix[5][1] = frustum.clippingBuffer[7] + frustum.clippingBuffer[6]
	frustum.Matrix[5][2] = frustum.clippingBuffer[11] + frustum.clippingBuffer[10]
	frustum.Matrix[5][3] = frustum.clippingBuffer[15] + frustum.clippingBuffer[14]
	frustum.NormalizePlane(5)
}

func (frustum *Frustum) PointInFrustum(x, y, z float32) bool {
	for i := 0; i < 6; i++ {
		if frustum.Matrix[i][0]*x+frustum.Matrix[i][1]*y+frustum.Matrix[i][2]*z+frustum.Matrix[i][3] <= 0.0 {
			return false
		}
	}
	return true
}

func (frustum *Frustum) SphereInFrustum(x, y, z, radius float32) bool {
	for i := range 6 {
		if frustum.Matrix[i][0]*x+frustum.Matrix[i][1]*y+frustum.Matrix[i][2]*z+frustum.Matrix[i][3] <= -radius {
			return false
		}
	}
	return true
}

func (frustum *Frustum) CubeFullyInFrustum(x1, y1, z1, x2, y2, z2 float32) bool {
	for i := range 6 {
		if frustum.Matrix[i][FrustumA]*x1+frustum.Matrix[i][FrustumB]*y1+frustum.Matrix[i][FrustumC]*z1+frustum.Matrix[i][FrustumD] > 0 {
			continue
		}
		if frustum.Matrix[i][FrustumA]*x2+frustum.Matrix[i][FrustumB]*y1+frustum.Matrix[i][FrustumC]*z1+frustum.Matrix[i][FrustumD] > 0 {
			continue
		}
		if frustum.Matrix[i][FrustumA]*x1+frustum.Matrix[i][FrustumB]*y2+frustum.Matrix[i][FrustumC]*z1+frustum.Matrix[i][FrustumD] > 0 {
			continue
		}
		if frustum.Matrix[i][FrustumA]*x2+frustum.Matrix[i][FrustumB]*y2+frustum.Matrix[i][FrustumC]*z1+frustum.Matrix[i][FrustumD] > 0 {
			continue
		}
		if frustum.Matrix[i][FrustumA]*x1+frustum.Matrix[i][FrustumB]*y1+frustum.Matrix[i][FrustumC]*z2+frustum.Matrix[i][FrustumD] > 0 {
			continue
		}
		if frustum.Matrix[i][FrustumA]*x2+frustum.Matrix[i][FrustumB]*y1+frustum.Matrix[i][FrustumC]*z2+frustum.Matrix[i][FrustumD] > 0 {
			continue
		}
		if frustum.Matrix[i][FrustumA]*x1+frustum.Matrix[i][FrustumB]*y2+frustum.Matrix[i][FrustumC]*z2+frustum.Matrix[i][FrustumD] > 0 {
			continue
		}
		if frustum.Matrix[i][FrustumA]*x2+frustum.Matrix[i][FrustumB]*y2+frustum.Matrix[i][FrustumC]*z2+frustum.Matrix[i][FrustumD] > 0 {
			continue
		}

		return false
	}

	return true
}

func (frustum *Frustum) CubeInFrustum(x1, y1, z1, x2, y2, z2 float32) bool {
	for i := range 6 {
		if !(frustum.Matrix[i][0]*x1+frustum.Matrix[i][1]*y1+frustum.Matrix[i][2]*z1+frustum.Matrix[i][3] > 0.0) &&
			!(frustum.Matrix[i][0]*x2+frustum.Matrix[i][1]*y1+frustum.Matrix[i][2]*z1+frustum.Matrix[i][3] > 0.0) &&
			!(frustum.Matrix[i][0]*x1+frustum.Matrix[i][1]*y2+frustum.Matrix[i][2]*z1+frustum.Matrix[i][3] > 0.0) &&
			!(frustum.Matrix[i][0]*x2+frustum.Matrix[i][1]*y2+frustum.Matrix[i][2]*z1+frustum.Matrix[i][3] > 0.0) &&
			!(frustum.Matrix[i][0]*x1+frustum.Matrix[i][1]*y1+frustum.Matrix[i][2]*z2+frustum.Matrix[i][3] > 0.0) &&
			!(frustum.Matrix[i][0]*x2+frustum.Matrix[i][1]*y1+frustum.Matrix[i][2]*z2+frustum.Matrix[i][3] > 0.0) &&
			!(frustum.Matrix[i][0]*x1+frustum.Matrix[i][1]*y2+frustum.Matrix[i][2]*z2+frustum.Matrix[i][3] > 0.0) &&
			!(frustum.Matrix[i][0]*x2+frustum.Matrix[i][1]*y2+frustum.Matrix[i][2]*z2+frustum.Matrix[i][3] > 0.0) {
			return false
		}
	}

	return true

}

func (frustum *Frustum) AABBInFrustum(aabb phys.AABB) bool {
	return frustum.CubeInFrustum(aabb.X0, aabb.Y0, aabb.Z0, aabb.X1, aabb.Y1, aabb.Z1)
}
