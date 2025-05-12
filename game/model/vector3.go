package model

type Vector3 struct {
	X, Y, Z float32
}

func NewVector3(x, y, z float32) Vector3 {
	return Vector3{x, y, z}
}

func (vector3 Vector3) InterpolateTo(target Vector3, partialTicks float32) Vector3 {
	xTarget := vector3.X + (target.X-vector3.X)*partialTicks
	yTarget := vector3.Y + (target.Y-vector3.Y)*partialTicks
	zTarget := vector3.Z + (target.Z-vector3.Z)*partialTicks

	return NewVector3(xTarget, yTarget, zTarget)
}

func (vector3 Vector3) Set(x, y, z float32) Vector3 {
	vector3.X = x
	vector3.Y = y
	vector3.Z = z

	return vector3
}
