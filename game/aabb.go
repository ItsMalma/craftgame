package game

type AABB struct {
	epsilon float64

	MinX, MinY, MinZ, MaxX, MaxY, MaxZ float64
}

func NewAABB(minX, minY, minZ, maxX, maxY, maxZ float64) *AABB {
	aabb := new(AABB)

	aabb.epsilon = 0.0

	aabb.MinX = minX
	aabb.MinY = minY
	aabb.MinZ = minZ
	aabb.MaxX = maxX
	aabb.MaxY = maxY
	aabb.MaxZ = maxZ

	return aabb
}

func (aabb *AABB) Expand(x, y, z float64) *AABB {
	minX := aabb.MinX
	minY := aabb.MinY
	minZ := aabb.MinZ
	maxX := aabb.MaxX
	maxY := aabb.MaxY
	maxZ := aabb.MaxZ

	if x < 0.0 {
		minX += x
	} else {
		maxX += x
	}

	if y < 0.0 {
		minY += y
	} else {
		maxY += y
	}

	if z < 0.0 {
		minZ += z
	} else {
		maxZ += z
	}

	return NewAABB(minX, minY, minZ, maxX, maxY, maxZ)
}

func (aabb *AABB) ClipXCollide(otherBoundingBox *AABB, x float64) float64 {
	if otherBoundingBox.MaxY <= aabb.MinY || otherBoundingBox.MinY >= aabb.MaxY {
		return x
	}

	if otherBoundingBox.MaxZ <= aabb.MinZ || otherBoundingBox.MinZ >= aabb.MaxZ {
		return x
	}

	if x > 0.0 && otherBoundingBox.MaxX <= aabb.MinX {
		max := aabb.MinX - otherBoundingBox.MaxX - aabb.epsilon
		if max < x {
			x = max
		}
	}

	if x < 0.0 && otherBoundingBox.MinX >= aabb.MaxX {
		max := aabb.MaxX - otherBoundingBox.MinX + aabb.epsilon
		if max > x {
			x = max
		}
	}

	return x
}

func (aabb *AABB) ClipYCollide(otherBoundingBox *AABB, y float64) float64 {
	if otherBoundingBox.MaxX <= aabb.MinX || otherBoundingBox.MinX >= aabb.MaxX {
		return y
	}

	if otherBoundingBox.MaxZ <= aabb.MinZ || otherBoundingBox.MinZ >= aabb.MaxZ {
		return y
	}

	if y > 0.0 && otherBoundingBox.MaxY <= aabb.MinY {
		max := aabb.MinY - otherBoundingBox.MaxY - aabb.epsilon
		if max < y {
			y = max
		}
	}

	if y < 0.0 && otherBoundingBox.MinY >= aabb.MaxY {
		max := aabb.MaxY - otherBoundingBox.MinY + aabb.epsilon
		if max > y {
			y = max
		}
	}

	return y
}

func (aabb *AABB) ClipZCollide(otherBoundingBox *AABB, z float64) float64 {
	if otherBoundingBox.MaxX <= aabb.MinX || otherBoundingBox.MinX >= aabb.MaxX {
		return z
	}

	if otherBoundingBox.MaxY <= aabb.MinY || otherBoundingBox.MinY >= aabb.MaxY {
		return z
	}

	if z > 0.0 && otherBoundingBox.MaxZ <= aabb.MinZ {
		max := aabb.MinZ - otherBoundingBox.MaxZ - aabb.epsilon
		if max < z {
			z = max
		}
	}

	if z < 0.0 && otherBoundingBox.MinZ >= aabb.MaxZ {
		max := aabb.MaxZ - otherBoundingBox.MinZ + aabb.epsilon
		if max > z {
			z = max
		}
	}

	return z
}

func (aabb *AABB) Move(x, y, z float64) {
	aabb.MinX += x
	aabb.MinY += y
	aabb.MinZ += z
	aabb.MaxX += x
	aabb.MaxY += y
	aabb.MaxZ += z
}
