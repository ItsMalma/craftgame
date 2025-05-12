package game

type AABB struct {
	epsilon float32

	MinX, MinY, MinZ, MaxX, MaxY, MaxZ float32
}

func NewAABB(minX, minY, minZ, maxX, maxY, maxZ float32) AABB {
	aabb := AABB{}

	aabb.epsilon = 0.0

	aabb.MinX = minX
	aabb.MinY = minY
	aabb.MinZ = minZ
	aabb.MaxX = maxX
	aabb.MaxY = maxY
	aabb.MaxZ = maxZ

	return aabb
}

func (aabb AABB) Clone() AABB {
	return NewAABB(aabb.MinX, aabb.MinY, aabb.MinZ, aabb.MaxX, aabb.MaxY, aabb.MaxZ)
}

func (aabb AABB) Expand(x, y, z float32) AABB {
	minX := aabb.MinX
	minY := aabb.MinY
	minZ := aabb.MinZ
	maxX := aabb.MaxX
	maxY := aabb.MaxY
	maxZ := aabb.MaxZ

	if x < 0.0 {
		minX += x
	}
	if x > 0.0 {
		maxX += x
	}

	if y < 0.0 {
		minY += y
	}
	if y > 0.0 {
		maxY += y
	}

	if z < 0.0 {
		minZ += z
	}
	if z > 0.0 {
		maxZ += z
	}

	return NewAABB(minX, minY, minZ, maxX, maxY, maxZ)
}

func (aabb AABB) Grow(x, y, z float32) AABB {
	return NewAABB(
		aabb.MinX-x, aabb.MinY-y, aabb.MinZ-z,
		aabb.MaxX+x, aabb.MaxY+y, aabb.MaxZ+z,
	)
}

func (aabb AABB) ClipXCollide(otherBoundingBox AABB, x float32) float32 {
	if !(otherBoundingBox.MaxY <= aabb.MinY) && !(otherBoundingBox.MinY >= aabb.MaxY) {
		if !(otherBoundingBox.MaxZ <= aabb.MinZ) && !(otherBoundingBox.MinZ >= aabb.MaxZ) {
			var max float32 = 0.0
			if x > 0.0 && otherBoundingBox.MaxX <= aabb.MinX {
				max = aabb.MinX - otherBoundingBox.MaxX - aabb.epsilon
				if max < x {
					x = max
				}
			}

			if x < 0.0 && otherBoundingBox.MinX >= aabb.MaxX {
				max = aabb.MaxX - otherBoundingBox.MinX + aabb.epsilon
				if max > x {
					x = max
				}
			}

			return x
		} else {
			return x
		}
	} else {
		return x
	}
}

func (aabb AABB) ClipYCollide(c AABB, y float32) float32 {
	if !(c.MaxX <= aabb.MinX) && !(c.MinX >= aabb.MaxX) {
		if !(c.MaxZ <= aabb.MinZ) && !(c.MinZ >= aabb.MaxZ) {
			var max float32 = 0.0
			if y > 0.0 && c.MaxY <= aabb.MinY {
				max = aabb.MinY - c.MaxY - aabb.epsilon
				if max < y {
					y = max
				}
			}

			if y < 0.0 && c.MinY >= aabb.MaxY {
				max = aabb.MaxY - c.MinY + aabb.epsilon
				if max > y {
					y = max
				}
			}

			return y
		} else {
			return y
		}
	} else {
		return y
	}
}

func (aabb AABB) ClipZCollide(c AABB, z float32) float32 {
	if !(c.MaxX <= aabb.MinX) && !(c.MinX >= aabb.MaxX) {
		if !(c.MaxY <= aabb.MinY) && !(c.MinY >= aabb.MaxY) {
			var max float32 = 0.0
			if z > 0.0 && c.MaxZ <= aabb.MinZ {
				max = aabb.MinZ - c.MaxZ - aabb.epsilon
				if max < z {
					z = max
				}
			}

			if z < 0.0 && c.MinZ >= aabb.MaxZ {
				max = aabb.MaxZ - c.MinZ + aabb.epsilon
				if max > z {
					z = max
				}
			}

			return z
		} else {
			return z
		}
	} else {
		return z
	}
}

func (aabb AABB) Intersects(c AABB) bool {
	if !(c.MaxX <= aabb.MinX) && !(c.MinX >= aabb.MaxX) {
		if !(c.MaxY <= aabb.MinY) && !(c.MinY >= aabb.MaxY) {
			return !(c.MaxZ <= aabb.MinZ) && !(c.MinZ >= aabb.MaxZ)
		} else {
			return false
		}
	} else {
		return false
	}
}

func (aabb AABB) Move(x, y, z float32) AABB {
	aabb.MinX += x
	aabb.MinY += y
	aabb.MinZ += z
	aabb.MaxX += x
	aabb.MaxY += y
	aabb.MaxZ += z

	return aabb
}

func (aabb AABB) Offset(x, y, z float32) AABB {
	return NewAABB(
		aabb.MinX+x, aabb.MinY+y, aabb.MinZ+z,
		aabb.MaxX+x, aabb.MaxY+y, aabb.MaxZ+z,
	)
}
