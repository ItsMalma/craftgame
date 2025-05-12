package game

type HitResult struct {
	X, Y, Z    int
	Type, Face int
}

func NewHitResult(x, y, z, typ, face int) *HitResult {
	hitResult := new(HitResult)
	hitResult.X = x
	hitResult.Y = y
	hitResult.Z = z
	hitResult.Type = typ
	hitResult.Face = face

	return hitResult
}
