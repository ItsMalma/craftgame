package game

import (
	"github.com/chewxy/math32"
	"github.com/go-gl/gl/v2.1/gl"
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
	values [6][4]float32

	modelView  [16]float32
	projection [16]float32
}

var frustumInstance *Frustum = NewFrustum()

func NewFrustum() *Frustum {
	frustum := new(Frustum)

	frustum.values = [6][4]float32{}

	frustum.modelView = [16]float32{}
	frustum.projection = [16]float32{}

	return frustum
}

func GetFrustum() *Frustum {
	frustumInstance.CalculateFrustum()
	return frustumInstance
}

func (frustum *Frustum) NormalizePlane(frustumValues [6][4]float32, side FrustumSide) {
	magnitude := math32.Sqrt(
		frustumValues[side][FrustumA]*frustumValues[side][FrustumA] +
			frustumValues[side][FrustumB]*frustumValues[side][FrustumB] +
			frustumValues[side][FrustumC]*frustumValues[side][FrustumC],
	)

	frustumValues[side][FrustumA] /= magnitude
	frustumValues[side][FrustumB] /= magnitude
	frustumValues[side][FrustumC] /= magnitude
	frustumValues[side][FrustumD] /= magnitude
}

func (frustum *Frustum) CalculateFrustum() {
	projection := [16]float32{}
	modelView := [16]float32{}
	clipping := [16]float32{}

	gl.GetFloatv(gl.PROJECTION_MATRIX, &frustum.projection[0])
	copy(projection[:], frustum.projection[:])

	gl.GetFloatv(gl.MODELVIEW_MATRIX, &frustum.modelView[0])
	copy(modelView[:], frustum.modelView[:])

	clipping[0] = modelView[0]*projection[0] + modelView[1]*projection[4] + modelView[2]*projection[8] + modelView[3]*projection[12]
	clipping[1] = modelView[0]*projection[1] + modelView[1]*projection[5] + modelView[2]*projection[9] + modelView[3]*projection[13]
	clipping[2] = modelView[0]*projection[2] + modelView[1]*projection[6] + modelView[2]*projection[10] + modelView[3]*projection[14]
	clipping[3] = modelView[0]*projection[3] + modelView[1]*projection[7] + modelView[2]*projection[11] + modelView[3]*projection[15]

	clipping[4] = modelView[4]*projection[0] + modelView[5]*projection[4] + modelView[6]*projection[8] + modelView[7]*projection[12]
	clipping[5] = modelView[4]*projection[1] + modelView[5]*projection[5] + modelView[6]*projection[9] + modelView[7]*projection[13]
	clipping[6] = modelView[4]*projection[2] + modelView[5]*projection[6] + modelView[6]*projection[10] + modelView[7]*projection[14]
	clipping[7] = modelView[4]*projection[3] + modelView[5]*projection[7] + modelView[6]*projection[11] + modelView[7]*projection[15]

	clipping[8] = modelView[8]*projection[0] + modelView[9]*projection[4] + modelView[10]*projection[8] + modelView[11]*projection[12]
	clipping[9] = modelView[8]*projection[1] + modelView[9]*projection[5] + modelView[10]*projection[9] + modelView[11]*projection[13]
	clipping[10] = modelView[8]*projection[2] + modelView[9]*projection[6] + modelView[10]*projection[10] + modelView[11]*projection[14]
	clipping[11] = modelView[8]*projection[3] + modelView[9]*projection[7] + modelView[10]*projection[11] + modelView[11]*projection[15]

	clipping[12] = modelView[12]*projection[0] + modelView[13]*projection[4] + modelView[14]*projection[8] + modelView[15]*projection[12]
	clipping[13] = modelView[12]*projection[1] + modelView[13]*projection[5] + modelView[14]*projection[9] + modelView[15]*projection[13]
	clipping[14] = modelView[12]*projection[2] + modelView[13]*projection[6] + modelView[14]*projection[10] + modelView[15]*projection[14]
	clipping[15] = modelView[12]*projection[3] + modelView[13]*projection[7] + modelView[14]*projection[11] + modelView[15]*projection[15]

	frustum.values[FrustumRight][FrustumA] = clipping[3] - clipping[0]
	frustum.values[FrustumRight][FrustumB] = clipping[7] - clipping[4]
	frustum.values[FrustumRight][FrustumC] = clipping[11] - clipping[8]
	frustum.values[FrustumRight][FrustumD] = clipping[15] - clipping[12]
	frustum.NormalizePlane(frustum.values, FrustumRight)

	frustum.values[FrustumLeft][FrustumA] = clipping[3] + clipping[0]
	frustum.values[FrustumLeft][FrustumB] = clipping[7] + clipping[4]
	frustum.values[FrustumLeft][FrustumC] = clipping[11] + clipping[8]
	frustum.values[FrustumLeft][FrustumD] = clipping[15] + clipping[12]
	frustum.NormalizePlane(frustum.values, FrustumLeft)

	frustum.values[FrustumBottom][FrustumA] = clipping[3] + clipping[1]
	frustum.values[FrustumBottom][FrustumB] = clipping[7] + clipping[5]
	frustum.values[FrustumBottom][FrustumC] = clipping[11] + clipping[9]
	frustum.values[FrustumBottom][FrustumD] = clipping[15] + clipping[13]
	frustum.NormalizePlane(frustum.values, FrustumBottom)

	frustum.values[FrustumTop][FrustumA] = clipping[3] - clipping[1]
	frustum.values[FrustumTop][FrustumB] = clipping[7] - clipping[5]
	frustum.values[FrustumTop][FrustumC] = clipping[11] - clipping[9]
	frustum.values[FrustumTop][FrustumD] = clipping[15] - clipping[13]
	frustum.NormalizePlane(frustum.values, FrustumTop)

	frustum.values[FrustumBack][FrustumA] = clipping[3] - clipping[2]
	frustum.values[FrustumBack][FrustumB] = clipping[7] - clipping[6]
	frustum.values[FrustumBack][FrustumC] = clipping[11] - clipping[10]
	frustum.values[FrustumBack][FrustumD] = clipping[15] - clipping[14]
	frustum.NormalizePlane(frustum.values, FrustumBack)

	frustum.values[FrustumFront][FrustumA] = clipping[3] + clipping[2]
	frustum.values[FrustumFront][FrustumB] = clipping[7] + clipping[6]
	frustum.values[FrustumFront][FrustumC] = clipping[11] + clipping[10]
	frustum.values[FrustumFront][FrustumD] = clipping[15] + clipping[14]
	frustum.NormalizePlane(frustum.values, FrustumFront)
}

func (frustum *Frustum) CubeInFrustum(minX, minY, minZ, maxX, maxY, maxZ float32) bool {
	for i := range 6 {
		if frustum.values[i][FrustumA]*minX+frustum.values[i][FrustumB]*minY+frustum.values[i][FrustumC]*minZ+frustum.values[i][FrustumD] > 0 {
			continue
		}
		if frustum.values[i][FrustumA]*maxX+frustum.values[i][FrustumB]*minY+frustum.values[i][FrustumC]*minZ+frustum.values[i][FrustumD] > 0 {
			continue
		}
		if frustum.values[i][FrustumA]*minX+frustum.values[i][FrustumB]*maxY+frustum.values[i][FrustumC]*minZ+frustum.values[i][FrustumD] > 0 {
			continue
		}
		if frustum.values[i][FrustumA]*maxX+frustum.values[i][FrustumB]*maxY+frustum.values[i][FrustumC]*minZ+frustum.values[i][FrustumD] > 0 {
			continue
		}
		if frustum.values[i][FrustumA]*minX+frustum.values[i][FrustumB]*minY+frustum.values[i][FrustumC]*maxZ+frustum.values[i][FrustumD] > 0 {
			continue
		}
		if frustum.values[i][FrustumA]*maxX+frustum.values[i][FrustumB]*minY+frustum.values[i][FrustumC]*maxZ+frustum.values[i][FrustumD] > 0 {
			continue
		}
		if frustum.values[i][FrustumA]*minX+frustum.values[i][FrustumB]*maxY+frustum.values[i][FrustumC]*maxZ+frustum.values[i][FrustumD] > 0 {
			continue
		}
		if frustum.values[i][FrustumA]*maxX+frustum.values[i][FrustumB]*maxY+frustum.values[i][FrustumC]*maxZ+frustum.values[i][FrustumD] > 0 {
			continue
		}

		return false
	}

	return true
}

func (frustum *Frustum) CubeInFrustumAABB(aabb *AABB) bool {
	return frustum.CubeInFrustum(
		float32(aabb.MinX),
		float32(aabb.MinY),
		float32(aabb.MinZ),
		float32(aabb.MaxX),
		float32(aabb.MaxY),
		float32(aabb.MaxZ),
	)
}
