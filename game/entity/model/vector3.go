package model

type Vector3 struct {
	X, Y, Z float32
}

func NewVector3(x, y, z float32) Vector3 {
	return Vector3{
		X: x,
		Y: y,
		Z: z,
	}
}

func (vector3 Vector3) InterpolateTo(t Vector3, p float32) Vector3 {
	return Vector3{
		X: vector3.X + (t.X-vector3.X)*p,
		Y: vector3.Y + (t.Y-vector3.Y)*p,
		Z: vector3.Z + (t.Z-vector3.Z)*p,
	}
}

func (vector3 Vector3) Set(x, y, z float32) Vector3 {
	vector3.X = x
	vector3.Y = y
	vector3.Z = z
	return vector3
}
